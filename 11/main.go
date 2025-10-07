package main

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

func findAnagrams(ss []string) map[string][]string {
	unique := make(map[string]struct{})
	groups := make(map[string][]string)

	for _, word := range ss {
		word = strings.ToLower(word)
		if _, ok := unique[word]; ok {
			continue
		}
		unique[word] = struct{}{}
		runes := []rune(word)
		slices.Sort(runes)
		sortedKey := string(runes)
		groups[sortedKey] = append(groups[sortedKey], word)
	}

	result := make(map[string][]string)
	for _, group := range groups {
		if len(group) > 1 {
			key := group[0]
			sort.Strings(group)
			result[key] = group
		}
	}

	return result
}

func main() {
	words := []string{
		"пятак", "пятка", "тяпка", "Пятка",
		"листок", "слиток", "столик", "стол",
	}

	anagramsMap := findAnagrams(words)

	for k, v := range anagramsMap {
		fmt.Printf("%s: %s\n", k, strings.Join(v, ", "))
	}
}
