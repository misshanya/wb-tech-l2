package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

func processCommand(input string, in io.Reader, w io.Writer) error {
	parts := strings.Fields(input)

	if len(parts) < 1 {
		return nil
	}

	switch parts[0] {

	case "cd":
		return cmdCd(parts)

	case "pwd":
		out, err := cmdPwd()
		if err != nil {
			return err
		}
		fmt.Fprintln(w, out)

	case "echo":
		out := cmdEcho(parts)
		fmt.Fprintln(w, out)

	case "ps":
		out, err := cmdPs()
		if err != nil {
			return err
		}
		for _, line := range out {
			fmt.Fprintln(w, line)
		}

	case "kill":
		return cmdKill(parts)

	case "exit":
		os.Exit(0)

	default:
		out, err := cmdExternal(parts, in)
		if err != nil {
			return err
		}
		for _, line := range out {
			fmt.Fprintln(w, line)
		}
	}

	return nil
}

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

		pipes := strings.Split(input, "|")
		if len(pipes) <= 1 {
			err := processCommand(input, os.Stdin, os.Stdout)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			continue
		}

		var wg sync.WaitGroup
		var inputReader io.Reader = os.Stdin
		for i, cmd := range pipes {
			if i == len(pipes)-1 {
				err := processCommand(cmd, inputReader, os.Stdout)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
				}
				continue
			}

			pipeReader, pipeWriter := io.Pipe()

			wg.Add(1)
			go func(cmd string, stdin io.Reader, stdout io.WriteCloser) {
				defer wg.Done()
				defer stdout.Close()

				err := processCommand(cmd, stdin, stdout)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
				}
			}(cmd, inputReader, pipeWriter)

			inputReader = pipeReader
		}

		wg.Wait()
	}
}
