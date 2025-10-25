package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

func download(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to request: %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to read response body: %w", err)
	}
	defer resp.Body.Close()

	return body, nil
}

func save(data []byte, path string) error {
	err := os.MkdirAll(filepath.Dir(path), 0o755)
	if err != nil {
		return fmt.Errorf("failed to create dirs: %w", err)
	}

	err = os.WriteFile(path, data, 0o644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func parseAndDownloadRecursive(baseLink string, depth int) error {
	base, err := url.Parse(baseLink)
	if err != nil {
		return fmt.Errorf("failed to parse input url: %w", err)
	}

	baseDir := fmt.Sprintf("parsed/%s", base.Hostname())
	mainFilePath := fmt.Sprintf("%s/index.html", baseDir)

	var f func(link string, path string, currentDepth, maxDepth int) error
	f = func(link string, path string, currentDepth, maxDepth int) error {
		if currentDepth >= maxDepth {
			return nil
		}

		base, err := url.Parse(link)
		if err != nil {
			return fmt.Errorf("failed to parse input url: %w", err)
		}
		body, err := download(link)
		if err != nil {
			return fmt.Errorf("failed to download file: %w", err)
		}

		_, err = os.Stat(path)
		if err != nil {
			if !os.IsNotExist(err) {
				return err
			}
		}

		err = save(body, path)
		if err != nil {
			return fmt.Errorf("failed to save file: %w", err)
		}

		links, err := findLinks(body)
		if err != nil {
			return fmt.Errorf("failed to find links: %w", err)
		}

		for _, link := range links {
			ref, err := url.Parse(link)
			if err != nil {
				fmt.Printf("failed to parse ref url: %s\n", err)
				continue
			}
			final := base.ResolveReference(ref)
			if final.Hostname() != base.Hostname() || final.Scheme != "http" && final.Scheme != "https" {
				continue
			}

			if strings.TrimPrefix(final.Path, "/") == "" {
				continue
			}
			nextPath := fmt.Sprintf("%s/%s", baseDir, strings.TrimPrefix(final.Path, "/"))
			if err := f(
				final.String(),
				nextPath,
				currentDepth+1, maxDepth,
			); err != nil {
				fmt.Printf("failed to parse and download: %s\n", err)
				continue
			}
		}

		return nil
	}

	return f(baseLink, mainFilePath, 0, depth)
}

func findLinks(body []byte) ([]string, error) {
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

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: miniget <url>")
		os.Exit(1)
	}

	inputURL := os.Args[1]

	if err := parseAndDownloadRecursive(inputURL, 10); err != nil {
		fmt.Printf("failed to parse and download: %s\n", err)
		os.Exit(1)
	}
}
