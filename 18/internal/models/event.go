package models

import "time"

type Event struct {
	ID     int
	UserID int
	Title  string
	Date   time.Time
}
