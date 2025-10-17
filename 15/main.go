package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
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
			var err error
			if len(parts) > 1 {
				err = os.Chdir(parts[1])
			} else {
				var homedir string
				homedir, err = os.UserHomeDir()
				if err != nil {
					fmt.Println("failed to get user home dir")
					continue
				}
				err = os.Chdir(homedir)
			}
			if err != nil {
				fmt.Println("failed to execute cd:", err)
				continue
			}

		case "pwd":
			workDir, err := os.Getwd()
			if err != nil {
				fmt.Println("failed to get current workdir:", err)
				continue
			}
			fmt.Println(workDir)

		case "echo":
			if len(parts) > 1 {
				fmt.Println(strings.Join(parts[1:], " "))
			}

		case "exit":
			os.Exit(0)

		default:
			cmd := exec.Command(parts[0], parts[1:]...)

			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout

			err := cmd.Run()
			if err != nil {
				fmt.Println("failed to execute external command:", err)
				continue
			}
		}
	}
}
