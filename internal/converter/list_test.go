package converter

import (
	"testing"

	"github.com/dstotijn/go-notion"
	"github.com/gomarkdown/markdown/ast"
	"github.com/stretchr/testify/assert"
)

func TestIsList(t *testing.T) {
	listNode := &ast.List{}
	assert.True(t, isList(listNode), "Expected isList to return true for *ast.List node")

	paragraphNode := &ast.Paragraph{}
	assert.False(t, isList(paragraphNode), "Expected isList to return false for *ast.Paragraph node")
}

func TestConvertList(t *testing.T) {

	t.Run("can convert simple list", func(t *testing.T) {
		listNode := &ast.List{
			Container: ast.Container{
				Children: []ast.Node{
					&ast.ListItem{
						ListFlags: ast.ListItemBeginningOfList,
						Container: ast.Container{
							Children: []ast.Node{
								&ast.Paragraph{
									Container: ast.Container{
										Children: []ast.Node{
											&ast.Leaf{
												Literal: []byte("Item 1"),
											},
										},
									},
								},
							},
						},
					},
					&ast.ListItem{
						Container: ast.Container{
							Children: []ast.Node{
								&ast.Paragraph{
									Container: ast.Container{
										Children: []ast.Node{
											&ast.Leaf{
												Literal: []byte("Item 2"),
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}

		blocks, err := convertList(listNode)
		assert.NoError(t, err, "Expected convertList to not return an error")
		assert.Len(t, blocks, 2, "Expected convertList to return 1 block")
	})

	t.Run("handles empty list", func(t *testing.T) {
		listNode := &ast.List{}

		blocks, err := convertList(listNode)
		assert.NoError(t, err, "Expected convertList to not return an error")
		assert.Empty(t, blocks, "Expected convertList to return no blocks")
	})
}

func TestConvertListItem(t *testing.T) {
	t.Run("converts list item with paragraph", func(t *testing.T) {
		listItem := &ast.ListItem{
			ListFlags: ast.ListItemBeginningOfList,
			Container: ast.Container{
				Children: []ast.Node{
					&ast.Paragraph{
						Container: ast.Container{
							Children: []ast.Node{
								&ast.Leaf{
									Literal: []byte("Item 1"),
								},
							},
						},
					},
				},
			},
		}

		block, err := convertListItem(listItem)
		assert.NoError(t, err, "Expected convertListItem to not return an error")

		bulletedListItem, ok := block.(notion.BulletedListItemBlock)
		assert.True(t, ok, "Expected convertListItem to return a notion.BulletedListItemBlock")
		assert.Len(t, bulletedListItem.RichText, 1, "Expected bulletedListItem to contain 1 rich text item")
		assert.Equal(t, "Item 1", bulletedListItem.RichText[0].PlainText, "Expected rich text content to be 'Item 1'")
	})

	t.Run("handles nil list item", func(t *testing.T) {
		block, err := convertListItem(nil)
		assert.NoError(t, err, "Expected convertListItem to not return an error for nil list item")
		assert.Nil(t, block, "Expected block to be nil for nil list item")
	})
}

func TestListItemContent(t *testing.T) {
	t.Run("returns the content of the list item", func(t *testing.T) {
		listItemNode := &ast.ListItem{
			ListFlags: ast.ListItemBeginningOfList,
			Container: ast.Container{
				Children: []ast.Node{
					&ast.Paragraph{
						Container: ast.Container{
							Children: []ast.Node{
								&ast.Leaf{
									Literal: []byte("Item 1"),
								},
							},
						},
					},
				},
			},
		}

		content := listItemContent(listItemNode)
		expect := []notion.RichText{
			{
				Type:      notion.RichTextTypeText,
				PlainText: "Item 1",
				Text: &notion.Text{
					Content: "Item 1",
				},
			},
		}

		assert.Equal(t, expect, content)
	})

	t.Run("handles nil list item content", func(t *testing.T) {
		content := listItemContent(nil)
		assert.Nil(t, content)
	})
}
