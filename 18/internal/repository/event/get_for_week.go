package event

import (
	"time"

	"github.com/misshanya/wb-tech-l2/18/internal/models"
)

// GetForWeek retrieves events for week
func (r *repo) GetForWeek(date time.Time) []*models.Event {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []*models.Event
	for id, event := range r.storage {
		y1, w1 := date.ISOWeek()
		y2, w2 := event.Date.ISOWeek()
		if y1 == y2 && w1 == w2 {
			e := eventToModel(event)
			e.ID = id
			result = append(result, e)
		}
	}

	return result
}
