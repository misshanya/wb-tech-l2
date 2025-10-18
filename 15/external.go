package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func cmdExternal(args []string, in io.Reader) ([]string, error) {
	cmd := exec.Command(args[0], args[1:]...)

	var stdoutBuf bytes.Buffer
	cmd.Stdout = &stdoutBuf

	cmd.Stdin = in
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
