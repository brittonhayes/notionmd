package converter

import (
	"github.com/brittonhayes/notionmd/chunk"
	"github.com/dstotijn/go-notion"
	"github.com/gomarkdown/markdown/ast"
)

func isStyledText(node ast.Node) bool {
	_, isEmph := node.(*ast.Emph)
	_, isStrong := node.(*ast.Strong)

	return isEmph || isStrong
}

func convertStyledTextToBlock(node ast.Node) []notion.RichText {
	if node == nil {
		return nil
	}

	var annotations notion.Annotations
	switch node.(type) {
	case *ast.Emph:
		annotations.Italic = true
	case *ast.Strong:
		annotations.Bold = true
	}

	var blocks []notion.RichText
	for _, child := range node.GetChildren() {
		if child.AsLeaf() == nil {
			continue
		}

		content := string(child.AsLeaf().Literal)
		blocks = append(blocks, chunk.RichText(content, &annotations)...)
	}

	if blocks == nil {
		return nil
	}

	return blocks
}
