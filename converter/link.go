package converter

import (
	"errors"

	"github.com/dstotijn/go-notion"
	"github.com/gomarkdown/markdown/ast"
)

var (
	ErrExpectedLinkNode = errors.New("expected *ast.Link node, got nil")
)

// TODO add support for images

func isLink(node ast.Node) bool {
	_, ok := node.(*ast.Link)
	return ok
}

func extractURL(node *ast.Link) string {
	return string(node.Destination)
}

func extractTitle(node *ast.Link) string {
	return string(node.AsContainer().GetChildren()[0].AsLeaf().Literal)
}

// convertLink converts an AST link node to a Notion block.
// It takes a pointer to an ast.Link node and returns a Notion block and an error.
func convertLink(node *ast.Link) (*notion.ParagraphBlock, error) {
	if node == nil {
		return nil, ErrExpectedLinkNode
	}

	richText, err := convertLinkToTextBlock(node)
	if err != nil {
		return nil, err
	}

	return &notion.ParagraphBlock{
		RichText: []notion.RichText{richText},
	}, nil
}

// convertLinkToTextBlock converts an AST link node to a Notion text block.
// It takes a pointer to an ast.Link node and returns a Notion text block and an error.
func convertLinkToTextBlock(node *ast.Link) (notion.RichText, error) {
	if node == nil {
		return notion.RichText{}, ErrExpectedLinkNode
	}

	return notion.RichText{
		Type: notion.RichTextTypeText,
		Text: &notion.Text{
			Content: extractTitle(node),
			Link: &notion.Link{
				URL: extractURL(node),
			},
		},
		PlainText: extractTitle(node),
	}, nil
}
