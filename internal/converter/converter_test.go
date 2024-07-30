package converter

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/dstotijn/go-notion"
	"github.com/stretchr/testify/assert"
)

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

This is a paragraph
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
						Text:      &notion.Text{Content: "This is a paragraph\nlorem ipsum dolor sit amet."},
						PlainText: "This is a paragraph\nlorem ipsum dolor sit amet.",
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
