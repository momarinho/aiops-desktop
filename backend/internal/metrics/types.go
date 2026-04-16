package metrics

import "time"

type MetricType string

const (
	MetricTypeCPU     MetricType = "cpu"
	MetricTypeMemory  MetricType = "memory"
	MetricTypeDisk    MetricType = "disk"
	MetricTypeNetwork MetricType = "network"
)

type Metric struct {
	Type      MetricType        `json:"type"`
	Name      string            `json:"name"`
	Value     float64           `json:"value"`
	Unit      string            `json:"unit"`
	Timestamp time.Time         `json:"timestamp"`
	Labels    map[string]string `json:"labels,omitempty"`
}

type Snapshot struct {
	Timestamp time.Time `json:"timestamp"`
	Metrics   []Metric  `json:"metrics"`
}

type SystemMetrics struct {
	CPU           float64 `json:"cpu_percent"`
	Memory        float64 `json:"memory_used_bytes"`
	MemoryPercent float64 `json:"memory_percent"`
	Disk          float64 `json:"disk_used_bytes"`
	DiskPercent   float64 `json:"disk_percent"`
	NetworkTX     float64 `json:"network_tx_bytes"`
	NetworkRX     float64 `json:"network_rx_bytes"`
}
