package converter

import (
	"log/slog"

	"github.com/dstotijn/go-notion"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"
)

// Convert takes a markdown document as text, parses it into an AST node,
// and iterates over the tree with the convertNode function, converting each
// of the nodes to Notion blocks.
func Convert(s string) ([]notion.Block, error) {

	// Parse the markdown document into an AST node
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse([]byte(s))

	var blocks []notion.Block
	ast.WalkFunc(doc, func(node ast.Node, entering bool) ast.WalkStatus {

		if isHeading(node) {
			block, err := convertHeading(node.(*ast.Heading))
			if err != nil {
				slog.Error("error converting heading", "error", err)
				return ast.Terminate
			}
			blocks = append(blocks, block)
			return ast.GoToNext
		}

		if isParagraph(node) {
			block, err := convertParagraph(node.(*ast.Paragraph))
			if err != nil {
				slog.Error("error converting paragraph", "error", err)
				return ast.Terminate
			}
			blocks = append(blocks, block)
			return ast.GoToNext
		}

		if node.GetChildren() == nil {
			return ast.Terminate
		}

		return ast.GoToNext
	})

	return blocks, nil
}
