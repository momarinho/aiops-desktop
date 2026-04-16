package metrics

import (
	"log/slog"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

type Collector struct {
	logger       *slog.Logger
	hostInfo     *host.InfoStat
	lastNetStats map[string]net.IOCountersStat
}

func NewCollector(logger *slog.Logger) *Collector {
	hostInfo, err := host.Info()
	if err != nil {
		logger.Warn("Failed to get host info", "error", err)
	}

	return &Collector{
		logger:       logger,
		hostInfo:     hostInfo,
		lastNetStats: make(map[string]net.IOCountersStat),
	}
}

func (c *Collector) CollectCPU() (float64, error) {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return 0, err
	}
	if len(percent) == 0 {
		return 0, nil
	}
	return percent[0], nil
}

func (c *Collector) CollectMemory() (usedBytes, usedPercent float64, err error) {
	stats, err := mem.VirtualMemory()
	if err != nil {
		return 0, 0, err
	}
	return float64(stats.Used), stats.UsedPercent, nil
}

func (c *Collector) CollectDisk() (usedBytes, usedPercent float64, err error) {
	stats, err := disk.Usage("/")
	if err != nil {
		return 0, 0, err
	}
	return float64(stats.Used), stats.UsedPercent, nil
}

func (c *Collector) CollectNetwork() (tx, rx float64, err error) {
	stats, err := net.IOCounters(false)
	if err != nil {
		return 0, 0, err
	}

	if len(stats) == 0 {
		return 0, 0, nil
	}

	var totalTX, totalRX float64
	for _, stat := range stats {
		totalTX += float64(stat.BytesSent)
		totalRX += float64(stat.BytesRecv)
	}

	return totalTX, totalRX, nil
}

func (c *Collector) CollectAll() (*SystemMetrics, error) {
	cpuVal, err := c.CollectCPU()
	if err != nil {
		c.logger.Warn("Failed to collect CPU", "error", err)
	}

	memVal, memPercent, err := c.CollectMemory()
	if err != nil {
		c.logger.Warn("Failed to collect memory", "error", err)
	}

	diskVal, diskPercent, err := c.CollectDisk()
	if err != nil {
		c.logger.Warn("Failed to collect disk", "error", err)
	}

	txVal, rxVal, err := c.CollectNetwork()
	if err != nil {
		c.logger.Warn("Failed to collect network", "error", err)
	}

	return &SystemMetrics{
		CPU:           cpuVal,
		Memory:        memVal,
		MemoryPercent: memPercent,
		Disk:          diskVal,
		DiskPercent:   diskPercent,
		NetworkTX:     txVal,
		NetworkRX:     rxVal,
	}, nil
}
