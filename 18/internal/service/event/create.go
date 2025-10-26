package event

import "github.com/misshanya/wb-tech-l2/18/internal/models"

// Create creates new event
func (s *service) Create(e *models.Event) *models.Event {
	return s.repo.Create(e)
}
