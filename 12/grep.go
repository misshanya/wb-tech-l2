package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type params struct {
	ContextAfterSize  int
	ContextBeforeSize int
	OnlyAmountOfLines bool
	IgnoreCase        bool
	InvertFilter      bool
	FixedLine         bool
	NumbersOfLines    bool
}

func grep(input []string, pattern string, p *params) ([]string, error) {
	var result []string
	linesToPrint := make(map[int]struct{})

	for i, line := range input {
		matchResult, err := isMatch(line, pattern, p)
		if err != nil {
			return []string{}, err
		}

		if p.InvertFilter {
			matchResult = !matchResult
		}

		if matchResult {
			linesToPrint[i] = struct{}{}

			for j := 1; j <= p.ContextBeforeSize && i-j >= 0; j++ {
				linesToPrint[i-j] = struct{}{}
			}

			for j := 1; j <= p.ContextAfterSize && i+j < len(input); j++ {
				linesToPrint[i+j] = struct{}{}
			}
		}
	}

	fmt.Println(linesToPrint)

	if p.OnlyAmountOfLines {
		return []string{strconv.Itoa(len(linesToPrint))}, nil
	}

	sortedIndexes := make([]int, 0, len(linesToPrint))

	for i := range linesToPrint {
		sortedIndexes = append(sortedIndexes, i)
	}
	sort.Ints(sortedIndexes)

	for _, i := range sortedIndexes {
		line := input[i]
		if p.NumbersOfLines {
			line = fmt.Sprintf("%v: %s", i+1, line)
		}
		result = append(result, line)
	}

	return result, nil
}

func isMatch(line, pattern string, p *params) (bool, error) {
	if p.IgnoreCase {
		line = strings.ToLower(line)
		pattern = strings.ToLower(pattern)
	}

	if p.FixedLine {
		return strings.Contains(line, pattern), nil
	}

	return regexp.MatchString(pattern, line)
}
