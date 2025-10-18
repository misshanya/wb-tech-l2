package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func cmdPs() ([]string, error) {
	var out []string

	out = append(out, fmt.Sprintf("%-8s %-20s", "PID", "PROGRAM"))

	entries, err := os.ReadDir("/proc")
	if err != nil {
		return []string{}, fmt.Errorf("failed to read /proc: %w", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		pid, err := strconv.Atoi(entry.Name())
		if err != nil {
			continue
		}

		comm, err := os.ReadFile(filepath.Join("/proc", entry.Name(), "comm"))
		if err != nil {
			return []string{}, fmt.Errorf("failed to read comm of pid %v: %w", pid, err)
		}
		program := strings.TrimSpace(string(comm))

		out = append(out, fmt.Sprintf("%-8d %-20s", pid, program))
	}

	return out, nil
}
