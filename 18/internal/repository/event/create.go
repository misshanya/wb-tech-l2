package event

import (
	"github.com/misshanya/wb-tech-l2/18/internal/models"
)

// Create creates a new event in a database
func (r *repo) Create(e *models.Event) *models.Event {
	internalEvent := modelToEvent(e)

	r.mu.Lock()
	defer r.mu.Unlock()

	internalEvent.ID = r.lastID
	r.storage[internalEvent.ID] = internalEvent
	r.lastID++
	return eventToModel(internalEvent)
}
