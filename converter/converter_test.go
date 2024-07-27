package converter

import (
	"testing"

	"github.com/dstotijn/go-notion"
	"github.com/gomarkdown/markdown/ast"
	"github.com/stretchr/testify/assert"
)

func TestConvertNode(t *testing.T) {

	t.Run("can convert heading node", func(t *testing.T) {
		// Test heading conversion
		heading := &ast.Heading{
			Level: 1,
			Container: ast.Container{
				Children: []ast.Node{
					&ast.Leaf{
						Content: []byte("# Hello, World!"),
						Literal: []byte("Hello, World!"),
					},
				},
			},
		}
		block, err := convertNode(heading)
		assert.NoError(t, err)

		assert.IsType(t, notion.Heading1Block{}, block)
		assert.Equal(t, "Hello, World!", block.(notion.Heading1Block).RichText[0].PlainText)
	})

	t.Run("can convert paragraph node", func(t *testing.T) {
		// Test paragraph conversion
		paragraph := &ast.Paragraph{
			Container: ast.Container{
				Children: []ast.Node{
					&ast.Leaf{
						Content: []byte("Hello, World!"),
						Literal: []byte("Hello, World!"),
					},
				},
			},
		}
		block, err := convertNode(paragraph)
		assert.NoError(t, err)

		assert.IsType(t, &notion.ParagraphBlock{}, block)
		assert.Equal(t, "Hello, World!", block.(*notion.ParagraphBlock).RichText[0].PlainText)
	})

	//TODO Add test for list conversion
	// t.Run("can convert list node", func(t *testing.T) {
	// 	// Test list conversion
	// 	list := &ast.List{}
	// 	block, err := convertNode(list)
	// 	assert.NoError(t, err)

	// 	assert.IsType(t, notion.BulletedListItemBlock{}, block)
	// })

	t.Run("can convert unknown node", func(t *testing.T) {
		// Test unknown node conversion
		unknownNode := &ast.CodeBlock{}
		block, err := convertNode(unknownNode)
		assert.Error(t, err)
		assert.Nil(t, block)
	})
}
