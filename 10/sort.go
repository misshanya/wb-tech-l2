package main

import (
	"strconv"
	"strings"
)

// sorter is a custom type for complex sort operations
type sorter struct {
	lines   []string
	column  int
	number  bool
	reverse bool
}

func (s *sorter) Len() int {
	return len(s.lines)
}

func (s *sorter) Swap(i, j int) {
	s.lines[i], s.lines[j] = s.lines[j], s.lines[i]
}

// Less performs checks and returns if one is less than other based on custom rules
func (s *sorter) Less(i, j int) bool {
	first := s.lines[i]
	second := s.lines[j]

	resultFirst := first
	resultSecond := second

	if s.column != 0 {
		firstSplitted := strings.Split(first, "\t")
		secondSplitted := strings.Split(second, "\t")

		if s.column <= len(firstSplitted) && s.column <= len(secondSplitted) {
			resultFirst = firstSplitted[s.column-1]
			resultSecond = secondSplitted[s.column-1]
		}
	}

	if s.reverse {
		resultFirst, resultSecond = resultSecond, resultFirst
	}

	if s.number {
		if numFirst, err := strconv.Atoi(resultFirst); err == nil {
			if numSecond, err := strconv.Atoi(resultSecond); err == nil {
				return numFirst < numSecond
			}
		}
	}

	return resultFirst < resultSecond
}
