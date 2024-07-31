package converter

import (
	"testing"

	"github.com/dstotijn/go-notion"
	"github.com/gomarkdown/markdown/ast"
	"github.com/stretchr/testify/assert"
)

func TestConvertCodeBlock(t *testing.T) {
	t.Run("converts code block with content and language", func(t *testing.T) {
		codeBlockNode := &ast.CodeBlock{
			Leaf: ast.Leaf{
				Literal: []byte("fmt.Println(\"Hello, World!\")"),
			},
			Info: []byte("go"),
		}

		language := extractLanguage(codeBlockNode)
		expected := notion.CodeBlock{
			RichText: []notion.RichText{
				{
					Type:      notion.RichTextTypeText,
					Text:      &notion.Text{Content: "fmt.Println(\"Hello, World!\")"},
					PlainText: "fmt.Println(\"Hello, World!\")",
				},
			},
			Language: &language,
		}

		codeBlock := convertCodeBlock(codeBlockNode)
		assert.Equal(t, expected, *codeBlock, "Expected code block to match")
	})

	t.Run("handles nil code block", func(t *testing.T) {
		codeBlock := convertCodeBlock(nil)
		assert.Nil(t, codeBlock, "Expected code block to be nil for nil code block")
	})

	t.Run("handles empty code block", func(t *testing.T) {
		codeBlockNode := &ast.CodeBlock{}

		codeBlock := convertCodeBlock(codeBlockNode)
		assert.Nil(t, codeBlock, "Expected code block to be nil for empty code block")
	})

	t.Run("handles code block with content but no language", func(t *testing.T) {
		codeBlockNode := &ast.CodeBlock{
			Leaf: ast.Leaf{
				Literal: []byte("fmt.Println(\"Hello, World!\")"),
			},
		}

		expected := notion.CodeBlock{
			RichText: []notion.RichText{
				{
					Type:      notion.RichTextTypeText,
					Text:      &notion.Text{Content: "fmt.Println(\"Hello, World!\")"},
					PlainText: "fmt.Println(\"Hello, World!\")",
				},
			},
		}

		codeBlock := convertCodeBlock(codeBlockNode)
		assert.Equal(t, expected, *codeBlock)
	})
}
