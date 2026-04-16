package metrics

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"
)

func Handler(logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			logger.Warn("Invalid method", "method", r.Method, "path", r.URL.Path)
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		logger.Debug("Metrics request received")

		response := MetricsResponse{
			Timestamp: time.Now().UTC().Format(time.RFC3339),
			Metrics: []Metric{
				{
					Name:  "cpu_usage_percent",
					Value: 45.7,
					Unit:  "percent",
					Labels: map[string]string{
						"host": "server-01",
					},
				},
				{
					Name:  "memory_usage_bytes",
					Value: 2147483648.0,
					Unit:  "bytes",
					Labels: map[string]string{
						"host": "server-01",
					},
				},
				{
					Name:  "request_rate",
					Value: 1234.5,
					Unit:  "requests_per_second",
					Labels: map[string]string{
						"endpoint": "/api/data",
						"method":   "GET",
					},
				},
			},
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(response); err != nil {
			logger.Error("Failed to encode metrics", "error", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}
}
