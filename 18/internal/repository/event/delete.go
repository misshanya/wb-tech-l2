package event

// Delete removes the event from the database
func (r *repo) Delete(id int) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.storage, id)
}
