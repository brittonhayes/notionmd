package converter

import (
	"github.com/dstotijn/go-notion"
	"github.com/gomarkdown/markdown/ast"
)

func isLink(node ast.Node) bool {
	_, ok := node.(*ast.Link)
	return ok
}

// convertLink converts an AST link node to a Notion block.
// It takes a pointer to an ast.Link node and returns a Notion block and an error.
func convertLink(node ast.Node) (notion.Block, error) {
	if node == nil {
		return notion.ParagraphBlock{}, nil
	}

	return notion.ParagraphBlock{
		RichText: []notion.RichText{
			{
				Type: notion.RichTextTypeText,
				Text: &notion.Text{
					Link: &notion.Link{
						URL: string(node.AsContainer().Content),
					},
				},
				PlainText: string(node.AsContainer().Content),
			}},
	}, nil
}
