package event

// Delete removes the event
func (s *service) Delete(id int) {
	s.repo.Delete(id)
	s.l.Info("deleted event", "id", id)
}
