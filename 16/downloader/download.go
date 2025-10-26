package downloader

import (
	"fmt"
	"io"
	"net/http"
)

func Download(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to request: %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to read response body: %w", err)
	}
	defer resp.Body.Close()

	return body, nil
}
