package converter

import (
	"fmt"

	"github.com/dstotijn/go-notion"
	"github.com/gomarkdown/markdown/ast"
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
func convertList(node *ast.List) ([]notion.Block, error) {
	if node == nil {
		return nil, nil
	}

	var items []notion.Block
	for _, listItem := range node.GetChildren() {
		item, err := convertListItem(listItem)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func listItemContent(node ast.Node) string {
	if node == nil {
		return ""
	}

	return string(node.AsContainer().GetChildren()[0].AsContainer().GetChildren()[0].AsLeaf().Literal)
}

func convertListItem(listItem ast.Node) (notion.Block, error) {
	if listItem == nil {
		return nil, nil
	}

	// Check if the list item contains text
	if !isListItem(listItem) {
		return nil, fmt.Errorf("unsupported list item type: %T", listItem)
	}

	// Extract the text content
	text := listItemContent(listItem)

	// Create a notion.Text object with the extracted text
	richText := notion.BulletedListItemBlock{
		RichText: []notion.RichText{
			{
				Type: notion.RichTextTypeText,
				Text: &notion.Text{
					Content: text,
				},
				PlainText: text,
			},
		},
	}

	return richText, nil
}
