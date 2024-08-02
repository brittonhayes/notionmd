package converter

import (
	"errors"
	"net/url"

	"github.com/brittonhayes/notionmd/internal/chunk"
	"github.com/dstotijn/go-notion"
	"github.com/gomarkdown/markdown/ast"
)

var (
	ErrInvalidURL       = errors.New("invalid URL, must be a valid URI")
	ErrExpectedLinkNode = errors.New("expected *ast.Link node, got nil")
)

func isValidURL(uri string) bool {
	_, err := url.ParseRequestURI(uri)
	return err == nil
}

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

// convertLinkToTextBlock converts an AST link node to a Notion text block.
// It takes a pointer to an ast.Link node and returns a Notion text block.
func convertLinkToTextBlock(node *ast.Link) []notion.RichText {
	if node == nil {
		return nil
	}

	ok := isValidURL(extractURL(node))
	if !ok {
		return nil
	}

	return chunk.RichTextWithLink(extractTitle(node), extractURL(node))
}
