package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func cmdKill(args []string) error {
	if len(args) < 2 {
		return errors.New("usage: kill {pid}")
	}

	pid, err := strconv.Atoi(args[1])
	if err != nil {
		return errors.New("invalid pid")
	}

	process, err := os.FindProcess(pid)
	if err != nil {
		return fmt.Errorf("failed to find process: %w", err)
	}

	if err := process.Kill(); err != nil {
		return fmt.Errorf("failed to kill process: %w", err)
	}

	return nil
}
