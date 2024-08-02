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
	extensions := parser.CommonExtensions
	p := parser.NewWithExtensions(extensions)
	document := p.Parse([]byte(markdown))

	var blocks []notion.Block

	ast.WalkFunc(document, func(node ast.Node, entering bool) ast.WalkStatus {
		if !entering {
			return ast.GoToNext
		}

		if isImage(node) {
			return ast.GoToNext
		}

		if isList(node) {
			list := convertList(node.(*ast.List))
			blocks = append(blocks, list...)
			return ast.SkipChildren
		}

		if isBlockquote(node) {
			quote := convertBlockquote(node.(*ast.BlockQuote))
			blocks = append(blocks, quote)
			return ast.SkipChildren
		}

		if isHeading(node) {
			block := convertHeading(node.(*ast.Heading))
			blocks = append(blocks, block)
			return ast.GoToNext
		}

		if isParagraph(node) {
			block := convertParagraph(node.(*ast.Paragraph))
			if block != nil {
				blocks = append(blocks, block)
			}

			return ast.SkipChildren
		}

		if isCodeBlock(node) {
			codeBlock := convertCodeBlock(node.(*ast.CodeBlock))
			if codeBlock != nil {
				blocks = append(blocks, codeBlock)
			}

			return ast.SkipChildren
		}

		return ast.GoToNext
	})

	return blocks, nil
}

func convertChildNodesToRichText(node ast.Node) []notion.RichText {
	if node == nil {
		return nil
	}

	var blocks []notion.RichText
	for _, child := range node.GetChildren() {
		if isLink(child) {
			linkBlock := convertLinkToTextBlock(child.(*ast.Link))
			if linkBlock != nil {
				blocks = append(blocks, linkBlock...)
			}
			continue
		}

		if isStyledText(child) {
			styledBlock := convertStyledTextToBlock(child)
			if styledBlock != nil {
				blocks = append(blocks, styledBlock...)
			}
			continue
		}

		value := child.AsLeaf()
		if value == nil {
			continue
		}

		content := string(child.AsLeaf().Literal)
		if content != "" {
			blocks = append(blocks, chunk.RichText(content, nil)...)
		}
	}

	if blocks == nil {
		return nil
	}

	return blocks
}
