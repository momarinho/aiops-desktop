package ai

import (
	"fmt"
	"strings"
)

func BuildPrompt(input ExplainInput) string {
	var lines []string

	lines = append(lines,
		"You are an operations assistant. Explain the alert in concise, structured language.",
		fmt.Sprintf("Alert ID: %s", input.Alert.ID),
		fmt.Sprintf("Metric: %s", input.Alert.MetricName),
		fmt.Sprintf("Severity: %s", input.Alert.Severity),
		fmt.Sprintf("Status: %s", input.Alert.Status),
		fmt.Sprintf("Description: %s", input.Alert.Description),
		fmt.Sprintf("Current Value: %.2f", input.Alert.CurrentValue),
		fmt.Sprintf("Threshold: %.2f", input.Alert.Threshold),
	)

	if input.Host != "" {
		lines = append(lines, fmt.Sprintf("Observed Host: %s", input.Host))
	}

	if input.Context == nil {
		return strings.Join(lines, "\n")
	}

	if input.Context.Hostname != "" {
		lines = append(lines, fmt.Sprintf("Context Hostname: %s", input.Context.Hostname))
	}
	if input.Context.Service != "" {
		lines = append(lines, fmt.Sprintf("Context Service: %s", input.Context.Service))
	}
	if len(input.Context.RecentEvents) > 0 {
		lines = append(lines, fmt.Sprintf("Recent Events: %s", strings.Join(input.Context.RecentEvents, " | ")))
	}
	if len(input.Context.RecentActions) > 0 {
		lines = append(lines, fmt.Sprintf("Recent Actions: %s", strings.Join(input.Context.RecentActions, " | ")))
	}
	if input.Context.AdditionalNotes != "" {
		lines = append(lines, fmt.Sprintf("Additional Notes: %s", input.Context.AdditionalNotes))
	}

	return strings.Join(lines, "\n")
}
