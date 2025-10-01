package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
)

var ErrFileNotExist = errors.New("file does not exist")

func sortStrings(ss []string) []string {
	sort.Strings(ss)
	return ss
}

func scan(r io.Reader) ([]string, error) {
	var result []string

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}

	if err := scanner.Err(); err != nil {
		return []string{}, fmt.Errorf("failed to scan lines: %w", err)
	}

	return result, nil
}

func readFromStdin() ([]string, error) {
	return scan(os.Stdin)
}

func readFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []string{}, ErrFileNotExist
		}
		return []string{}, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	return scan(file)
}

func main() {
	var text []string
	var err error

	if len(os.Args) < 2 {
		text, err = readFromStdin()
		if err != nil {
			fmt.Printf("failed to read from stdin: %s\n", err)
			os.Exit(1)
		}
	} else {
		text, err = readFromFile(os.Args[1])
		if err != nil {
			if errors.Is(err, ErrFileNotExist) {
				fmt.Println("file does not exist")
				os.Exit(1)
			}
			fmt.Printf("failed to read from file: %s\n", err)
			os.Exit(1)
		}
	}

	sorted := sortStrings(text)
	for _, s := range sorted {
		fmt.Println(s)
	}
}
