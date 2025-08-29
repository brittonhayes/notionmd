package converter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertToJSON(t *testing.T) {
	t.Run("can convert simple markdown to JSON", func(t *testing.T) {
		markdown := `# Hello World

This is a paragraph.

- List item 1
- List item 2`

		jsonBlocks, err := ConvertToJSON(markdown)
		assert.NoError(t, err)
		assert.Len(t, jsonBlocks, 4)

		// Check first block (heading)
		firstBlock := jsonBlocks[0]

		// The notion blocks don't have a "type" field, they have the type as a key
		// Check if it has a heading_1 key
		_, hasHeading1 := firstBlock["heading_1"]
		assert.True(t, hasHeading1, "First block should have heading_1 key")

		heading1, ok := firstBlock["heading_1"].(map[string]any)
		assert.True(t, ok)

		richText, ok := heading1["rich_text"].([]any)
		assert.True(t, ok)
		assert.Len(t, richText, 1)

		textBlock := richText[0].(map[string]any)
		assert.Equal(t, "text", textBlock["type"])

		text, ok := textBlock["text"].(map[string]any)
		assert.True(t, ok)
		assert.Equal(t, "Hello World", text["content"])
	})

	t.Run("can convert markdown with blockquote to JSON", func(t *testing.T) {
		markdown := `> This is a blockquote`

		jsonBlocks, err := ConvertToJSON(markdown)
		assert.NoError(t, err)
		assert.Len(t, jsonBlocks, 1)

		firstBlock := jsonBlocks[0]

		// Check if it has a quote key
		_, hasQuote := firstBlock["quote"]
		assert.True(t, hasQuote, "First block should have quote key")

		quote, ok := firstBlock["quote"].(map[string]any)
		assert.True(t, ok)

		richText, ok := quote["rich_text"].([]any)
		assert.True(t, ok)
		assert.Len(t, richText, 1)
	})

	t.Run("can convert markdown with code block to JSON", func(t *testing.T) {
		markdown := "```go\nfmt.Println(\"Hello\")\n```"

		jsonBlocks, err := ConvertToJSON(markdown)
		assert.NoError(t, err)
		assert.Len(t, jsonBlocks, 1)

		firstBlock := jsonBlocks[0]

		// Check if it has a code key
		_, hasCode := firstBlock["code"]
		assert.True(t, hasCode, "First block should have code key")

		code, ok := firstBlock["code"].(map[string]any)
		assert.True(t, ok)
		assert.Equal(t, "go", code["language"])
	})
}
