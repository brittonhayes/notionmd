package converter

import (
	"github.com/brittonhayes/notionmd/internal/chunk"
	"github.com/dstotijn/go-notion"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"
)

// Convert takes a markdown document as text, parses it into an AST node,
// and iterates over the tree with the convertNode function, converting each
// of the nodes to Notion blocks.
func Convert(markdown string) ([]notion.Block, error) {

	// Parse the markdown document into an AST node
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	p := parser.NewWithExtensions(extensions)
	document := p.Parse([]byte(markdown))

	var blocks []notion.Block

	ast.WalkFunc(document, func(node ast.Node, entering bool) ast.WalkStatus {
		if !entering {
			return ast.GoToNext
		}

		if isHeading(node) {
			block, err := convertHeading(node.(*ast.Heading))
			if err != nil {
				return ast.Terminate
			}
			blocks = append(blocks, block)
			return ast.SkipChildren
		}

		if isList(node) {
			list, err := convertList(node.(*ast.List))
			if err != nil {
				return ast.Terminate
			}
			blocks = append(blocks, list...)
			return ast.SkipChildren
		}

		if isParagraph(node) {
			block, err := convertParagraph(node.(*ast.Paragraph))
			if err != nil {
				return ast.Terminate
			}
			if block != nil {
				blocks = append(blocks, block)
			}

			return ast.SkipChildren
		}

		if isCodeBlock(node) {
			codeBlock, err := convertCodeBlock(node.(*ast.CodeBlock))
			if err != nil {
				return ast.Terminate
			}
			if codeBlock != nil {
				blocks = append(blocks, codeBlock)
			}

			return ast.SkipChildren
		}

		return ast.GoToNext
	})

	return blocks, nil
}

func convertChildNodesToRichText(node ast.Node) ([]notion.RichText, error) {
	var blocks []notion.RichText

	for _, child := range node.GetChildren() {
		if isLink(child) {
			linkBlock, err := convertLinkToTextBlock(child.(*ast.Link))
			if err != nil {
				return nil, err
			}

			if linkBlock != nil {
				blocks = append(blocks, linkBlock...)
			}

		} else if isStyledText(child) {
			styledBlock := convertStyledTextToBlock(child)
			if styledBlock != nil {
				blocks = append(blocks, styledBlock...)
			}
		} else {
			content := string(child.AsLeaf().Literal)
			if content != "" {
				blocks = append(blocks, chunk.RichText(content, nil)...)
			}
		}
	}

	if blocks == nil {
		return nil, nil
	}

	return blocks, nil
}
