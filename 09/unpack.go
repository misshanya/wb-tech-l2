package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrIncorrectString = errors.New("incorrect string")

func unpack(s string) (string, error) {
	var unpacked strings.Builder

	runes := []rune(s)

	var screened bool

	for i, char := range runes {
		if screened {
			unpacked.WriteRune(char)
			screened = false
			continue
		}

		if isDigit(char) {
			if i == 0 {
				return "", ErrIncorrectString
			}

			num, err := strconv.Atoi(string(char))
			if err != nil {
				return "", fmt.Errorf("failed to convert rune to int: %w", err)
			}

			unpacked.WriteString(strings.Repeat(string(runes[i-1]), num-1))
			continue
		}

		screened = char == '\\'
		if !screened {
			unpacked.WriteRune(char)
		}
	}

	if screened {
		return "", ErrIncorrectString
	}

	return unpacked.String(), nil
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}
