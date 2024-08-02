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

	valid := validateLanguage(string(node.Info))
	if !valid {
		return ""
	}

	return string(node.Info)
}

// convertCodeBlock converts an AST code block node to a Notion code block.
// It takes a pointer to an ast.CodeBlock node and returns a notion.CodeBlock and an error.
func convertCodeBlock(node *ast.CodeBlock) *notion.CodeBlock {
	if node == nil || node.Literal == nil {
		return nil
	}

	content := string(node.Literal)
	if content == "" {
		return nil
	}

	result := &notion.CodeBlock{
		RichText: chunk.RichText(content, nil),
	}

	language := extractLanguage(node)
	if language != "" {
		result.Language = &language
	}

	return result
}

// validateLanguage checks if the code language is a valid option
// for Notion's code block.
//
// https://developers.notion.com/reference/block#code
func validateLanguage(language string) bool {
	validLanguages := map[string]bool{
		"abap":          true,
		"arduino":       true,
		"bash":          true,
		"basic":         true,
		"c":             true,
		"clojure":       true,
		"coffeescript":  true,
		"c++":           true,
		"c#":            true,
		"css":           true,
		"dart":          true,
		"diff":          true,
		"docker":        true,
		"elixir":        true,
		"elm":           true,
		"erlang":        true,
		"flow":          true,
		"fortran":       true,
		"f#":            true,
		"gherkin":       true,
		"glsl":          true,
		"go":            true,
		"graphql":       true,
		"groovy":        true,
		"haskell":       true,
		"html":          true,
		"java":          true,
		"javascript":    true,
		"json":          true,
		"julia":         true,
		"kotlin":        true,
		"latex":         true,
		"less":          true,
		"lisp":          true,
		"livescript":    true,
		"lua":           true,
		"makefile":      true,
		"markdown":      true,
		"markup":        true,
		"matlab":        true,
		"mermaid":       true,
		"nix":           true,
		"objective-c":   true,
		"ocaml":         true,
		"pascal":        true,
		"perl":          true,
		"php":           true,
		"plain text":    true,
		"powershell":    true,
		"prolog":        true,
		"protobuf":      true,
		"python":        true,
		"r":             true,
		"reason":        true,
		"ruby":          true,
		"rust":          true,
		"sass":          true,
		"scala":         true,
		"scheme":        true,
		"scss":          true,
		"shell":         true,
		"sql":           true,
		"swift":         true,
		"typescript":    true,
		"vb.net":        true,
		"verilog":       true,
		"vhdl":          true,
		"visual basic":  true,
		"webassembly":   true,
		"xml":           true,
		"yaml":          true,
		"java/c/c++/c#": true,
	}

	_, ok := validLanguages[language]
	return ok
}
