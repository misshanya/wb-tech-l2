package main

import "strings"

func cmdEcho(args []string) string {
	if len(args) > 1 {
		return strings.Join(args[1:], " ")
	}
	return ""
}
