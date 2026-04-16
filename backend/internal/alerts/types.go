package alerts

import "time"

type Severity string
type Status string

const (
	SeverityWarning  Severity = "warning"
	SeverityCritical Severity = "critical"

	StatusFiring       Status = "firing"
	StatusAcknowledged Status = "acknowledged"
	StatusSilenced     Status = "silenced"
	StatusResolved     Status = "resolved"
)

type Rule struct {
	ID          string
	MetricName  string
	Threshold   float64
	Window      int
	Severity    Severity
	Description string
}

type Alert struct {
	ID             string     `json:"id"`
	Severity       Severity   `json:"severity"`
	Status         Status     `json:"status"`
	Description    string     `json:"description"`
	MetricName     string     `json:"metric_name"`
	Threshold      float64    `json:"threshold"`
	CurrentValue   float64    `json:"current_value"`
	StartedAt      *time.Time `json:"started_at,omitempty"`
	UpdatedAt      time.Time  `json:"updated_at"`
	AcknowledgedAt *time.Time `json:"acknowledged_at,omitempty"`
	SilencedAt     *time.Time `json:"silenced_at,omitempty"`
	ResolvedAt     *time.Time `json:"resolved_at,omitempty"`
}
