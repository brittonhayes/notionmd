package notionmd

import (
	"github.com/brittonhayes/notionmd/internal/converter"
	"github.com/dstotijn/go-notion"
)

// Convert takes a markdown document as text, parses it into an AST node,
// and iterates over the tree with the convertNode function, converting each
// of the nodes to Notion blocks.
func Convert(markdown string) ([]notion.Block, error) {
	return converter.Convert(markdown)
}

// ConvertToJSON takes a markdown document and returns JSON-compatible maps.
// This allows users to work with the data without external dependencies.
func ConvertToJSON(markdown string) ([]map[string]any, error) {
	return converter.ConvertToJSON(markdown)
}
