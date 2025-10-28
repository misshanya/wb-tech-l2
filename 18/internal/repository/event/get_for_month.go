package event

import (
	"time"

	"github.com/misshanya/wb-tech-l2/18/internal/models"
)

// GetForMonth retrieves events for month
func (r *repo) GetForMonth(date time.Time) []*models.Event {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []*models.Event
	for id, event := range r.storage {
		if date.Year() == event.Date.Year() &&
			date.Month() == event.Date.Month() {
			e := eventToModel(event)
			e.ID = id
			result = append(result, e)
		}
	}

	return result
}
