package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")

		input, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Println()
				break
			}

			fmt.Println("failed to read line:", err)
			continue
		}

		input = strings.TrimSpace(input)

		parts := strings.Fields(input)

		if len(parts) < 1 {
			continue
		}

		switch parts[0] {

		case "cd":
			err := cmdCd(parts)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}

		case "pwd":
			out, err := cmdPwd()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			fmt.Println(out)

		case "echo":
			out := cmdEcho(parts)
			fmt.Println(out)

		case "ps":
			out, err := cmdPs()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}

			for _, line := range out {
				fmt.Println(line)
			}

		case "kill":
			err := cmdKill(parts)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}

		case "exit":
			os.Exit(0)

		default:
			out, err := cmdExternal(parts)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			for _, line := range out {
				fmt.Println(line)
			}
		}
	}
}
