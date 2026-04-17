package processes

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
)

func ListHandler(monitor *Monitor, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Debug("Listing processes")

		processes, err := monitor.GetProcessList()
		if err != nil {
			logger.Error("Failed to get processes", "error", err)
			http.Error(w, "Failed to get processes", http.StatusInternalServerError)
			return
		}

		writeJSON(w, http.StatusOK, processes)
	}
}

func GetByPIDHandler(monitor *Monitor, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pidStr := r.PathValue("pid")
		pid, err := strconv.Atoi(pidStr)
		if err != nil {
			http.Error(w, "Invalid PID", http.StatusBadRequest)
			return
		}

		logger.Debug("Getting process by PID", "pid", pid)

		process, err := monitor.GetProcessByPID(pid)
		if err != nil {
			logger.Error("Process not found", "pid", pid, "error", err)
			http.Error(w, "Process not found", http.StatusNotFound)
			return
		}

		writeJSON(w, http.StatusOK, process)
	}
}

func SystemInfoHandler(monitor *Monitor, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Debug("Getting system info")

		info := monitor.GetSystemInfo()
		writeJSON(w, http.StatusOK, info)
	}
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}
