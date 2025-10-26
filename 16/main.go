package main

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/misshanya/wb-tech-l2/15/downloader"
	"github.com/misshanya/wb-tech-l2/15/parser"
)

func parseAndDownloadRecursive(baseLink string, depth int) error {
	base, err := url.Parse(baseLink)
	if err != nil {
		return fmt.Errorf("failed to parse input url: %w", err)
	}

	baseDir := fmt.Sprintf("parsed/%s", base.Hostname())
	mainFilePath := fmt.Sprintf("%s/index.html", baseDir)

	visited := make(map[string]struct{})

	var f func(link string, path string, currentDepth, maxDepth int) error
	f = func(link string, path string, currentDepth, maxDepth int) error {
		if currentDepth >= maxDepth {
			return nil
		}

		if _, ok := visited[link]; ok {
			return nil
		}
		visited[link] = struct{}{}

		base, err := url.Parse(link)
		if err != nil {
			return fmt.Errorf("failed to parse input url: %w", err)
		}
		body, err := downloader.Download(link)
		if err != nil {
			return fmt.Errorf("failed to download file: %w", err)
		}

		err = downloader.Save(body, path)
		if err != nil {
			return fmt.Errorf("failed to save file: %w", err)
		}

		links, err := parser.FindLinks(body)
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

			nextPath := filepath.Join(baseDir, final.Path)
			if strings.HasSuffix(nextPath, string(filepath.Separator)) || filepath.Ext(nextPath) == "" {
				nextPath = filepath.Join(nextPath, "index.html")
			}
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
