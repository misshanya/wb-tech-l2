package main

import (
	"fmt"
	"os"
)

func cmdPwd() (string, error) {
	workDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get current workdir: %w", err)
	}
	return workDir, nil
}
