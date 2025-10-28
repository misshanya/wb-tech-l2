package event

import (
	"time"

	"github.com/misshanya/wb-tech-l2/18/internal/models"
)

// GetForWeek retrieves events for week
func (s *service) GetForWeek(date time.Time) []*models.Event {
	return s.repo.GetForWeek(date)
}
