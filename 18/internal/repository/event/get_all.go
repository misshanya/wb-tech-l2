package event

import "github.com/misshanya/wb-tech-l2/18/internal/models"

// GetAll returns all events
func (r *repo) GetAll() []*models.Event {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]*models.Event, len(r.storage))
	for id, event := range r.storage {
		e := eventToModel(event)
		e.ID = id
		result = append(result, e)
	}

	return result
}
