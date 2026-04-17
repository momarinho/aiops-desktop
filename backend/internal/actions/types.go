package actions

import "time"

type ActionType string
type ActionStatus string

const (
	ActionTypeKillProcess      ActionType = "kill_process"
	ActionTypeRestartContainer ActionType = "restart_container"
	ActionTypeScaleContainer   ActionType = "scale_container"

	ActionStatusPending ActionStatus = "pending"
	ActionStatusSuccess ActionStatus = "success"
	ActionStatusFailed  ActionStatus = "failed"
)

type Action struct {
	ID          string         `json:"id"`
	Type        ActionType     `json:"type"`
	Target      string         `json:"target"`
	Parameters  map[string]any `json:"parameters"`
	User        string         `json:"user"`
	RequestTime time.Time      `json:"request_time"`
	StartTime   *time.Time     `json:"start_time,omitempty"`
	EndTime     *time.Time     `json:"end_time,omitempty"`
	Status      ActionStatus   `json:"status"`
	Output      string         `json:"output,omitempty"`
	Error       string         `json:"error,omitempty"`
	Risky       bool           `json:"risky"`
}
