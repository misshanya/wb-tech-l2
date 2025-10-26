package event

import (
	"time"

	"github.com/misshanya/wb-tech-l2/18/internal/models"
)

// Get retrieves events for day
func (s *service) GetForDay(date time.Time) []*models.Event {
	return s.repo.GetForDay(date)
}
