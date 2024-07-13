# Pull request description-like document

## Construction

### Description

Defined XMLNode to be handled appropriately. `#parseXML` is broken down into three processes.

1. start xml node
  This detects the starting position of the xml tag. And for xmlNode, it keeps the tag name and tag, node nesting structure.
2. close xml node
  This detects the end position of the xml tag. Proceed the index.
3. text
  Get non-tag name text and record it for the currentNode.Text


### Modules

- main.go ... execution file
- parser ... extracted parser files, this includes xml parser that is created from scratch and its test file
  - xml.go
  - xml_test.go


## How to execute code

```bash
$ go run main.go "<People age="1">hello world</People>"
```

### Tests

```bash
$ go test -v ./parser
```

## Supplement

If an empty string is passed, an error is output as panic.

Here are some examples

```bash
➜ go run main.go ""
panic: runtime error: index out of range [0] with length 0

goroutine 1 [running]:
main/parser.parseXML({0x16f1c2f29, 0x0})
main/parser.AssertXml(...)
main.main()
exit status 2
```

