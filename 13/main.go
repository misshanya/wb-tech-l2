package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/pflag"
)

var (
	ErrFileNotExist          = errors.New("file does not exist")
	ErrInvalidFieldsArgument = errors.New("invalid fields argument")
)

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

// parseFields performs a parse of colomns to print into a slice of indexes
func parseFields(input string) (map[int]struct{}, error) {
	result := make(map[int]struct{})
	parts := strings.SplitSeq(input, ",")

	for part := range parts {
		// Range of nums
		if strings.Contains(part, "-") {
			numsStr := strings.Split(part, "-")

			firstNum, err := strconv.Atoi(numsStr[0])
			if err != nil {
				return map[int]struct{}{}, err
			}

			secondNum, err := strconv.Atoi(numsStr[1])
			if err != nil {
				return map[int]struct{}{}, err
			}

			if firstNum > secondNum {
				return map[int]struct{}{}, ErrInvalidFieldsArgument
			}

			for i := firstNum; i <= secondNum; i++ {
				result[i] = struct{}{}
			}
		} else {
			num, err := strconv.Atoi(part)
			if err != nil {
				return map[int]struct{}{}, err
			}
			result[num] = struct{}{}
		}
	}

	return result, nil
}

func main() {
	fields := pflag.StringP("fields", "f", "", "specify the colomns to print")
	delimiter := pflag.StringP("delimiter", "d", "\t", "specify delimiter")
	separated := pflag.BoolP("separated", "s", false, "print only lines contain delimiter")

	pflag.Parse()

	fieldsParsed, err := parseFields(*fields)
	if err != nil {
		log.Fatalf("failed to parse fields: %s\n", err)
	}

	var input []string

	if pflag.NArg() < 1 {
		input, err = readFromStdin()
		if err != nil {
			fmt.Printf("failed to read from stdin: %s\n", err)
			os.Exit(1)
		}
	} else {
		input, err = readFromFile(pflag.Arg(0))
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
		Fields:    fieldsParsed,
		Delimiter: *delimiter,
		Separated: *separated,
	}
	result := separate(input, p)

	for _, s := range result {
		fmt.Println(s)
	}
}
