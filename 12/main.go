package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/pflag"
)

var ErrFileNotExist = errors.New("file does not exist")

// scan performs read from reader by line
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

// readFromStdin performs read from standard input
func readFromStdin() ([]string, error) {
	return scan(os.Stdin)
}

// readFromFile performs a read of file content with specified filename
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

// makeUnique returns a slice of unique strings of an input slice
func makeUnique(text []string) []string {
	var result []string
	uniqueMap := make(map[string]struct{})
	for _, line := range text {
		if _, ok := uniqueMap[line]; !ok {
			uniqueMap[line] = struct{}{}
			result = append(result, line)
		}
	}
	return result
}

func main() {
	contextAfterSize := pflag.IntP("context-after", "A", 0, "print N lines after each line")
	contextBeforeSize := pflag.IntP("context-before", "B", 0, "print N lines before each line")
	contextAllSize := pflag.IntP("context", "C", 0, "print N lines around each line")
	onlyAmountOfLines := pflag.BoolP("count", "c", false, "print only amount of lines")
	ignoreCase := pflag.BoolP("ignore-case", "i", false, "ignore case")
	invertFilter := pflag.BoolP("invert", "v", false, "invert filter")
	fixedLine := pflag.BoolP("fixed", "F", false, "perceive the template as a fixed string")
	numbersOfLines := pflag.BoolP("numbers", "n", false, "print the number of each line")

	pflag.Parse()

	// Add context to equal -C N = -A N -B N
	*contextAfterSize += *contextAllSize
	*contextBeforeSize += *contextAllSize

	var input []string
	var err error

	if pflag.NArg() == 0 {
		log.Fatal("search pattern not provided")
	}
	pattern := pflag.Arg(0)

	if pflag.NArg() < 2 {
		input, err = readFromStdin()
		if err != nil {
			fmt.Printf("failed to read from stdin: %s\n", err)
			os.Exit(1)
		}
	} else {
		input, err = readFromFile(pflag.Arg(1))
		if err != nil {
			if errors.Is(err, ErrFileNotExist) {
				fmt.Println("file does not exist")
				os.Exit(1)
			}
			fmt.Printf("failed to read from file: %s\n", err)
			os.Exit(1)
		}
	}

	p := &params{
		ContextAfterSize:  *contextAfterSize,
		ContextBeforeSize: *contextBeforeSize,
		ContextAllSize:    *contextAllSize,
		OnlyAmountOfLines: *onlyAmountOfLines,
		IgnoreCase:        *ignoreCase,
		InvertFilter:      *invertFilter,
		FixedLine:         *fixedLine,
		NumbersOfLines:    *numbersOfLines,
	}
	result, err := grep(input, pattern, p)
	if err != nil {
		log.Fatal(err)
	}

	for _, s := range result {
		fmt.Println(s)
	}
}
