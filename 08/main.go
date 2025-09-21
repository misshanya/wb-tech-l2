package main

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/beevik/ntp"
)

const ntpServer = "0.ru.pool.ntp.org"

func getCurrentTime() (time.Time, error) {
	return ntp.Time(ntpServer)
}

func main() {
	currentTime, err := getCurrentTime()
	if err != nil {
		slog.Error("failed to get current time",
			"error", err,
		)
		os.Exit(1)
	}

	fmt.Printf("Current time: %s\n", currentTime)
}
