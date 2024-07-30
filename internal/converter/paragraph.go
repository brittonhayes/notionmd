package converter

import (
	"github.com/brittonhayes/notionmd/internal/chunk"
	"github.com/dstotijn/go-notion"
	"github.com/gomarkdown/markdown/ast"
)

// TODO add support for bold text
// TODO add support for italic text
// TODO add support for underline text

func isParagraph(node ast.Node) bool {
	_, ok := node.(*ast.Paragraph)
	return ok
}

// convertParagraph converts an AST paragraph node to a Notion paragraph block.
// It takes a pointer to an ast.Paragraph node and returns a notion.ParagraphBlock and an error.
func convertParagraph(node *ast.Paragraph) (*notion.ParagraphBlock, error) {
	if node.Children == nil || node == nil {
		return nil, nil
	}

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

		} else {
			content := string(child.AsLeaf().Literal)
			if content != "" {
				blocks = append(blocks, chunk.RichText(content)...)
			}
		}
	}

	if blocks == nil {
		return nil, nil
	}

	return &notion.ParagraphBlock{
		RichText: blocks,
	}, nil
}
