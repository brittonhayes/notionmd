package converter

import (
	"testing"

	"github.com/dstotijn/go-notion"
	"github.com/gomarkdown/markdown/ast"
	"github.com/stretchr/testify/assert"
)

func TestIsHeading(t *testing.T) {
	t.Run("Heading node", func(t *testing.T) {
		node := &ast.Heading{}
		assert.True(t, isHeading(node))
	})

	t.Run("Not a heading node", func(t *testing.T) {
		node := &ast.Paragraph{}
		assert.False(t, isHeading(node))
	})
}

func TestConvertHeading(t *testing.T) {
	t.Run("can convert heading level 1", func(t *testing.T) {
		input := "Heading level 1"
		node := &ast.Heading{
			Level: 1,
			Container: ast.Container{
				Children: []ast.Node{
					&ast.Leaf{
						Content: []byte(input),
						Literal: []byte(input),
					},
				},
			},
		}

		expected := notion.Heading1Block{
			RichText: []notion.RichText{
				{
					Type:      notion.RichTextTypeText,
					Text:      &notion.Text{Content: input},
					PlainText: input,
				},
			},
		}

		result, err := convertHeading(node)
		assert.NoError(t, err)
		assertHeadingBlockEqual(t, expected, result)
	})

	t.Run("can convert heading level 2", func(t *testing.T) {
		input := "Heading level 2"
		node := &ast.Heading{
			Level: 2,
			Container: ast.Container{
				Children: []ast.Node{
					&ast.Leaf{
						Content: []byte("## " + input),
						Literal: []byte(input),
					},
				},
			},
		}

		expected := notion.Heading2Block{
			RichText: []notion.RichText{
				{
					Type:      notion.RichTextTypeText,
					Text:      &notion.Text{Content: input},
					PlainText: input,
				},
			},
		}

		result, err := convertHeading(node)
		assert.NoError(t, err)
		assertHeadingBlockEqual(t, expected, result)
	})

	t.Run("can convert heading level 3", func(t *testing.T) {
		input := "Heading level 3"
		node := &ast.Heading{
			Level: 3,
			Container: ast.Container{
				Children: []ast.Node{
					&ast.Leaf{
						Content: []byte(input),
						Literal: []byte(input),
					},
				},
			},
		}

		expected := notion.Heading3Block{
			RichText: []notion.RichText{
				{
					Type:      notion.RichTextTypeText,
					Text:      &notion.Text{Content: input},
					PlainText: input,
				},
			},
		}

		result, err := convertHeading(node)
		assert.NoError(t, err)
		assertHeadingBlockEqual(t, expected, result)
	})

	t.Run("can convert heading level 4", func(t *testing.T) {
		input := "Heading level 4"
		node := &ast.Heading{
			Level: 4,
			Container: ast.Container{
				Children: []ast.Node{
					&ast.Leaf{
						Content: []byte("#### " + input),
						Literal: []byte(input),
					},
				},
			},
		}

		result, err := convertHeading(node)
		assert.NoError(t, err)
		assert.Nil(t, result)
	})
}

func assertHeadingBlockEqual(t *testing.T, expected, actual notion.Block) {
	switch expected := expected.(type) {
	case notion.Heading1Block:
		actual, ok := actual.(notion.Heading1Block)
		assert.True(t, ok)
		assert.Equal(t, expected.RichText[0].PlainText, actual.RichText[0].PlainText)
		assert.Equal(t, expected.RichText[0].Text.Content, actual.RichText[0].Text.Content)
	case notion.Heading2Block:
		actual, ok := actual.(notion.Heading2Block)
		assert.True(t, ok)
		assert.Equal(t, expected.RichText[0].PlainText, actual.RichText[0].PlainText)
		assert.Equal(t, expected.RichText[0].Text.Content, actual.RichText[0].Text.Content)
	case notion.Heading3Block:
		actual, ok := actual.(notion.Heading3Block)
		assert.True(t, ok)
		assert.Equal(t, expected.RichText[0].PlainText, actual.RichText[0].PlainText)
		assert.Equal(t, expected.RichText[0].Text.Content, actual.RichText[0].Text.Content)
	default:
		t.Fatalf("unexpected block type: %T", expected)
	}
}
