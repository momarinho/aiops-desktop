package alerts

func DefaultRules() []Rule {
	return []Rule{
		{
			ID:          "cpu-high",
			MetricName:  "cpu_usage_percent",
			Threshold:   85,
			Window:      3,
			Severity:    SeverityCritical,
			Description: "CPU usage is above 85% for 3 consecutive samples",
		},
		{
			ID:          "memory-high",
			MetricName:  "memory_usage_percent",
			Threshold:   90,
			Window:      3,
			Severity:    SeverityCritical,
			Description: "Memory usage is above 90% for 3 consecutive samples",
		},
		{
			ID:          "disk-high",
			MetricName:  "disk_usage_percent",
			Threshold:   90,
			Window:      3,
			Severity:    SeverityWarning,
			Description: "Disk usage is above 90% for 3 consecutive samples",
		},
	}
}
