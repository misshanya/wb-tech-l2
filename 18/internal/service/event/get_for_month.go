package event

import (
	"time"

	"github.com/misshanya/wb-tech-l2/18/internal/models"
)

// GetForMonth retrieves events for month
func (s *service) GetForMonth(date time.Time) []*models.Event {
	events := s.repo.GetAll()

	var result []*models.Event
	for _, event := range events {
		if date.Year() == event.Date.Year() &&
			date.Month() == event.Date.Month() {
			result = append(result, event)
		}
	}

	return result
}
