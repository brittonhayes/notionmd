package converter

import (
	"testing"

	"github.com/dstotijn/go-notion"
	"github.com/gomarkdown/markdown/ast"
	"github.com/stretchr/testify/assert"
)

func TestConvertBlockquote(t *testing.T) {
	t.Run("can convert blockquote with rich text", func(t *testing.T) {
		childNode := &ast.Paragraph{
			Container: ast.Container{
				Children: []ast.Node{
					&ast.Leaf{
						Content: []byte("This is a blockquote."),
						Literal: []byte("This is a blockquote."),
					},
				},
			},
		}

		node := &ast.BlockQuote{
			Container: ast.Container{
				Children: []ast.Node{childNode},
			},
		}

		expected := &notion.QuoteBlock{
			RichText: []notion.RichText{
				{
					Type:      notion.RichTextTypeText,
					Text:      &notion.Text{Content: "This is a blockquote."},
					PlainText: "This is a blockquote.",
				},
			},
		}

		result := convertBlockquote(node)

		assert.Equal(t, expected, result)
	})

	t.Run("handles empty blockquote", func(t *testing.T) {
		node := &ast.BlockQuote{}

		result := convertBlockquote(node)

		assert.Nil(t, result)
	})

	t.Run("handles nil blockquote", func(t *testing.T) {
		var node *ast.BlockQuote

		result := convertBlockquote(node)

		assert.Nil(t, result)
	})
}
