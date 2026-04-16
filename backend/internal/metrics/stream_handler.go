package metrics

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

func StreamHandler(store *Store, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Config headers SSE
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// create ticker
		ticker := time.NewTicker(2 * time.Second)
		defer ticker.Stop()

		// Send initial data
		if latest := store.GetLatest(); latest != nil {
			sendSSEEvent(w, latest, logger)
		}

		// loop
		for {
			select {
			case <-r.Context().Done():
				logger.Debug("Client disconnected from SSE")
				return
			case <-ticker.C:
				if latest := store.GetLatest(); latest != nil {
					if err := sendSSEEvent(w, latest, logger); err != nil {
						logger.Error("Failed to send SSE event", "error")
						return
					}
				}
			}
		}
	}
}

func sendSSEEvent(w http.ResponseWriter, snapshot *Snapshot, logger *slog.Logger) error {
	data, err := json.Marshal(snapshot)
	if err != nil {
		return err
	}

	// SSE Format
	_, err = fmt.Fprintf(w, "data: %s\n\n", data)
	if err != nil {
		return err
	}

	//Flush
	if flusher, ok := w.(http.Flusher); ok {
		flusher.Flush()
	}

	return nil
}
