package event

import (
	"time"

	"github.com/misshanya/wb-tech-l2/18/internal/models"
)

// GetForDay retrieves events for day
func (s *service) GetForDay(date time.Time) []*models.Event {
	events := s.repo.GetAll()

	var result []*models.Event
	for _, event := range events {
		if event.Date.Year() == date.Year() &&
			event.Date.Month() == date.Month() &&
			event.Date.Day() == date.Day() {
			result = append(result, event)
		}
	}

	return result
}
