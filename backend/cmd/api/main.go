package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"aiops-desktop/backend/internal/actions"
	"aiops-desktop/backend/internal/ai"
	"aiops-desktop/backend/internal/alerts"
	"aiops-desktop/backend/internal/config"
	"aiops-desktop/backend/internal/httpapi"
	"aiops-desktop/backend/internal/logger"
	"aiops-desktop/backend/internal/metrics"
	"aiops-desktop/backend/internal/processes"
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
	alertStore := alerts.NewStore()
	alertEvaluator := alerts.NewEvaluator(alertStore, alerts.DefaultRules())
	actionStore := actions.NewStore()
	actionExecutor := actions.NewExecutor(log)
	processMonitor := processes.NewMonitor(log)
	hostname, _ := os.Hostname()
	aiService := ai.NewService(
		alertStore,
		ai.NewProvider(cfg.AIProvider),
		time.Duration(cfg.AITimeoutSeconds)*time.Second,
		log,
		hostname,
	)

	// Create collector
	collector := metrics.NewCollector(log)

	// Create and start collector loop
	collectorLoop := metrics.NewCollectorLoop(collector, store, 2*time.Second, alertEvaluator.Evaluate)

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
	mux.HandleFunc("GET /alerts", alerts.ListHandler(alertStore, log))
	mux.HandleFunc("GET /alerts/{id}", alerts.GetByIDHandler(alertStore, log))
	mux.HandleFunc("POST /alerts/{id}/acknowledge", alerts.AcknowledgeHandler(alertStore, log))
	mux.HandleFunc("POST /alerts/{id}/silence", alerts.SilenceHandler(alertStore, log))
	mux.HandleFunc("POST /ai/explain-alert", aiService.ExplainAlertHandler())
	mux.HandleFunc("POST /actions", actions.ExecuteHandler(actionStore, actionExecutor, log))
	mux.HandleFunc("GET /actions", actions.ListHandler(actionStore, log))
	mux.HandleFunc("GET /actions/{id}", actions.GetByIDHandler(actionStore, log))
	mux.HandleFunc("GET /processes", processes.ListHandler(processMonitor, log))
	mux.HandleFunc("GET /processes/{pid}", processes.GetByPIDHandler(processMonitor, log))
	mux.HandleFunc("GET /system/info", processes.SystemInfoHandler(processMonitor, log))

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
