package event

import "github.com/misshanya/wb-tech-l2/18/internal/models"

// Update modifies the event
func (s *service) Update(e *models.Event) error {
	return s.repo.Update(e)
}
