package converter

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/dstotijn/go-notion"
	"github.com/stretchr/testify/assert"
)

func TestMarkdownConverter_fixtures(t *testing.T) {
	basePath := "../../fixtures/"
	files, err := os.ReadDir(basePath)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("can convert fixtures to notion blocks without failing", func(t *testing.T) {
		for _, file := range files {
			fixture, err := os.ReadFile(basePath + file.Name())
			if err != nil {
				t.Fatal(err)
			}

			result, err := Convert(string(fixture))
			assert.NoError(t, err)
			assert.NotNil(t, result)
		}
	})
}

func TestMarkdownConverter(t *testing.T) {

	t.Run("can convert markdown to notion blocks", func(t *testing.T) {
		markdownText := `# H1 Example`
		expected := []notion.Block{
			notion.Heading1Block{
				RichText: []notion.RichText{
					{
						Type:      notion.RichTextTypeText,
						Text:      &notion.Text{Content: "H1 Example"},
						PlainText: "H1 Example",
					},
				},
			},
		}

		result, err := Convert(markdownText)
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("can convert markdown with multiple blocks", func(t *testing.T) {
		markdownText := `# H1 Example

## H2 Example

*This is a paragraph*
lorem ipsum dolor sit amet.
`

		expected := []notion.Block{
			notion.Heading1Block{
				RichText: []notion.RichText{
					{
						Type:      notion.RichTextTypeText,
						Text:      &notion.Text{Content: "H1 Example"},
						PlainText: "H1 Example",
					},
				},
			},
			notion.Heading2Block{
				RichText: []notion.RichText{
					{
						Type:      notion.RichTextTypeText,
						Text:      &notion.Text{Content: "H2 Example"},
						PlainText: "H2 Example",
					},
				},
			},
			&notion.ParagraphBlock{
				RichText: []notion.RichText{
					{
						Type:      notion.RichTextTypeText,
						Text:      &notion.Text{Content: "This is a paragraph"},
						PlainText: "This is a paragraph",
						Annotations: &notion.Annotations{
							Italic: true,
						},
					},
					{
						Type:      notion.RichTextTypeText,
						Text:      &notion.Text{Content: "\nlorem ipsum dolor sit amet."},
						PlainText: "\nlorem ipsum dolor sit amet.",
					},
				},
			},
		}

		result, err := Convert(markdownText)

		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("can convert markdown heading and bulleted list to notion blocks", func(t *testing.T) {
		markdownText := `# H1 Example

- Item 1
- Item 2
`

		expected := []notion.Block{
			notion.Heading1Block{
				RichText: []notion.RichText{
					{
						Type:      notion.RichTextTypeText,
						Text:      &notion.Text{Content: "H1 Example"},
						PlainText: "H1 Example",
					},
				},
			},
			notion.BulletedListItemBlock{
				RichText: []notion.RichText{
					{
						Type:      notion.RichTextTypeText,
						Text:      &notion.Text{Content: "Item 1"},
						PlainText: "Item 1",
					},
				},
			},
			notion.BulletedListItemBlock{
				RichText: []notion.RichText{
					{
						Type:      notion.RichTextTypeText,
						Text:      &notion.Text{Content: "Item 2"},
						PlainText: "Item 2",
					},
				},
			},
		}

		result, err := Convert(markdownText)

		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("can convert lists with links and styling", func(t *testing.T) {
		markdownText := `
- [google](https://google.com)
- *GitHub*
`
		expected := []notion.Block{
			notion.BulletedListItemBlock{
				RichText: []notion.RichText{
					{
						Type:      notion.RichTextTypeText,
						PlainText: "google",
						Text: &notion.Text{
							Content: "google",
							Link: &notion.Link{
								URL: "https://google.com",
							},
						},
					},
				},
			},
			notion.BulletedListItemBlock{
				RichText: []notion.RichText{
					{
						Type:      notion.RichTextTypeText,
						PlainText: "GitHub",
						Text: &notion.Text{
							Content: "GitHub",
						},
						Annotations: &notion.Annotations{
							Italic: true,
						},
					},
				},
			},
		}

		result, err := Convert(markdownText)
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})
	
	t.Run("can handle bullet list with relative path links", func(t *testing.T) {
		markdownText := `# How and Why Alpine Linux is used

- [Building](./building.md)
`
		expected := []notion.Block{
			notion.Heading1Block{
				RichText: []notion.RichText{
					{
						Type:      notion.RichTextTypeText,
						Text:      &notion.Text{Content: "How and Why Alpine Linux is used"},
						PlainText: "How and Why Alpine Linux is used",
					},
				},
			},
			notion.BulletedListItemBlock{
				RichText: []notion.RichText{
					{
						Type:      notion.RichTextTypeText,
						Text:      &notion.Text{Content: "Building"},
						PlainText: "Building",
					},
				},
			},
		}

		result, err := Convert(markdownText)
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})
}

func ExampleConvert() {
	markdown := `# H1 Example

- Item 1
`

	blocks, err := Convert(markdown)
	if err != nil {
		panic(err)
	}

	result, _ := json.MarshalIndent(blocks, "", "  ")
	fmt.Println(string(result))

	// Output:
	// [
	//   {
	//     "heading_1": {
	//       "rich_text": [
	//         {
	//           "type": "text",
	//           "plain_text": "H1 Example",
	//           "text": {
	//             "content": "H1 Example"
	//           }
	//         }
	//       ],
	//       "is_toggleable": false
	//     }
	//   },
	//   {
	//     "bulleted_list_item": {
	//       "rich_text": [
	//         {
	//           "type": "text",
	//           "plain_text": "Item 1",
	//           "text": {
	//             "content": "Item 1"
	//           }
	//         }
	//       ]
	//     }
	//   }
	// ]
}
