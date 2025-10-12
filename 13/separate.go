package main

import "strings"

type params struct {
	Fields    map[int]struct{}
	Delimiter string
	Separated bool
}

func separate(input []string, p *params) []string {
	var result []string

	for _, line := range input {
		var preResult []string

		parts := strings.Split(line, p.Delimiter)
		if len(parts) == 1 && p.Separated {
			continue
		}

		for i, part := range parts {
			if _, ok := p.Fields[i+1]; ok {
				preResult = append(preResult, part)
			}
		}

		if len(preResult) > 0 {
			result = append(result, strings.Join(preResult, p.Delimiter))
		}
	}

	return result
}
