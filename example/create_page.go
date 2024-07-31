package main

import (
	"context"
	"log"
	"os"

	"github.com/brittonhayes/notionmd"
	"github.com/dstotijn/go-notion"
)

func main() {
	// Read the Markdown file
	markdown, err := os.ReadFile("example.md")
	if err != nil {
		log.Fatalf("Error reading Markdown file: %v", err)
	}

	// Convert the Markdown content to Notion blocks
	blocks, err := notionmd.Convert(string(markdown))
	if err != nil {
		log.Fatalf("Error converting Markdown to Notion blocks: %v", err)
	}

	// Initialize the Notion client
	client := notion.NewClient(os.Getenv("NOTION_API_KEY"))

	// Create a new page in Notion
	parentPageID := "page-id" // Replace with your actual parent page ID
	newPage, err := client.CreatePage(context.Background(), notion.CreatePageParams{
		ParentType: notion.ParentTypePage,
		ParentID:   parentPageID,
		Title: []notion.RichText{
			{
				Text: &notion.Text{
					Content: "Markdown to Notion Example",
				},
			},
		},
		Children: blocks,
	})
	if err != nil {
		log.Fatalf("Error creating Notion page: %v", err)
	}

	log.Printf("Successfully created Notion page with ID: %s", newPage.ID)
}
