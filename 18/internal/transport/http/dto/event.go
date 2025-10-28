package dto

import "time"

type DateString string // date should be string to parse it from "YYYY-MM-DD"

type Event struct {
	ID     int       `json:"id"`
	UserID int       `json:"user_id"`
	Title  string    `json:"title"`
	Date   time.Time `json:"date"`
}

type EventCreateRequest struct {
	UserID int        `json:"user_id"`
	Title  string     `json:"title"`
	Date   DateString `json:"date"`
}

type EventCreateResponse Event

type EventUpdateRequest struct {
	ID     int        `json:"id"`
	UserID int        `json:"user_id"`
	Title  string     `json:"title"`
	Date   DateString `json:"date"`
}

type EventDeleteRequest struct {
	ID int `json:"id"`
}

type EventsGetForDayResponse []Event
