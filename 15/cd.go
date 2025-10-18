package main

import (
	"fmt"
	"os"
)

func cmdCd(args []string) error {
	var err error
	if len(args) > 1 {
		err = os.Chdir(args[1])
	} else {
		var homedir string
		homedir, err = os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("failed to get user home dir: %w", err)
		}
		err = os.Chdir(homedir)
	}
	if err != nil {
		return fmt.Errorf("failed to execute cd: %w", err)
	}

	return nil
}
