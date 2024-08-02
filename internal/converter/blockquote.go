package converter

import (
	"github.com/dstotijn/go-notion"
	"github.com/gomarkdown/markdown/ast"
)

func isBlockquote(node ast.Node) bool {
	_, ok := node.(*ast.BlockQuote)
	return ok
}

func convertBlockquote(node *ast.BlockQuote) *notion.QuoteBlock {
	if node == nil {
		return nil
	}

	var blocks []notion.RichText
	for _, child := range node.GetChildren() {
		childBlocks := convertChildNodesToRichText(child)
		if childBlocks != nil {
			blocks = append(blocks, childBlocks...)
		}
	}

	if blocks == nil {
		return nil
	}

	return &notion.QuoteBlock{
		RichText: blocks,
	}
}
