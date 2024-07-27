package converter

import (
	"fmt"

	"github.com/dstotijn/go-notion"
	"github.com/gomarkdown/markdown/ast"
)

// Converter is an interface that defines the contract for converting markdown to Notion blocks.
type Converter interface {
	Convert(markdown string) ([]notion.Block, error)
}

// convertNode converts an AST node to a Notion block.
// It checks the type of the node and calls the corresponding conversion function.
// If the node type is unsupported, it returns nil and no error.
func convertNode(node ast.Node) (notion.Block, error) {
	if isHeading(node) {
		return convertHeading(node.(*ast.Heading))
	}

	if isParagraph(node) {
		return convertParagraph(node.(*ast.Paragraph))
	}

	// if isList(node) {
	// 	return convertList(node.(*ast.List))
	// }

	// if isLink(node) {
	// 	return convertLink(node.(*ast.Link))
	// }

	return nil, fmt.Errorf("unsupported node type: %T", node)
}
