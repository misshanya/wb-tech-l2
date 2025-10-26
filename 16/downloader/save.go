package downloader

import (
	"fmt"
	"os"
	"path/filepath"
)

func Save(data []byte, path string) error {
	err := os.MkdirAll(filepath.Dir(path), 0o755)
	if err != nil {
		return fmt.Errorf("failed to create dirs: %w", err)
	}

	err = os.WriteFile(path, data, 0o644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
