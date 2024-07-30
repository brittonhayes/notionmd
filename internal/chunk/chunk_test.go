package chunk

import (
	"testing"

	"github.com/dstotijn/go-notion"
	"github.com/stretchr/testify/assert"
)

func TestRichText(t *testing.T) {

	t.Run("can convert rich text under character limit", func(t *testing.T) {
		content := "Lorem ipsum dolor sit amet, consectetur adipiscing elit."
		expectedBlocks := []notion.RichText{
			{
				Type: notion.RichTextTypeText,
				Text: &notion.Text{
					Content: content,
				},
				PlainText: content,
			},
		}

		blocks := RichText(content)

		assert.Equal(t, len(expectedBlocks), len(blocks), "Expected %d blocks, but got %d", len(expectedBlocks), len(blocks))

		for i, block := range blocks {
			expectedBlock := expectedBlocks[i]

			assert.Equal(t, expectedBlock.Type, block.Type, "Expected block type %s, but got %s", expectedBlock.Type, block.Type)
			assert.Equal(t, expectedBlock.PlainText, block.PlainText, "Expected plain text %s, but got %s", expectedBlock.PlainText, block.PlainText)
			assert.Equal(t, expectedBlock.Text.Content, block.Text.Content, "Expected text content %s, but got %s", expectedBlock.Text.Content, block.Text.Content)
		}
	})

	t.Run("can convert rich text over character limit", func(t *testing.T) {
		// Create a string over 2000 characters
		var longContent string
		for i := 0; i <= CharacterLimit; i++ {
			longContent += "a"
		}

		expectedBlocks := []notion.RichText{
			{
				Type: notion.RichTextTypeText,
				Text: &notion.Text{
					Content: longContent[:CharacterLimit],
				},
				PlainText: longContent[:CharacterLimit],
			},
			{
				Type: notion.RichTextTypeText,
				Text: &notion.Text{
					Content: longContent[CharacterLimit:],
				},
				PlainText: longContent[CharacterLimit:],
			},
		}

		result := RichText(longContent)

		assert.Equal(t, len(expectedBlocks), len(result), "Expected %d blocks, but got %d", len(expectedBlocks), len(result))
		for i, block := range result {
			expectedBlock := expectedBlocks[i]

			assert.Equal(t, expectedBlock.Type, block.Type, "Expected block type %s, but got %s", expectedBlock.Type, block.Type)
			assert.Equal(t, expectedBlock.PlainText, block.PlainText, "Expected plain text %s, but got %s", expectedBlock.PlainText, block.PlainText)
			assert.Equal(t, expectedBlock.Text.Content, block.Text.Content, "Expected text content %s, but got %s", expectedBlock.Text.Content, block.Text.Content)
		}
	})
}
