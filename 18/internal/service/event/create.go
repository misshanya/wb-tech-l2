package event

import "github.com/misshanya/wb-tech-l2/18/internal/models"

// Create creates new event
func (s *service) Create(e *models.Event) *models.Event {
	finalEvent := s.repo.Create(e)
	s.l.Info("created event", "event", e)
	return finalEvent
}
