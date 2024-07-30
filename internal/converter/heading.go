package converter

import (
	"github.com/brittonhayes/notionmd/internal/chunk"
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
	if node.GetChildren() == nil {
		return nil, nil
	}

	if node.Level == 1 {
		return notion.Heading1Block{
			RichText: chunk.RichText(string(node.Children[0].AsLeaf().Literal)),
		}, nil
	}

	if node.Level == 2 {
		return notion.Heading2Block{
			RichText: chunk.RichText(string(node.Children[0].AsLeaf().Literal)),
		}, nil
	}

	if node.Level == 3 {
		return notion.Heading3Block{
			RichText: chunk.RichText(string(node.Children[0].AsLeaf().Literal)),
		}, nil
	}

	return nil, nil
}
