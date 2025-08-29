package converter

import (
	"testing"

	"github.com/dstotijn/go-notion"
	"github.com/gomarkdown/markdown/ast"
	"github.com/stretchr/testify/assert"
)

func TestIsValidURL(t *testing.T) {
	t.Run("can validate URL", func(t *testing.T) {
		validURL := "https://example.com"
		invalidURL := "example.com"

		assert.True(t, isValidURL(validURL))
		assert.False(t, isValidURL(invalidURL))
	})

	t.Run("throws err when given table of contents", func(t *testing.T) {
		toc := "#internal-link"
		assert.False(t, isValidURL(toc))
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

	t.Run("falls back to plain text for relative paths", func(t *testing.T) {
		node := &ast.Link{
			Destination: []byte("./building.md"),
			Container: ast.Container{
				Children: []ast.Node{
					&ast.Leaf{
						Literal: []byte("Building"),
					},
				},
			},
		}

		// Expect plain text without link for invalid URLs
		expected := []notion.RichText{{
			Type:      notion.RichTextTypeText,
			PlainText: "Building",
			Text: &notion.Text{
				Content: "Building",
			},
		}}

		result := convertLinkToTextBlock(node)
		assert.Equal(t, expected, result)
	})
}
