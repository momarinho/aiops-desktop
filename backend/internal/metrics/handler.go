package metrics

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func Handler(store *Store, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			logger.Warn("Invalid method", "method", r.Method, "path", r.URL.Path)
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		logger.Debug("Metrics request received")

		latest := store.GetLatest()
		if latest == nil {
			logger.Warn("No metrics available")
			http.Error(w, "No metrics available", http.StatusServiceUnavailable)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(latest); err != nil {
			logger.Error("Failed to encode metrics", "error", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}
}
