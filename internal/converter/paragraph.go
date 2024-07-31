package converter

import (
	"github.com/dstotijn/go-notion"
	"github.com/gomarkdown/markdown/ast"
)

// TODO add support for bold text
// TODO add support for italic text
// TODO add support for underline text

func isParagraph(node ast.Node) bool {
	_, ok := node.(*ast.Paragraph)
	return ok
}

// convertParagraph converts an AST paragraph node to a Notion paragraph block.
// It takes a pointer to an ast.Paragraph node and returns a notion.ParagraphBlock and an error.
func convertParagraph(node *ast.Paragraph) (*notion.ParagraphBlock, error) {
	if node.Children == nil || node == nil {
		return nil, nil
	}

	blocks, err := convertChildNodesToRichText(node)
	if err != nil {
		return nil, err
	}

	if blocks == nil {
		return nil, nil
	}

	return &notion.ParagraphBlock{
		RichText: blocks,
	}, nil
}
