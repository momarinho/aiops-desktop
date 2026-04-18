package ai

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"aiops-desktop/backend/internal/alerts"
)

type Service struct {
	alertStore *alerts.Store
	provider   Provider
	timeout    time.Duration
	logger     *slog.Logger
	host       string
}

func NewService(alertStore *alerts.Store, provider Provider, timeout time.Duration, logger *slog.Logger, host string) *Service {
	return &Service{
		alertStore: alertStore,
		provider:   provider,
		timeout:    timeout,
		logger:     logger,
		host:       host,
	}
}

func (s *Service) ExplainAlertHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req ExplainAlertRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "invalid request body"})
			return
		}

		if strings.TrimSpace(req.AlertID) == "" {
			writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "alert_id is required"})
			return
		}

		alert, found := s.alertStore.GetByID(req.AlertID)
		if !found {
			writeJSON(w, http.StatusNotFound, ErrorResponse{Error: "alert not found"})
			return
		}

		input := ExplainInput{
			Alert:      alert,
			Context:    req.Context,
			Host:       s.host,
			ReceivedAt: time.Now().UTC(),
		}
		prompt := BuildPrompt(input)

		ctx := r.Context()
		cancel := func() {}
		if s.timeout > 0 {
			ctx, cancel = context.WithTimeout(r.Context(), s.timeout)
		}
		defer cancel()

		response, err := s.provider.ExplainAlert(ctx, prompt, input)
		if err != nil {
			switch {
			case errors.Is(err, context.DeadlineExceeded):
				writeJSON(w, http.StatusGatewayTimeout, ErrorResponse{Error: "ai explanation timed out"})
			case errors.Is(err, ErrProviderUnavailable):
				writeJSON(w, http.StatusServiceUnavailable, ErrorResponse{Error: "ai provider unavailable"})
			default:
				s.logger.Error("Failed to explain alert", "alert_id", req.AlertID, "error", err)
				writeJSON(w, http.StatusInternalServerError, ErrorResponse{Error: "failed to explain alert"})
			}
			return
		}

		writeJSON(w, http.StatusOK, normalizeResponse(response))
	}
}

func normalizeResponse(response *ExplainAlertResponse) *ExplainAlertResponse {
	if response == nil {
		return &ExplainAlertResponse{
			Summary:       "Alert context was analyzed, but no summary was returned.",
			ProbableCause: "The alert exceeded its configured threshold and needs operator review.",
			SuggestedActions: []string{
				"Review the current metric trend and confirm the alert is still active.",
				"Inspect the affected workload before applying a restart or scale operation.",
				"Record the mitigation outcome in the action history.",
			},
		}
	}

	normalized := &ExplainAlertResponse{
		Summary:          strings.TrimSpace(response.Summary),
		ProbableCause:    strings.TrimSpace(response.ProbableCause),
		SuggestedActions: make([]string, 0, len(response.SuggestedActions)),
	}

	if normalized.Summary == "" {
		normalized.Summary = "Alert context was analyzed, but the explanation provider returned an empty summary."
	}
	if normalized.ProbableCause == "" {
		normalized.ProbableCause = "The alert exceeded its configured threshold and needs operator review."
	}

	for _, action := range response.SuggestedActions {
		trimmed := strings.TrimSpace(action)
		if trimmed == "" {
			continue
		}
		normalized.SuggestedActions = append(normalized.SuggestedActions, trimmed)
	}

	if len(normalized.SuggestedActions) == 0 {
		normalized.SuggestedActions = []string{
			"Review the current metric trend and confirm the alert is still active.",
			"Inspect the affected workload before applying a restart or scale operation.",
			"Record the mitigation outcome in the action history.",
		}
	}

	if len(normalized.SuggestedActions) > 4 {
		normalized.SuggestedActions = normalized.SuggestedActions[:4]
	}

	return normalized
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}
