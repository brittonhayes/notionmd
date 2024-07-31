package converter

import (
	"testing"

	"github.com/dstotijn/go-notion"
	"github.com/gomarkdown/markdown/ast"
	"github.com/stretchr/testify/assert"
)

func TestConvertLink(t *testing.T) {
	t.Run("can convert link", func(t *testing.T) {
		node := &ast.Link{
			Destination: []byte("https://example.com"),
			Container: ast.Container{
				Children: []ast.Node{
					&ast.Leaf{
						Literal: []byte("Example"),
					},
				},
			},
		}
		expectedBlock := &notion.ParagraphBlock{
			RichText: []notion.RichText{
				{
					Type: notion.RichTextTypeText,
					Text: &notion.Text{
						Content: "Example",
						Link: &notion.Link{
							URL: "https://example.com",
						},
					},
					PlainText: "Example",
				},
			},
		}

		block, err := convertLink(node)
		assert.NoError(t, err)
		assert.Equal(t, expectedBlock, block)
	})

	t.Run("can convert link from markdown", func(t *testing.T) {
		markdownText := `[Example](https://example.com)`
		expected := []notion.Block{
			&notion.ParagraphBlock{
				RichText: []notion.RichText{
					{
						Type: notion.RichTextTypeText,
						Text: &notion.Text{
							Content: "Example",
							Link: &notion.Link{
								URL: "https://example.com",
							},
						},
						PlainText: "Example",
					},
				},
			},
		}

		result, err := Convert(markdownText)
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})
}

func TestIsLink(t *testing.T) {
	linkNode := &ast.Link{}
	paragraphNode := &ast.Paragraph{}

	assert.True(t, isLink(linkNode))
	assert.False(t, isLink(paragraphNode))
}

func TestExtractURL(t *testing.T) {
	t.Run("can extract URL", func(t *testing.T) {
		node := &ast.Link{
			Destination: []byte("https://example.com"),
			Container: ast.Container{
				Children: []ast.Node{
					&ast.Leaf{
						Literal: []byte("Example"),
					},
				},
			},
		}

		expectedURL := "https://example.com"

		url := extractURL(node)
		assert.Equal(t, expectedURL, url)
	})
}

func TestExtractTitle(t *testing.T) {
	t.Run("can extract title", func(t *testing.T) {
		node := &ast.Link{
			Destination: []byte("https://example.com"),
			Container: ast.Container{
				Children: []ast.Node{
					&ast.Leaf{
						Literal: []byte("Example"),
					},
				},
			},
		}

		expectedTitle := "Example"

		title := extractTitle(node)
		assert.Equal(t, expectedTitle, title)
	})
}

func TestConvertLinkToTextBlock(t *testing.T) {
	t.Run("can convert link to text block", func(t *testing.T) {
		node := &ast.Link{
			Destination: []byte("https://example.com"),
			Container: ast.Container{
				Children: []ast.Node{
					&ast.Leaf{
						Literal: []byte("Example"),
					},
				},
			},
		}

		expected := []notion.RichText{{
			Type:      notion.RichTextTypeText,
			PlainText: "Example",
			Text: &notion.Text{
				Content: "Example",
				Link: &notion.Link{
					URL: "https://example.com",
				},
			},
		}}

		result := convertLinkToTextBlock(node)
		assert.Equal(t, expected, result)
	})
}
