package actions

import (
	"context"
	"fmt"
	"log/slog"
	"os/exec"
	"strconv"
	"time"
)

type Executor struct {
	logger    *slog.Logger
	allowlist []ActionDefinition
}

func NewExecutor(logger *slog.Logger) *Executor {
	return &Executor{
		logger:    logger,
		allowlist: GetAllowlist(),
	}
}

func (e *Executor) Execute(ctx context.Context, action *Action) error {
	// validate allowlist action
	def, ok := GetActionDefinition(action.Type)
	if !ok {
		return fmt.Errorf("action type not allowed: %s", action.Type)
	}

	// validate target and parameteres
	if err := def.Validate(action.Target, action.Parameters); err != nil {
		return fmt.Errorf("validation failed %w", err)
	}

	// Mark init
	now := time.Now().UTC()
	action.StartTime = &now
	action.Status = ActionStatusPending

	// Exec action
	var output string
	var err error

	switch action.Type {
	case ActionTypeKillProcess:
		output, err = e.executeKillProcess(ctx, action.Target)
	case ActionTypeRestartContainer:
		output, err = e.executeRestartContainer(ctx, action.Target)
	case ActionTypeScaleContainer:
		output, err = e.executeScaleContainer(ctx, action.Target, action.Parameters)
	default:
		err = fmt.Errorf("unimplemented action type %s", action.Type)
	}

	// mark action as done
	endTime := time.Now().UTC()
	action.EndTime = &endTime

	if err != nil {
		action.Status = ActionStatusFailed
		action.Error = err.Error()
		e.logger.Error("action failed",
			"action_type", action.Type,
			"target", action.Target,
			"error", err)
		return err
	}

	action.Status = ActionStatusSuccess
	action.Output = output
	e.logger.Info("action succeeded",
		"action_type", action.Type,
		"target", action.Target)

	return nil
}

func (e *Executor) executeKillProcess(ctx context.Context, pid string) (string, error) {
	// First, check if process exists
	checkCmd := exec.CommandContext(ctx, "kill", "-0", pid)
	if err := checkCmd.Run(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() == 1 {
			return "", fmt.Errorf("process %s does not exist", pid)
		}
		return "", fmt.Errorf("failed to check process %s: %w", pid, err)
	}

	// Process exists, try to kill it
	cmd := exec.CommandContext(ctx, "kill", pid)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to kill process %s: %w (output: %s)", pid, err, string(output))
	}
	return string(output), nil
}

func (e *Executor) executeRestartContainer(ctx context.Context, containerID string) (string, error) {
	// Use docker restart command
	cmd := exec.CommandContext(ctx, "docker", "restart", containerID)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to restart container %s: %w", containerID, err)
	}
	return string(output), nil
}

func (e *Executor) executeScaleContainer(ctx context.Context, service string, params map[string]any) (string, error) {
	// Get replicas from parameters
	replicas, ok := params["replicas"]
	if !ok {
		return "", fmt.Errorf("missing required parameter: replicas")
	}

	replicasStr, ok := replicas.(string)
	if !ok {
		switch value := replicas.(type) {
		case int:
			replicasStr = strconv.Itoa(value)
		case int32:
			replicasStr = strconv.FormatInt(int64(value), 10)
		case int64:
			replicasStr = strconv.FormatInt(value, 10)
		case float64:
			if value != float64(int(value)) {
				return "", fmt.Errorf("replicas must be an integer")
			}
			replicasStr = strconv.Itoa(int(value))
		default:
			return "", fmt.Errorf("replicas must be a string or number")
		}
	}

	// Use docker scale command
	cmd := exec.CommandContext(ctx, "docker", "service", "scale", fmt.Sprintf("%s=%s", service, replicasStr))
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to scale service %s: %w", service, err)
	}
	return string(output), nil
}
