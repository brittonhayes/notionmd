package converter

import (
	"github.com/dstotijn/go-notion"
	"github.com/gomarkdown/markdown/ast"
)

func isList(node ast.Node) bool {
	_, ok := node.(*ast.List)
	return ok
}

// convertList converts an *ast.List node to a notion.BulletedListItemBlock.
// It iterates over the list items and their children, extracting the text content
// and creating a notion.RichText slice. The resulting notion.BulletedListItemBlock
// contains the extracted text as rich text.
func convertList(node *ast.List) (notion.Block, error) {
	if node == nil {
		return notion.BulletedListItemBlock{}, nil
	}

	var items []notion.RichText
	// TODO implement list conversion logic
	return notion.BulletedListItemBlock{
		RichText: items,
	}, nil
}
