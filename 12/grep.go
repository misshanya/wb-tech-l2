package main

import (
	"regexp"
	"strconv"
	"strings"
)

type params struct {
	ContextAfterSize  int
	ContextBeforeSize int
	ContextAllSize    int
	OnlyAmountOfLines bool
	IgnoreCase        bool
	InvertFilter      bool
	FixedLine         bool
	NumbersOfLines    bool
}

func grep(input []string, pattern string, p *params) ([]string, error) {
	var result []string

	for _, s := range input {
		matchResult, err := isMatch(s, pattern, p)
		if err != nil {
			return []string{}, err
		}
		if p.InvertFilter {
			matchResult = !matchResult
		}

		if matchResult {
			result = append(result, s)
		}
	}

	if p.OnlyAmountOfLines {
		return []string{strconv.Itoa(len(result))}, nil
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
