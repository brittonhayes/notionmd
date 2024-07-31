package converter

import (
	"github.com/brittonhayes/notionmd/internal/chunk"
	"github.com/dstotijn/go-notion"
	"github.com/gomarkdown/markdown/ast"
)

// isCodeBlock checks if a node is a code block.
func isCodeBlock(node ast.Node) bool {
	_, ok := node.(*ast.CodeBlock)
	return ok
}

// extractLanguage extracts the language from a code block node.
func extractLanguage(node *ast.CodeBlock) string {
	if node == nil {
		return ""
	}
	return string(node.Info)
}

// convertCodeBlock converts an AST code block node to a Notion code block.
// It takes a pointer to an ast.CodeBlock node and returns a notion.CodeBlock and an error.
func convertCodeBlock(node *ast.CodeBlock) (*notion.CodeBlock, error) {
	if node == nil || node.Literal == nil {
		return nil, nil
	}

	content := string(node.Literal)
	if content == "" {
		return nil, nil
	}

	result := &notion.CodeBlock{
		RichText: chunk.RichText(content, nil),
	}

	language := extractLanguage(node)
	if language != "" {
		result.Language = &language
	}

	return result, nil
}
