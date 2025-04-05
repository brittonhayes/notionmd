package converter

import (
	"errors"

	"github.com/brittonhayes/notionmd/chunk"
	"github.com/dstotijn/go-notion"
	"github.com/gomarkdown/markdown/ast"
)

var (
	ErrUnsupportedListItemType = errors.New("unsupported list item type")
)

func isList(node ast.Node) bool {
	_, ok := node.(*ast.List)
	return ok
}

func isListItem(node ast.Node) bool {
	_, ok := node.(*ast.ListItem)

	return ok
}

// convertList converts an *ast.List node to a notion.BulletedListItemBlock.
// It iterates over the list items and their children, extracting the text content
// and creating a notion.RichText slice. The resulting notion.BulletedListItemBlock
// contains the extracted text as rich text.
func convertList(node *ast.List) []notion.Block {
	if node == nil {
		return nil
	}

	var items []notion.Block
	for _, listItem := range node.GetChildren() {
		item := convertListItem(listItem)
		items = append(items, item)
	}

	return items
}

func listItemContent(node ast.Node) []notion.RichText {
	if node == nil {
		return nil
	}

	var richText []notion.RichText
	ast.WalkFunc(node, func(n ast.Node, entering bool) ast.WalkStatus {
		if !entering {
			return ast.GoToNext
		}

		if isLink(n) {
			linkBlock := convertLinkToTextBlock(n.(*ast.Link))
			if linkBlock != nil {
				richText = append(richText, linkBlock...)
			}
			return ast.SkipChildren
		}

		if isStyledText(n) {
			styledTextBlock := convertStyledTextToBlock(n)
			if styledTextBlock != nil {
				richText = append(richText, styledTextBlock...)
			}
			return ast.SkipChildren
		}

		content := string(node.AsContainer().GetChildren()[0].AsContainer().GetChildren()[0].AsLeaf().Literal)
		if content != "" {
			richText = append(richText, chunk.RichText(content, nil)...)
			return ast.SkipChildren
		}

		return ast.GoToNext
	})

	return richText
}

func convertListItem(listItem ast.Node) notion.Block {
	if listItem == nil {
		return nil
	}

	// Check if the list item contains text
	if !isListItem(listItem) {
		return nil
	}

	// Extract the text content
	content := listItemContent(listItem)

	// add support for ordered lists
	if listItem.(*ast.ListItem).ListFlags&ast.ListTypeOrdered != 0 {
		return notion.NumberedListItemBlock{
			RichText: content,
		}
	}

	return notion.BulletedListItemBlock{
		RichText: content,
	}
}
