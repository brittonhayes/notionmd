package converter

import (
	"testing"

	"github.com/dstotijn/go-notion"
	"github.com/gomarkdown/markdown/ast"
	"github.com/stretchr/testify/assert"
)

func TestIsParagraph(t *testing.T) {
	t.Run("is paragraph", func(t *testing.T) {
		node := &ast.Paragraph{}
		assert.True(t, isParagraph(node))
	})

	t.Run("is not paragraph", func(t *testing.T) {
		node := &ast.Heading{}
		assert.False(t, isParagraph(node))
	})
}

func TestConvertParagraph(t *testing.T) {
	t.Run("can convert simple paragraph", func(t *testing.T) {
		input := "This is a simple paragraph."
		expected := notion.ParagraphBlock{
			RichText: []notion.RichText{
				{
					Type:      notion.RichTextTypeText,
					Text:      &notion.Text{Content: input},
					PlainText: input,
				},
			},
		}

		result := testConvertParagraph(t, input)
		assert.Len(t, result.RichText, 1)
		assert.Equal(t, expected.RichText[0].PlainText, result.RichText[0].PlainText)
		assert.Equal(t, expected.RichText[0].Text.Content, result.RichText[0].Text.Content)
	})

	t.Run("handles empty paragraph", func(t *testing.T) {
		input := ""
		expected := &notion.ParagraphBlock{}

		result := testConvertParagraph(t, input)
		assert.IsType(t, expected, result)
		assert.False(t, result.HasChildren())
	})

	t.Run("can convert paragraph with multiple lines of text", func(t *testing.T) {
		input := "This is a paragraph.\nThis is another line."
		expected := notion.ParagraphBlock{
			RichText: []notion.RichText{
				{
					Type:      notion.RichTextTypeText,
					Text:      &notion.Text{Content: input},
					PlainText: input,
				},
			},
		}

		result := testConvertParagraph(t, input)
		// Assert the content is correct
		assert.Len(t, result.RichText, 1)
		assert.Equal(t, expected.RichText[0].PlainText, result.RichText[0].PlainText)
		assert.Equal(t, expected.RichText[0].Text.Content, result.RichText[0].Text.Content)
	})
}

func testConvertParagraph(t *testing.T, input string) *notion.ParagraphBlock {
	// Convert the paragraph
	result, err := convertParagraph(&ast.Paragraph{
		Container: ast.Container{
			Children: []ast.Node{
				&ast.Leaf{
					Content: []byte(input),
					Literal: []byte(input),
				},
			},
		},
	})

	// Assert no error occurred
	assert.NoError(t, err)

	// Assert the result is of type ParagraphBlock
	assert.IsType(t, &notion.ParagraphBlock{}, result, "Result should be of type ParagraphBlock")

	return result
}
