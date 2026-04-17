package actions

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// ExecuteActionRequest represents the request body for executing an action
type ExecuteActionRequest struct {
	Type       ActionType     `json:"type"`
	Target     string         `json:"target"`
	Parameters map[string]any `json:"parameters"`
	User       string         `json:"user"`
}

// ListHandler returns all actions in history
func ListHandler(store *Store, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Debug("Listing actions")
		writeJSON(w, http.StatusOK, store.GetAll())
	}
}

// GetByIDHandler returns a specific action by ID
func GetByIDHandler(store *Store, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		action, found := store.GetByID(id)
		if !found {
			http.Error(w, "action not found", http.StatusNotFound)
			return
		}

		writeJSON(w, http.StatusOK, action)
	}
}

// ExecuteHandler executes a new action
func ExecuteHandler(store *Store, executor *Executor, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse request body
		var req ExecuteActionRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			logger.Error("Failed to decode request", "error", err)
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		// Validate required fields
		if req.Type == "" {
			http.Error(w, "action type is required", http.StatusBadRequest)
			return
		}
		if req.Target == "" {
			http.Error(w, "action target is required", http.StatusBadRequest)
			return
		}

		// Check if action is allowed and get its definition
		actionDef, ok := GetActionDefinition(req.Type)
		if !ok {
			http.Error(w, "action type not allowed", http.StatusForbidden)
			return
		}

		// Validate target and parameters
		if err := actionDef.Validate(req.Target, req.Parameters); err != nil {
			logger.Warn("Action validation failed", "type", req.Type, "error", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Check if action is risky (could add additional confirmation logic here)
		if actionDef.Risky {
			logger.Info("Executing risky action", "type", req.Type, "target", req.Target, "user", req.User)
		}

		// Create action
		now := time.Now().UTC()
		action := &Action{
			ID:          uuid.New().String(),
			Type:        req.Type,
			Target:      req.Target,
			Parameters:  req.Parameters,
			User:        req.User,
			RequestTime: now,
			Status:      ActionStatusPending,
			Risky:       actionDef.Risky,
		}

		// Add to store before executing
		store.Add(action)

		// Execute action asynchronously
		go func() {
			// Create new context for background execution (not tied to HTTP request)
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
			defer cancel()

			if err := executor.Execute(ctx, action); err != nil {
				logger.Error("Action execution failed", "id", action.ID, "error", err)
			}
		}()

		// Return the created action (it will be pending initially)
		writeJSON(w, http.StatusAccepted, action)
	}
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}
