package metrics

type Metric struct {
	Name   string            `json:"name"`
	Value  float64           `json:"value"`
	Unit   string            `json:"unit"`
	Labels map[string]string `json:"labels,omitempty"`
}

type MetricsResponse struct {
	Timestamp string   `json:"timestamp"`
	Metrics   []Metric `json:"metrics"`
}
