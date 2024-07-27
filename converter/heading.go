package converter

import (
	"github.com/dstotijn/go-notion"
	"github.com/gomarkdown/markdown/ast"
)

const (
	ErrNotHeading = "received node is not a heading, got %T"
)

func isHeading(node ast.Node) bool {
	_, ok := node.(*ast.Heading)
	return ok
}

// convertHeading converts an AST heading node to a Notion block.
//
// It takes a pointer to an ast.Heading node as input and returns a Notion block and an error.
// The function extracts the text content from the heading node and creates a Notion block
// based on the heading level. If the heading level is not 1, 2, or 3, it treats the node as a paragraph.
// The function returns the corresponding Notion block and an error if any.
func convertHeading(node *ast.Heading) (notion.Block, error) {
	var block notion.Block

	if node.Children == nil {
		return nil, nil
	}

	switch node.Level {
	case 1:
		block = notion.Heading1Block{
			RichText: []notion.RichText{
				{
					Type: notion.RichTextTypeText,
					Text: &notion.Text{
						Content: string(node.Children[0].AsLeaf().Literal),
					},
					PlainText: string(node.Children[0].AsLeaf().Literal),
				},
			},
		}
	case 2:
		block = notion.Heading2Block{
			RichText: []notion.RichText{
				{
					Type: notion.RichTextTypeText,
					Text: &notion.Text{
						Content: string(node.Children[0].AsLeaf().Literal),
					},
					PlainText: string(node.Children[0].AsLeaf().Literal),
				},
			},
		}
	case 3:
		block = notion.Heading3Block{
			RichText: []notion.RichText{
				{
					Type: notion.RichTextTypeText,
					Text: &notion.Text{
						Content: string(node.Children[0].AsLeaf().Literal),
					},
					PlainText: string(node.Children[0].AsLeaf().Literal),
				},
			},
		}
	}

	return block, nil
}
