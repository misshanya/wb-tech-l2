package event

import (
	"time"

	"github.com/misshanya/wb-tech-l2/18/internal/models"
)

// Get retrieves events for day
func (r *repo) GetForDay(date time.Time) []*models.Event {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []*models.Event
	for id, event := range r.storage {
		if event.Date.Year() == date.Year() &&
			event.Date.Month() == date.Month() &&
			event.Date.Day() == date.Day() {
			e := eventToModel(event)
			e.ID = id
			result = append(result, e)
		}
	}

	return result
}
