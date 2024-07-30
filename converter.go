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
