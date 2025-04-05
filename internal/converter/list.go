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
		if item != nil {
			items = append(items, item)
		}
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

		if n.AsLeaf() != nil {
			content := string(n.AsLeaf().Literal)
			if content != "" {
				richText = append(richText, chunk.RichText(content, nil)...)
			}
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

	// Get the list item's children
	children := listItem.GetChildren()
	if len(children) == 0 {
		return nil
	}

	// The first child should be a paragraph containing the text
	paragraph, ok := children[0].(*ast.Paragraph)
	if !ok {
		return nil
	}

	// Extract the text content from the paragraph
	content := listItemContent(paragraph)

	// Process any remaining children (which could be nested lists)
	var nestedBlocks []notion.Block
	for _, child := range children[1:] {
		if isList(child) {
			nestedBlocks = append(nestedBlocks, convertList(child.(*ast.List))...)
		}
	}

	// Create the appropriate list item block
	if listItem.(*ast.ListItem).ListFlags&ast.ListTypeOrdered != 0 {
		return notion.NumberedListItemBlock{
			RichText: content,
			Children: nestedBlocks,
		}
	}

	return notion.BulletedListItemBlock{
		RichText: content,
		Children: nestedBlocks,
	}
}
