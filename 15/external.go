package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func cmdExternal(args []string) ([]string, error) {
	cmd := exec.Command(args[0], args[1:]...)

	var stdoutBuf bytes.Buffer
	cmd.Stdout = &stdoutBuf

	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return []string{}, fmt.Errorf("failed to execute external command: %w", err)
	}

	out := strings.Split(
		strings.TrimSpace(stdoutBuf.String()),
		"\n",
	)

	return out, nil
}
