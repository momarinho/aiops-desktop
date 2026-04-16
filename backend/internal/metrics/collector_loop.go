package metrics

import (
	"context"
	"time"
)

type CollectorLoop struct {
	collector  *Collector
	store      *Store
	interval   time.Duration
	onSnapshot func(*Snapshot)
}

func NewCollectorLoop(
	collector *Collector,
	store *Store,
	interval time.Duration,
	onSnapshot func(*Snapshot),
) *CollectorLoop {
	if interval <= 0 {
		interval = 2 * time.Second
	}

	return &CollectorLoop{
		collector:  collector,
		store:      store,
		interval:   interval,
		onSnapshot: onSnapshot,
	}
}

func (cl *CollectorLoop) Start(ctx context.Context) error {
	if err := cl.collectAndStore(); err != nil {
		return err
	}

	ticker := time.NewTicker(cl.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		case <-ticker.C:
			if err := cl.collectAndStore(); err != nil {
			}
		}
	}
}

func (cl *CollectorLoop) collectAndStore() error {
	systemMetrics, err := cl.collector.CollectAll()
	if err != nil {
		return err
	}

	now := time.Now().UTC()

	snapshot := &Snapshot{
		Timestamp: now,
		Metrics: []Metric{
			{
				Type:      MetricTypeCPU,
				Name:      "cpu_usage_percent",
				Value:     systemMetrics.CPU,
				Unit:      "percent",
				Timestamp: now,
			},
			{
				Type:      MetricTypeMemory,
				Name:      "memory_usage_bytes",
				Value:     systemMetrics.Memory,
				Unit:      "bytes",
				Timestamp: now,
			},
			{
				Type:      MetricTypeMemory,
				Name:      "memory_usage_percent",
				Value:     systemMetrics.MemoryPercent,
				Unit:      "percent",
				Timestamp: now,
			},
			{
				Type:      MetricTypeDisk,
				Name:      "disk_usage_bytes",
				Value:     systemMetrics.Disk,
				Unit:      "bytes",
				Timestamp: now,
			},
			{
				Type:      MetricTypeDisk,
				Name:      "disk_usage_percent",
				Value:     systemMetrics.DiskPercent,
				Unit:      "percent",
				Timestamp: now,
			},
			{
				Type:      MetricTypeNetwork,
				Name:      "network_tx_bytes",
				Value:     systemMetrics.NetworkTX,
				Unit:      "bytes",
				Timestamp: now,
			},
			{
				Type:      MetricTypeNetwork,
				Name:      "network_rx_bytes",
				Value:     systemMetrics.NetworkRX,
				Unit:      "bytes",
				Timestamp: now,
			},
		},
	}

	cl.store.Put(snapshot)

	if cl.onSnapshot != nil {
		cl.onSnapshot(snapshot)
	}

	return nil
}
