package chunk

import "github.com/dstotijn/go-notion"

const (
	// BlockLimit is the maximum number of blocks that can be sent in a single request.
	BlockLimit = 100

	// CharacterLimit is the maximum number of characters that can be sent in a single rich text block.
	CharacterLimit = 2000
)

// Blocks splits an array of blocks into chunks of 100 blocks each.
func Blocks(blocks []notion.Block) [][]notion.Block {
	var chunks [][]notion.Block
	chunkSize := BlockLimit

	for i := 0; i < len(blocks); i += chunkSize {
		end := i + chunkSize
		if end > len(blocks) {
			end = len(blocks)
		}
		chunks = append(chunks, blocks[i:end])
	}

	return chunks
}

// RichText builds a new rich text block every 2000 characters of the provided string content.
func RichText(content string, annotations *notion.Annotations) []notion.RichText {
	var blocks []notion.RichText

	if len(content) <= CharacterLimit {
		richText := notion.RichText{
			Type: notion.RichTextTypeText,
			Text: &notion.Text{
				Content: content,
			},
			PlainText:   content,
			Annotations: annotations,
		}

		blocks = append(blocks, richText)
	} else {
		for i := 0; i < len(content); i += CharacterLimit {
			end := i + CharacterLimit
			if end > len(content) {
				end = len(content)
			}

			chunk := content[i:end]
			richText := notion.RichText{
				Type: notion.RichTextTypeText,
				Text: &notion.Text{
					Content: chunk,
				},
				PlainText:   chunk,
				Annotations: annotations,
			}

			blocks = append(blocks, richText)
		}
	}

	return blocks
}

// RichTextWithLink builds a new rich text block every 2000 characters of the provided string content with a link.
func RichTextWithLink(content string, link string) []notion.RichText {
	var blocks []notion.RichText

	if len(content) <= CharacterLimit {
		richText := notion.RichText{
			Type: notion.RichTextTypeText,
			Text: &notion.Text{
				Content: content,
				Link: &notion.Link{
					URL: link,
				},
			},
			PlainText: content,
		}

		blocks = append(blocks, richText)
	} else {
		for i := 0; i < len(content); i += CharacterLimit {
			end := i + CharacterLimit
			if end > len(content) {
				end = len(content)
			}

			chunk := content[i:end]
			richText := notion.RichText{
				Type: notion.RichTextTypeText,
				Text: &notion.Text{
					Content: chunk,
					Link: &notion.Link{
						URL: link,
					},
				},
				PlainText: chunk,
			}

			blocks = append(blocks, richText)
		}
	}

	return blocks
}
