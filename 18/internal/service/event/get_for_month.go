package event

import (
	"time"

	"github.com/misshanya/wb-tech-l2/18/internal/models"
)

// GetForMonth retrieves events for month
func (s *service) GetForMonth(date time.Time) []*models.Event {
	return s.repo.GetForMonth(date)
}
