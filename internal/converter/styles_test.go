package converter

import (
	"testing"

	"github.com/dstotijn/go-notion"
	"github.com/gomarkdown/markdown/ast"
	"github.com/stretchr/testify/assert"
)

func TestConvertTextStyles(t *testing.T) {
	// Create a test node
	input := "Hello, World!"
	node := &ast.Emph{
		Container: ast.Container{
			Children: []ast.Node{
				&ast.Leaf{
					Literal: []byte(input),
				},
			},
		},
	}

	expected := []notion.RichText{
		{
			Type: notion.RichTextTypeText,
			Text: &notion.Text{
				Content: "Hello, World!",
			},
			PlainText: "Hello, World!",
			Annotations: &notion.Annotations{
				Italic: true,
			},
		},
	}

	result := convertStyledTextToBlock(node)

	assert.NotNil(t, result)
	assert.Equal(t, expected, result)

	for i, rt := range result {
		expectedRT := expected[i]

		assert.Equal(t, expectedRT.Type, rt.Type)
		assert.Equal(t, expectedRT.Text.Content, rt.Text.Content)
		assert.Equal(t, expectedRT.Annotations.Italic, rt.Annotations.Italic)
		assert.Equal(t, expectedRT.Annotations.Bold, rt.Annotations.Bold)
	}
}
