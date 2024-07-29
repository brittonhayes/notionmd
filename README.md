# NotionMD

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

Given the following markdown

```markdown
# H1 Example
- Item 1
- Item 2
```

You can convert it to Notion blocks like so:

```go
package main

import (
    "github.com/brittonhayes/notionmd/converter"
)

func main() {
    blocks, err := converter.Convert(markdownText)
    if err != nil {
        panic(err)
    }

    // Use the blocks with your Notion client
}
```

## Testing

To run the tests, use the following command:

```sh
go test ./...
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgements

- [go-notion](https://github.com/dstotijn/go-notion)
- [gomarkdown](https://github.com/gomarkdown/gomarkdown)

For more information, visit the [documentation](https://github.com/brittonhayes/notionmd).

---

This README provides a concise overview of the NotionMD project, including installation, usage, testing, and contribution guidelines.