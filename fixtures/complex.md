# Advanced Markdown Techniques

## Table of Contents
- [Advanced Markdown Techniques](#advanced-markdown-techniques)
  - [Table of Contents](#table-of-contents)
  - [Nested Lists](#nested-lists)
  - [Task Lists](#task-lists)
  - [Fenced Code Blocks](#fenced-code-blocks)
  - [Footnotes](#footnotes)
  - [Definition Lists](#definition-lists)
  - [Escaping Characters](#escaping-characters)

## Nested Lists

1. First level
    - Nested unordered list
    - Another item
        1. Nested ordered list
        2. Second item
2. Back to first level
    - [ ] Unchecked task
    - [x] Checked task

## Task Lists

- [x] Write the press release
- [ ] Update the website
- [ ] Contact the media

## Fenced Code Blocks

Here's a JavaScript code block with syntax highlighting:

```javascript
function calculateFactorial(n) {
    if (n === 0 || n === 1) {
        return 1;
    }
    return n * calculateFactorial(n - 1);
}

console.log(calculateFactorial(5)); // Output: 120
```

## Footnotes

Here's a sentence with a footnote[^1].

[^1]: This is the footnote content.

## Definition Lists

Term 1
: Definition 1

Term 2
: Definition 2a
: Definition 2b

## Escaping Characters

You can use \*asterisks\* without making text italic by escaping them with backslashes.

---

Here's a horizontal rule:

---

And finally, here's a ~~strikethrough~~ text.