package event

import (
	"time"

	"github.com/misshanya/wb-tech-l2/18/internal/models"
)

type event struct {
	ID     int
	UserID int
	Title  string
	Date   time.Time
}

func eventToModel(e *event) *models.Event {
	return &models.Event{
		ID:     e.ID,
		UserID: e.UserID,
		Title:  e.Title,
		Date:   e.Date,
	}
}

func modelToEvent(m *models.Event) *event {
	return &event{
		ID:     m.ID,
		UserID: m.UserID,
		Title:  m.Title,
		Date:   m.Date,
	}
}
