package converter

import (
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

This is a paragraph.`

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
			notion.ParagraphBlock{
				RichText: []notion.RichText{
					{
						Type:      notion.RichTextTypeText,
						Text:      &notion.Text{Content: "This is a paragraph."},
						PlainText: "This is a paragraph.",
					},
				},
			},
		}

		result, err := Convert(markdownText)
		assert.NoError(t, err)
		assert.NotNil(t, result)

		assert.Equal(t, expected, result)
		assert.Len(t, result, 2)
	})
}
