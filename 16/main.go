package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"golang.org/x/net/html"
)

func findLinks(body []byte) ([]string, error) {
	r := bytes.NewReader(body)
	node, err := html.Parse(r)
	if err != nil {
		return []string{}, fmt.Errorf("failed to parse html: %w", err)
	}

	var links []string

	var f func(*html.Node)
	f = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, a := range node.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
					break
				}
			}
		}
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(node)

	return links, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: miniget <url>")
		os.Exit(1)
	}

	url := os.Args[1]

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("failed to request: %s\n", err)
		os.Exit(1)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("failed to read response body: %s\n", err)
		os.Exit(1)
	}
	resp.Body.Close()

	path := "parsed/index.html"

	err = os.MkdirAll(filepath.Dir(path), 0o755)
	if err != nil {
		fmt.Printf("failed to create dirs: %s", err)
		os.Exit(1)
	}

	err = os.WriteFile(path, body, 0o644)
	if err != nil {
		fmt.Printf("failed to write file: %s\n", err)
		os.Exit(1)
	}

	links, err := findLinks(body)
	if err != nil {
		fmt.Printf("failed to find links: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("links:", links)
}
