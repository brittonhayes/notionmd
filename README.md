# NotionMD
Seamlessly Convert Markdown to Notion Blocks

[![Go Reference](https://pkg.go.dev/badge/github.com/brittonhayes/notionmd.svg)](https://pkg.go.dev/github.com/brittonhayes/notionmd)
[![Go Report Card](https://goreportcard.com/badge/github.com/brittonhayes/notionmd)](https://goreportcard.com/report/github.com/brittonhayes/notionmd)
![CI](https://github.com/brittonhayes/notionmd/actions/workflows/ci.yml/badge.svg)

NotionMD is a powerful Go package that bridges the gap between Markdown and Notion. It allows you to effortlessly convert your Markdown content into Notion blocks, making it easier than ever to integrate your existing Markdown documents into your Notion workspace.

## 🌟 Key Features

- **Markdown to Notion**: Convert Markdown documents to Notion blocks with a single function call
- **Rich Content Support**: Handles a variety of Markdown elements including headings, links, lists, and paragraphs
- **Large Document Handling**: Efficiently processes large documents by breaking blocks into manageable chunks
- **Easy Integration**: Designed to work seamlessly with Notion API clients

## 🚀 Quick Start

### Installation

Get started with NotionMD in your Go project:

```sh
go get github.com/brittonhayes/notionmd
```

### Basic Usage

Here's a simple example of how to use NotionMD:

```go
package main

import (
    "encoding/json"
    "log"
    "fmt"
    "github.com/brittonhayes/notionmd"
)

func main() {
    markdown := `
# Welcome to NotionMD
- Convert Markdown easily
- Integrate with Notion seamlessly`

    blocks, err := notionmd.Convert(markdown)
    if err != nil {
       log.Fatal(err) 
    }

    // Print the resulting Notion blocks
    result, _ := json.MarshalIndent(blocks, "", "  ")
    fmt.Println(string(result))
}
```

<details>
<summary>Click to see the output</summary>

```json
[
  {
    "heading_1": {
      "rich_text": [
        {
          "type": "text",
          "plain_text": "Welcome to NotionMD",
          "text": {
            "content": "Welcome to NotionMD"
          }
        }
      ],
      "is_toggleable": false
    }
  },
  {
    "bulleted_list_item": {
      "rich_text": [
        {
          "type": "text",
          "plain_text": "Convert Markdown easily",
          "text": {
            "content": "Convert Markdown easily"
          }
        }
      ]
    }
  },
  {
    "bulleted_list_item": {
      "rich_text": [
        {
          "type": "text",
          "plain_text": "Integrate with Notion seamlessly",
          "text": {
            "content": "Integrate with Notion seamlessly"
          }
        }
      ]
    }
  }
]
```
</details>

### Creating a Notion Page from Markdown

This example demonstrates how to create a Notion page using the blocks parsed from a Markdown document

```go
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
    parentPageID := "your-parent-page-id" // Replace with your actual parent page ID
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
```

> [!NOTE]  
>  Make sure to set the `NOTION_API_KEY` environment variable with your Notion API key and replace `"your-parent-page-id"` with the actual ID of the parent page where you want to create the new page.

## 🧪 Testing

Ensure the reliability of NotionMD by running the test suite:

```sh
go test ./... -v -cover
```

## 🤝 Contributing

We welcome contributions from the community! If you'd like to contribute:

1. Fork the repository
2. Create a new branch for your feature or bug fix
3. Make your changes and write tests if applicable
4. Submit a pull request with a clear description of your changes

Please open an issue if you find a bug or have a feature request.

## 📄 License

NotionMD is open-source software licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgements

- [go-notion](https://github.com/dstotijn/go-notion) for Notion API interactions
- [gomarkdown](https://github.com/gomarkdown/gomarkdown) for Markdown parsing

## 📚 Learn More

For detailed API documentation and advanced usage examples, visit our [Go Package Documentation](https://pkg.go.dev/github.com/brittonhayes/notionmd).

---

Built with 🖤 by [Britton Hayes](https://github.com/brittonhayes) and contributors. If you find NotionMD useful, consider giving it a star on GitHub!