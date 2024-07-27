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
func convertParagraph(node *ast.Paragraph) (*notion.ParagraphBlock, error) {
	if node.Children == nil {
		return nil, nil
	}

	var blocks []notion.RichText
	blocks = append(blocks, notion.RichText{
		Type: notion.RichTextTypeText,
		Text: &notion.Text{
			Content: string(node.Children[0].AsLeaf().Literal),
		},
		PlainText: string(node.Children[0].AsLeaf().Literal),
	})

	return &notion.ParagraphBlock{
		RichText: blocks,
	}, nil
}
