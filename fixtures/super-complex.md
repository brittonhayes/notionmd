# Advanced Markdown Techniques and Edge Cases

## Table of Contents
- [Advanced Markdown Techniques and Edge Cases](#advanced-markdown-techniques-and-edge-cases)
  - [Table of Contents](#table-of-contents)
  - [Multi-level Headings](#multi-level-headings)
- [H1](#h1)
  - [H2](#h2)
    - [H3](#h3)
      - [H4](#h4)
        - [H5](#h5)
          - [H6](#h6)
  - [Text Formatting Combinations](#text-formatting-combinations)
  - [Complex Lists](#complex-lists)
  - [Code Blocks with Annotations](#code-blocks-with-annotations)
  - [Blockquotes with Nested Elements](#blockquotes-with-nested-elements)
  - [Link Variations](#link-variations)
  - [Escaping and Special Characters](#escaping-and-special-characters)
  - [HTML and Markdown Mix](#html-and-markdown-mix)
  - [Definition Lists with Nested Elements](#definition-lists-with-nested-elements)
  - [Footnotes and References](#footnotes-and-references)

## Multi-level Headings

# H1
## H2
### H3
#### H4
##### H5
###### H6

####### This is not a valid heading (H7)

## Text Formatting Combinations

This text is ***bold and italic***. 
This one is ~~strikethrough and **bold**~~.
Here's some *italic text with **bold** inside*.
Let's try `inline code with *italic*` and **bold**.

## Complex Lists

1. First level item
    - Nested unordered list
    - Another item
        1. Nested ordered list
        2. Second item
            - Even deeper nesting
                1. Is this too deep?
                2. Probably, but it's a good test!
    - Back to second level
2. First level again
    - [ ] Unchecked task
    - [x] Checked task
        - [ ] Nested unchecked task
        - [x] Nested checked task

## Code Blocks with Annotations

```python
def factorial(n):
    """
    Calculate the factorial of a number.
    """
    if n == 0 or n == 1:
        return 1
    else:
        return n * factorial(n - 1)  # Recursive call

# Example usage:
result = factorial(5)
print(f"5! = {result}")  # Output: 5! = 120
```

Inline code: `print("Hello, World!")`

## Blockquotes with Nested Elements

> This is a blockquote.
> 
> It can contain multiple paragraphs.
>
>> This is a nested blockquote.
>>
>> - It can contain lists
>> - And other elements
>
> Back to the first level.
>
> 1. Ordered list in blockquote
> 2. Second item
>     ```
>     Code block in blockquote
>     ```

## Link Variations

[Basic link](https://www.example.com)
[Link with title](https://www.example.com "Example Website")
[Reference-style link][ref1]
[Numbered reference-style link][1]
[Link to heading](#complex-lists)

[ref1]: https://www.example.com
[1]: https://www.example.com/another-page

## Escaping and Special Characters

\*This text is surrounded by asterisks but not italic\*

\# This is not a heading

Use `\` to escape \*, \`, \#, and other special characters.

## HTML and Markdown Mix

<details>
<summary>Click to expand</summary>

This content is inside an HTML `details` tag.

- It can contain Markdown
- Like this list
- And *formatted* **text**

</details>

## Definition Lists with Nested Elements

Term 1
: Definition 1
: Another definition for Term 1
    > With a blockquote
    ```
    And a code block
    ```

Term 2
: Definition 2
    1. With an ordered list
    2. Second item

## Footnotes and References

Here's a sentence with a footnote[^1]. And another[^2]. And a named one[^note].

[^1]: This is the first footnote.
[^2]: This is the second footnote.
[^note]: This is a named footnote.

---

*This document ends with a horizontal rule and some emphasized text.*