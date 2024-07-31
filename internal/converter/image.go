package converter

import "github.com/gomarkdown/markdown/ast"

func isImage(node ast.Node) bool {
	_, ok := node.(*ast.Image)
	return ok
}
