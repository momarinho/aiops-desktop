package ai

import (
	"context"
	"time"

	"aiops-desktop/backend/internal/alerts"
)

type ExplainAlertRequest struct {
	AlertID string               `json:"alert_id"`
	Context *ExplainAlertContext `json:"context,omitempty"`
}

type ExplainAlertContext struct {
	Hostname        string   `json:"hostname,omitempty"`
	Service         string   `json:"service,omitempty"`
	RecentEvents    []string `json:"recent_events,omitempty"`
	RecentActions   []string `json:"recent_actions,omitempty"`
	AdditionalNotes string   `json:"additional_notes,omitempty"`
}

type ExplainAlertResponse struct {
	Summary          string   `json:"summary"`
	ProbableCause    string   `json:"probable_cause"`
	SuggestedActions []string `json:"suggested_actions"`
}

type ExplainInput struct {
	Alert      alerts.Alert
	Context    *ExplainAlertContext
	Host       string
	ReceivedAt time.Time
}

type Provider interface {
	Name() string
	ExplainAlert(ctx context.Context, prompt string, input ExplainInput) (*ExplainAlertResponse, error)
}

type ErrorResponse struct {
	Error string `json:"error"`
}
