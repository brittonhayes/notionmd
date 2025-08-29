package converter

import (
	"errors"
	"net/url"

	"github.com/brittonhayes/notionmd/chunk"
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
// If the URL is invalid (like a relative path), it falls back to returning just the text content
// without the link to prevent empty list items that would fail Notion API validation.
func convertLinkToTextBlock(node *ast.Link) []notion.RichText {
	if node == nil {
		return nil
	}

	url := extractURL(node)
	title := extractTitle(node)
	
	ok := isValidURL(url)
	if !ok {
		// Fallback to plain text when URL validation fails
		// This prevents empty list items when using relative paths
		return chunk.RichText(title, nil)
	}

	return chunk.RichTextWithLink(title, url)
}
