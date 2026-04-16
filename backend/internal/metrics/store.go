package metrics

import "sync"

type Store struct {
	mu         sync.RWMutex
	latest     *Snapshot
	history    []*Snapshot
	maxHistory int
}

func NewStore(maxHistory int) *Store {
	if maxHistory <= 0 {
		maxHistory = 1
	}

	return &Store{
		history:    make([]*Snapshot, 0, maxHistory),
		maxHistory: maxHistory,
	}
}

func (s *Store) Put(snapshot *Snapshot) {
	if snapshot == nil {
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.latest = snapshot
	s.history = append(s.history, snapshot)

	if len(s.history) > s.maxHistory {
		s.history = s.history[len(s.history)-s.maxHistory:]
	}
}

func (s *Store) GetHistory() []*Snapshot {
	s.mu.RLock()
	defer s.mu.RUnlock()

	history := make([]*Snapshot, len(s.history))
	copy(history, s.history)

	return history
}

func (s *Store) GetLatest() *Snapshot {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.latest
}
