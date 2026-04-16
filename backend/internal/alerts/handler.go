package alerts

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"time"
)

func ListHandler(store *Store, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Debug("Listing alerts")
		writeJSON(w, http.StatusOK, store.List())
	}
}

func AcknowledgeHandler(store *Store, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		alert, err := store.Acknowledge(id, time.Now().UTC())
		if err != nil {
			if errors.Is(err, ErrAlertNotFound) {
				http.Error(w, "alert not found", http.StatusNotFound)
				return
			}
			logger.Error("Failed to acknowledge alert", "id", id, "error", err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		writeJSON(w, http.StatusOK, alert)
	}
}

func SilenceHandler(store *Store, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		alert, err := store.Silence(id, time.Now().UTC())
		if err != nil {
			if errors.Is(err, ErrAlertNotFound) {
				http.Error(w, "alert not found", http.StatusNotFound)
				return
			}
			logger.Error("Failed to silence alert", "id", id, "error", err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		writeJSON(w, http.StatusOK, alert)
	}
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}
