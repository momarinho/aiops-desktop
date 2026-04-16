package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"aiops-desktop/backend/internal/config"
	"aiops-desktop/backend/internal/httpapi"
	"aiops-desktop/backend/internal/logger"
	"aiops-desktop/backend/internal/metrics"
)

func main() {
	// Configs
	cfg := config.Load()

	// Logger
	log := logger.New(cfg.LogLevel, cfg.Environment)
	log.Info("Starting server",
		"port", cfg.Port,
		"environment", cfg.Environment,
	)

	// Create store
	store := metrics.NewStore(100)

	// Create collector
	collector := metrics.NewCollector(log)

	// Create and start collector loop
	collectorLoop := metrics.NewCollectorLoop(collector, store, 2*time.Second)

	// Start em goroutine (background)
	go func() {
		if err := collectorLoop.Start(context.Background()); err != nil {
			log.Error("Collector loop failed", "error", err)
		}
	}()

	// Create router
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/metrics", metrics.Handler(store, log))
	mux.HandleFunc("/metrics/stream", metrics.StreamHandler(store, log))

	// Apply CORS middleware
	handler := httpapi.Middleware(mux)

	// Server init
	addr := ":" + cfg.Port
	log.Info("Server listening", "address", addr)

	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Error("Server failed", "error", err)
		panic(err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
}
