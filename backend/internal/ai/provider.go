package ai

import (
	"context"
	"errors"
	"fmt"
	"strings"
)

var ErrProviderUnavailable = errors.New("ai provider unavailable")

func NewProvider(name string) Provider {
	switch strings.ToLower(strings.TrimSpace(name)) {
	case "", "rule-based", "rule_based", "local":
		return RuleBasedProvider{}
	case "disabled", "off":
		return DisabledProvider{}
	default:
		return RuleBasedProvider{}
	}
}

type DisabledProvider struct{}

func (DisabledProvider) Name() string {
	return "disabled"
}

func (DisabledProvider) ExplainAlert(context.Context, string, ExplainInput) (*ExplainAlertResponse, error) {
	return nil, ErrProviderUnavailable
}

type RuleBasedProvider struct{}

func (RuleBasedProvider) Name() string {
	return "rule-based"
}

func (RuleBasedProvider) ExplainAlert(_ context.Context, _ string, input ExplainInput) (*ExplainAlertResponse, error) {
	summary := buildSummary(input)
	cause := buildProbableCause(input)
	actions := buildSuggestedActions(input)

	return &ExplainAlertResponse{
		Summary:          summary,
		ProbableCause:    cause,
		SuggestedActions: actions,
	}, nil
}

func buildSummary(input ExplainInput) string {
	severity := strings.ToUpper(string(input.Alert.Severity))
	metric := metricLabel(input.Alert.MetricName)
	base := fmt.Sprintf(
		"%s alert on %s: %s is at %.1f against a threshold of %.1f.",
		severity,
		scopeLabel(input),
		metric,
		input.Alert.CurrentValue,
		input.Alert.Threshold,
	)

	if input.Alert.Status == "acknowledged" || input.Alert.Status == "silenced" {
		return fmt.Sprintf("%s The alert is currently %s, so mitigation is already in progress.", base, input.Alert.Status)
	}

	return base
}

func buildProbableCause(input ExplainInput) string {
	switch input.Alert.MetricName {
	case "cpu_usage_percent":
		return "The host is under sustained compute pressure. A hot process, runaway worker, or insufficient capacity is the most likely cause."
	case "memory_usage_bytes":
		return "Memory pressure is elevated. This usually points to a leak, oversized workload, or a service holding memory longer than expected."
	case "disk_usage_bytes":
		return "Disk utilization is near the configured limit. Log growth, retained artifacts, or backlog files are the most likely drivers."
	default:
		return "The alert crossed its configured threshold and needs a quick operational review to confirm whether this is a transient spike or a sustained degradation."
	}
}

func buildSuggestedActions(input ExplainInput) []string {
	actions := []string{
		fmt.Sprintf("Validate the source of %s pressure on %s before making changes.", metricLabel(input.Alert.MetricName), scopeLabel(input)),
	}

	switch input.Alert.MetricName {
	case "cpu_usage_percent":
		actions = append(actions,
			"Inspect the top CPU-consuming processes or containers and compare usage against the time the alert started.",
			"Scale or restart the affected workload only after confirming the spike is sustained and user impact is increasing.",
		)
	case "memory_usage_bytes":
		actions = append(actions,
			"Review memory usage by process or container and look for recent deploys, cache growth, or leaks.",
			"Free pressure by restarting the affected workload or reducing concurrency if memory continues climbing.",
		)
	case "disk_usage_bytes":
		actions = append(actions,
			"Check which directories or volumes grew recently, especially logs, temp files, and retained artifacts.",
			"Reclaim space safely or expand capacity before the filesystem reaches a hard limit.",
		)
	default:
		actions = append(actions,
			"Confirm the metric trend from recent snapshots and compare it with recent operational changes.",
			"Choose the least risky mitigation first and record the action taken for follow-up.",
		)
	}

	if input.Context != nil && len(input.Context.RecentActions) > 0 {
		actions = append(actions, fmt.Sprintf("Review recent actions first: %s.", strings.Join(input.Context.RecentActions, ", ")))
	}

	return actions
}

func metricLabel(metricName string) string {
	switch metricName {
	case "cpu_usage_percent":
		return "CPU usage"
	case "memory_usage_bytes":
		return "memory usage"
	case "disk_usage_bytes":
		return "disk usage"
	default:
		return strings.ReplaceAll(metricName, "_", " ")
	}
}

func scopeLabel(input ExplainInput) string {
	if input.Context != nil && input.Context.Service != "" {
		return input.Context.Service
	}
	if input.Context != nil && input.Context.Hostname != "" {
		return input.Context.Hostname
	}
	if input.Host != "" {
		return input.Host
	}
	return "the monitored system"
}
