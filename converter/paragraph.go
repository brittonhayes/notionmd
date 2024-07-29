package converter

import (
	"github.com/dstotijn/go-notion"
	"github.com/gomarkdown/markdown/ast"
)

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

			if (linkBlock != notion.RichText{}) {
				blocks = append(blocks, linkBlock)
			}

		} else {
			content := string(child.AsLeaf().Literal)
			if content != "" {
				blocks = append(blocks, notion.RichText{
					Type: notion.RichTextTypeText,
					Text: &notion.Text{
						Content: string(child.AsLeaf().Literal),
					},
					PlainText: string(child.AsLeaf().Literal),
				})
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
