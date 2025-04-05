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

		blocks := convertList(listNode)
		assert.Len(t, blocks, 2, "Expected convertList to return 1 block")
	})

	t.Run("handles empty list", func(t *testing.T) {
		listNode := &ast.List{}

		blocks := convertList(listNode)
		assert.Empty(t, blocks, "Expected convertList to return no blocks")
	})

	t.Run("can convert nested list", func(t *testing.T) {
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
												Literal: []byte("Parent Item"),
											},
										},
									},
								},
								&ast.List{
									Container: ast.Container{
										Children: []ast.Node{
											&ast.ListItem{
												Container: ast.Container{
													Children: []ast.Node{
														&ast.Paragraph{
															Container: ast.Container{
																Children: []ast.Node{
																	&ast.Leaf{
																		Literal: []byte("Child Item 1"),
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
																		Literal: []byte("Child Item 2"),
																	},
																},
															},
														},
													},
												},
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

		blocks := convertList(listNode)
		assert.Len(t, blocks, 1, "Expected convertList to return 1 parent block")

		parentBlock, ok := blocks[0].(notion.BulletedListItemBlock)
		assert.True(t, ok, "Expected first block to be a bulleted list item")
		assert.Equal(t, "Parent Item", parentBlock.RichText[0].PlainText, "Expected parent item text to be 'Parent Item'")
		assert.NotNil(t, parentBlock.Children, "Expected parent block to have children")
		assert.Len(t, parentBlock.Children, 2, "Expected parent block to have 2 child blocks")

		child1, ok := parentBlock.Children[0].(notion.BulletedListItemBlock)
		assert.True(t, ok, "Expected first child to be a bulleted list item")
		assert.Equal(t, "Child Item 1", child1.RichText[0].PlainText, "Expected first child text to be 'Child Item 1'")

		child2, ok := parentBlock.Children[1].(notion.BulletedListItemBlock)
		assert.True(t, ok, "Expected second child to be a bulleted list item")
		assert.Equal(t, "Child Item 2", child2.RichText[0].PlainText, "Expected second child text to be 'Child Item 2'")
	})

	t.Run("can convert nested list from markdown", func(t *testing.T) {
		// This test validates the full conversion from markdown string to Notion blocks
		markdownText := `- list
  - nested list`

		expected := []notion.Block{
			notion.BulletedListItemBlock{
				RichText: []notion.RichText{
					{
						Type:      notion.RichTextTypeText,
						Text:      &notion.Text{Content: "list"},
						PlainText: "list",
					},
				},
				Children: []notion.Block{
					notion.BulletedListItemBlock{
						RichText: []notion.RichText{
							{
								Type:      notion.RichTextTypeText,
								Text:      &notion.Text{Content: "nested list"},
								PlainText: "nested list",
							},
						},
					},
				},
			},
		}

		result, err := Convert(markdownText)
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
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

		block := convertListItem(listItem)

		bulletedListItem, ok := block.(notion.BulletedListItemBlock)
		assert.True(t, ok, "Expected convertListItem to return a notion.BulletedListItemBlock")
		assert.Len(t, bulletedListItem.RichText, 1, "Expected bulletedListItem to contain 1 rich text item")
		assert.Equal(t, "Item 1", bulletedListItem.RichText[0].PlainText, "Expected rich text content to be 'Item 1'")
	})

	t.Run("handles nil list item", func(t *testing.T) {
		block := convertListItem(nil)
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
