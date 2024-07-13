package parser

import (
	"fmt"
	"strings"
)

type XMLNode struct {
	Tag        string
	Attributes map[string]string
	Parent     *XMLNode
	Children   []*XMLNode
	Text       string
}

func parseXML(xmlData string) (*XMLNode, error) {
	root := &XMLNode{}
	currentNode := root
	i := 0
	tagStack := []string{}

	if xmlData[0] != '<' {
		return nil, fmt.Errorf("invalid xml data")
	}
	for i < len(xmlData) {
		switch xmlData[i] {
		case '<':
			i++
			if xmlData[i] == '/' {
				i++
				endTag := ""
				for xmlData[i] != '>' {
					endTag += string(xmlData[i])
					i++
				}
				if len(tagStack) == 0 || tagStack[len(tagStack)-1] != endTag {
					return nil, fmt.Errorf("mismatched tags: %s and %s", endTag, tagStack[len(tagStack)-1])
				}
				currentNode = currentNode.Parent
				tagStack = tagStack[:len(tagStack)-1]
			} else {
				tagName := ""
				for xmlData[i] != '>' {
					tagName += string(xmlData[i])
					i++
				}
				if strings.Contains(tagName, "=") {
					return nil, fmt.Errorf("attribute is not supported")
				}
				newNode := &XMLNode{Tag: tagName, Parent: currentNode}
				currentNode.Children = append(currentNode.Children, newNode)
				currentNode = newNode
				tagStack = append(tagStack, tagName)
			}
		case '>':
			i++
		default:
			text := ""
			for i < len(xmlData) && xmlData[i] != '<' {
				text += string(xmlData[i])
				i++
			}
			currentNode.Text = text
		}
	}
	if len(tagStack) > 0 {
		return nil, fmt.Errorf("unclosed tags: %v", tagStack)
	}

	return root, nil
}

func AssertXml(xmlData string) bool {
	_, err := parseXML(xmlData)
	return err == nil
}
