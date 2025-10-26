package parser

import (
	"bytes"
	"fmt"

	"golang.org/x/net/html"
)

func FindLinks(body []byte) ([]string, error) {
	r := bytes.NewReader(body)
	node, err := html.Parse(r)
	if err != nil {
		return []string{}, fmt.Errorf("failed to parse html: %w", err)
	}

	var links []string

	findAndAppendAttr := func(attrs []html.Attribute, key string) {
		for _, a := range attrs {
			if a.Key == key {
				links = append(links, a.Val)
				break
			}
		}
	}

	var f func(*html.Node)
	f = func(node *html.Node) {
		if node.Type == html.ElementNode {
			switch node.Data {
			case "a", "link":
				findAndAppendAttr(node.Attr, "href")
			case "script", "img":
				findAndAppendAttr(node.Attr, "src")
			}
		}
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(node)

	return links, nil
}
