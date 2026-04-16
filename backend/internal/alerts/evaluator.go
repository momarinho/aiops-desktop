package alerts

import (
	"aiops-desktop/backend/internal/metrics"
	"sync"
)

type Evaluator struct {
	mu          sync.RWMutex
	store       *Store
	rules       []Rule
	consecutive map[string]int
}

func NewEvaluator(store *Store, rules []Rule) *Evaluator {
	return &Evaluator{
		store:       store,
		rules:       rules,
		consecutive: make(map[string]int),
	}
}

func (e *Evaluator) Evaluate(snapshot *metrics.Snapshot) {
	if snapshot == nil {
		return
	}

	values := make(map[string]float64, len(snapshot.Metrics))
	for _, metric := range snapshot.Metrics {
		values[metric.Name] = metric.Value
	}

	e.mu.Lock()
	defer e.mu.Unlock()

	for _, rule := range e.rules {
		value, ok := values[rule.MetricName]
		if !ok {
			continue
		}

		if value >= rule.Threshold {
			e.consecutive[rule.ID]++
			if e.consecutive[rule.ID] >= rule.Window {
				e.store.Activate(rule, value, snapshot.Timestamp)
			}
			continue
		}

		e.consecutive[rule.ID] = 0
		e.store.Resolve(rule.ID, value, snapshot.Timestamp)
	}
}
