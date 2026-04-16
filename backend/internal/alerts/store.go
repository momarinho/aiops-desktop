package alerts

import (
	"errors"
	"sort"
	"sync"
	"time"
)

var ErrAlertNotFound = errors.New("alert not found")

type Store struct {
	mu     sync.RWMutex
	alerts map[string]*Alert
}

func NewStore() *Store {
	return &Store{
		alerts: make(map[string]*Alert),
	}
}

func (s *Store) List() []Alert {
	s.mu.RLock()
	defer s.mu.RUnlock()

	out := make([]Alert, 0, len(s.alerts))
	for _, a := range s.alerts {
		out = append(out, *cloneAlert(a))
	}

	sort.Slice(out, func(i, j int) bool {
		return out[i].UpdatedAt.After(out[j].UpdatedAt)
	})

	return out
}

func (s *Store) Activate(rule Rule, value float64, at time.Time) {
	s.mu.Lock()
	defer s.mu.Unlock()

	alert, ok := s.alerts[rule.ID]
	if !ok {
		alert = &Alert{
			ID:          rule.ID,
			Severity:    rule.Severity,
			Description: rule.Description,
			MetricName:  rule.MetricName,
			Threshold:   rule.Threshold,
		}
		s.alerts[rule.ID] = alert
	}

	if alert.StartedAt == nil || alert.Status == StatusResolved {
		alert.StartedAt = timePtr(at)
		alert.AcknowledgedAt = nil
		alert.SilencedAt = nil
		alert.ResolvedAt = nil
		alert.Status = StatusFiring
	} else if alert.Status != StatusAcknowledged && alert.Status != StatusSilenced {
		alert.Status = StatusFiring
	}

	alert.CurrentValue = value
	alert.UpdatedAt = at
}

func (s *Store) Resolve(id string, value float64, at time.Time) {
	s.mu.Lock()
	defer s.mu.Unlock()

	alert, ok := s.alerts[id]
	if !ok {
		return
	}

	alert.CurrentValue = value
	alert.UpdatedAt = at
	alert.ResolvedAt = timePtr(at)
	alert.Status = StatusResolved
}

func (s *Store) Acknowledge(id string, at time.Time) (*Alert, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	alert, ok := s.alerts[id]
	if !ok {
		return nil, ErrAlertNotFound
	}

	alert.Status = StatusAcknowledged
	alert.AcknowledgedAt = timePtr(at)
	alert.UpdatedAt = at

	return cloneAlert(alert), nil
}

func (s *Store) Silence(id string, at time.Time) (*Alert, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	alert, ok := s.alerts[id]
	if !ok {
		return nil, ErrAlertNotFound
	}

	alert.Status = StatusSilenced
	alert.SilencedAt = timePtr(at)
	alert.UpdatedAt = at

	return cloneAlert(alert), nil
}

func cloneAlert(src *Alert) *Alert {
	if src == nil {
		return nil
	}

	dst := *src

	if src.StartedAt != nil {
		dst.StartedAt = timePtr(*src.StartedAt)
	}
	if src.AcknowledgedAt != nil {
		dst.AcknowledgedAt = timePtr(*src.AcknowledgedAt)
	}
	if src.SilencedAt != nil {
		dst.SilencedAt = timePtr(*src.SilencedAt)
	}
	if src.ResolvedAt != nil {
		dst.ResolvedAt = timePtr(*src.ResolvedAt)
	}

	return &dst
}

func timePtr(value time.Time) *time.Time {
	ptr := new(time.Time)
	*ptr = value
	return ptr
}
