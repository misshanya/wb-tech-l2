package event

import (
	"time"

	"github.com/misshanya/wb-tech-l2/18/internal/models"
)

// GetForWeek retrieves events for week
func (s *service) GetForWeek(date time.Time) []*models.Event {
	events := s.repo.GetAll()

	var result []*models.Event
	for _, event := range events {
		y1, w1 := date.ISOWeek()
		y2, w2 := event.Date.ISOWeek()
		if y1 == y2 && w1 == w2 {
			result = append(result, event)
		}
	}

	return result
}
