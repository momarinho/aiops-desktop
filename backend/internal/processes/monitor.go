package processes

import (
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/v3/process"
)

type Monitor struct {
	logger *slog.Logger
}

func NewMonitor(logger *slog.Logger) *Monitor {
	return &Monitor{
		logger: logger,
	}
}

func (m *Monitor) GetProcessList() ([]ProcessInfo, error) {
	processes, err := process.Processes()
	if err != nil {
		m.logger.Error("Failed to get process list", "error", err)
		return nil, fmt.Errorf("failed to get process list: %w", err)
	}

	var result []ProcessInfo
	now := time.Now()

	for _, p := range processes {
		pid := p.Pid

		// Get process name
		name, err := p.Name()
		if err != nil {
			name = "unknown"
		}

		// Get username
		username, err := p.Username()
		if err != nil {
			username = "unknown"
		}

		// Get CPU percent
		cpuPercent, err := p.CPUPercent()
		if err != nil {
			cpuPercent = 0
		}

		// Get memory info
		memInfo, err := p.MemoryInfo()
		var memoryMB float64
		if err == nil && memInfo != nil {
			memoryMB = float64(memInfo.RSS) / 1024 / 1024
		}

		// Get create time
		createTime, err := p.CreateTime()
		var createTimeStr string
		if err == nil {
			ct := time.Unix(createTime/1000, 0)
			createTimeStr = ct.Format(time.RFC3339)
		} else {
			createTimeStr = now.Format(time.RFC3339)
		}

		// Get command line
		cmdline, err := p.Cmdline()
		if err != nil {
			cmdline = name
		}

		// Check if critical
		isCritical := IsCriticalPID(int(pid))

		// Get status
		status := GetProcessStatus(int(pid))

		processInfo := ProcessInfo{
			PID:        int(pid),
			Name:       name,
			User:       username,
			CPUPercent: cpuPercent,
			MemoryMB:   memoryMB,
			CreateTime: createTimeStr,
			Command:    cmdline,
			IsCritical: isCritical,
			Status:     status,
		}

		result = append(result, processInfo)
	}

	// Sort by CPU usage (highest first)
	for i := 0; i < len(result)-1; i++ {
		for j := i + 1; j < len(result); j++ {
			if result[j].CPUPercent > result[i].CPUPercent {
				result[i], result[j] = result[j], result[i]
			}
		}
	}

	m.logger.Debug("Retrieved process list", "count", len(result))

	return result, nil
}

func (m *Monitor) GetProcessByPID(pid int) (*ProcessInfo, error) {
	processes, err := m.GetProcessList()
	if err != nil {
		return nil, err
	}

	for _, p := range processes {
		if p.PID == pid {
			return &p, nil
		}
	}

	return nil, fmt.Errorf("process with PID %d not found", pid)
}

func (m *Monitor) GetSystemInfo() map[string]interface{} {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	info := map[string]interface{}{
		"hostname":    "",
		"os":          runtime.GOOS,
		"arch":        runtime.GOARCH,
		"cpu_count":   runtime.NumCPU(),
		"go_version":  runtime.Version(),
		"memory_mb":   memStats.Sys / 1024 / 1024,
		"timestamp":   time.Now().Format(time.RFC3339),
	}

	if hostname, err := os.Hostname(); err == nil {
		info["hostname"] = hostname
	}

	return info
}
