package event

import (
	"github.com/misshanya/wb-tech-l2/18/internal/errorz"
	"github.com/misshanya/wb-tech-l2/18/internal/models"
)

// Update modifies event in the database
func (r *repo) Update(e *models.Event) error {
	internalEvent := modelToEvent(e)

	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.storage[internalEvent.ID]; !ok {
		return errorz.EventNotFound
	}

	r.storage[internalEvent.ID] = internalEvent
	return nil
}
