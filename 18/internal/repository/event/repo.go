package event

import "sync"

type repo struct {
	storage map[int]*event
	lastID  int
	mu      *sync.RWMutex
}

func New() *repo {
	return &repo{
		storage: make(map[int]*event),
		mu:      &sync.RWMutex{},
	}
}
