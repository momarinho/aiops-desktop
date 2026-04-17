package actions

import (
	"sync"
)

type Store struct {
	mu      sync.RWMutex
	actions []*Action
}

func NewStore() *Store {
	return &Store{
		actions: make([]*Action, 0),
	}
}

func (s *Store) Add(action *Action) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.actions = append(s.actions, action)
}

func (s *Store) GetAll() []*Action {
	s.mu.RLock()
	defer s.mu.RUnlock()

	actions := make([]*Action, len(s.actions))
	copy(actions, s.actions)

	return actions
}

func (s *Store) GetByID(id string) (*Action, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, action := range s.actions {
		if action.ID == id {
			return action, true
		}
	}

	return nil, false
}
