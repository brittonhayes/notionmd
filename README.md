# NotionMD

[![Go Reference](https://pkg.go.dev/badge/github.com/brittonhayes/notionmd.svg)](https://pkg.go.dev/github.com/brittonhayes/notionmd)
[![Go Report Card](https://goreportcard.com/badge/github.com/brittonhayes/notionmd)](https://goreportcard.com/report/github.com/brittonhayes/notionmd)
![CI](https://github.com/brittonhayes/notionmd/actions/workflows/ci.yml/badge.svg)

A Go package designed to convert Markdown content into Notion Blocks seamlessly.

## Features

- Convert Markdown documents to Notion blocks
- Support for headings, links, lists, and paragraphs
- Easy integration with other Go projects

## Installation

To install NotionMD, use `go get`:

```sh
go get github.com/brittonhayes/notionmd
```

## Usage

You can convert Markdown content to Notion blocks using the `Convert` function. The function takes a string of Markdown content and returns a slice of Notion blocks.

```go
package main

import (
    "github.com/brittonhayes/notionmd/converter"
)

func main() {
    markdown := `
# H1 Example
- Item 1`

    blocks, err := converter.Convert(markdown)
    if err != nil {
        panic(err)
    }

    // Use the blocks with your Notion client
    result, _ := json.MarshalIndent(blocks, "", "  ")
    fmt.Println(string(m))
}
```

<details>
<summary>Output</summary>

```json
[
    {
        "heading_1": {
            "rich_text": [
                {
                    "type": "text",
                    "plain_text": "H1 Example",
                    "text": {
                    "content": "H1 Example"
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
                    "plain_text": "Item 1",
                    "text": {
                    "content": "Item 1"
                    }
                }
            ]
        }
    }
]
```

</details>

## Testing

To run the tests, use the following command:

```sh
go test ./...-v -cover
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgements

- [go-notion](https://github.com/dstotijn/go-notion)
- [gomarkdown](https://github.com/gomarkdown/gomarkdown)

For more information, visit the [documentation](https://pkg.go.dev/github.com/brittonhayes/notionmd).

---

This README provides a concise overview of the NotionMD project, including installation, usage, testing, and contribution guidelines.