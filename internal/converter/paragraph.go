package converter

import (
	"github.com/dstotijn/go-notion"
	"github.com/gomarkdown/markdown/ast"
)

func isParagraph(node ast.Node) bool {
	_, ok := node.(*ast.Paragraph)
	return ok
}

// convertParagraph converts an AST paragraph node to a Notion paragraph block.
// It takes a pointer to an ast.Paragraph node and returns a notion.ParagraphBlock and an error.
func convertParagraph(node *ast.Paragraph) *notion.ParagraphBlock {
	if node.Children == nil || node == nil {
		return nil
	}

	blocks := convertChildNodesToRichText(node)
	if blocks == nil {
		return nil
	}

	return &notion.ParagraphBlock{
		RichText: blocks,
	}
}
