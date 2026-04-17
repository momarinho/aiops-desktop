package actions

import (
	"fmt"
	"strconv"
)

type ActionDefinition struct {
	Type           ActionType
	Description    string
	RequiredParams []string
	Risky          bool
	Validate       func(target string, params map[string]any) error
}

func GetAllowlist() []ActionDefinition {
	return []ActionDefinition{
		{
			Type:           ActionTypeKillProcess,
			Description:    "Kill a process by PID",
			RequiredParams: []string{},
			Risky:          true,
			Validate:       validateKillProcess,
		},
		{
			Type:           ActionTypeRestartContainer,
			Description:    "Restart a Docker container",
			RequiredParams: []string{},
			Risky:          true,
			Validate:       validateContainerTarget,
		},
		{
			Type:           ActionTypeScaleContainer,
			Description:    "Scale a Docker service",
			RequiredParams: []string{"replicas"},
			Risky:          false,
			Validate:       validateScaleContainer,
		},
	}
}

func GetActionDefinition(actionType ActionType) (*ActionDefinition, bool) {
	for _, def := range GetAllowlist() {
		if def.Type == actionType {
			return &def, true
		}
	}

	return nil, false
}

func validateKillProcess(target string, _ map[string]any) error {
	if target == "" {
		return fmt.Errorf("PID cannot be empty")
	}

	pid, err := strconv.Atoi(target)
	if err != nil || pid <= 0 {
		return fmt.Errorf("invalid PID: %s is not a positive integer", target)
	}

	return nil
}

func validateContainerTarget(target string, _ map[string]any) error {
	if target == "" {
		return fmt.Errorf("container name/ID cannot be empty")
	}

	if len(target) == 64 {
		for _, c := range target {
			if !isHexChar(c) {
				return fmt.Errorf("invalid container ID: contains non-hex character '%c'", c)
			}
		}

		return nil
	}

	if len(target) > 128 {
		return fmt.Errorf("container name too long: maximum 128 characters")
	}

	firstChar := rune(target[0])
	if !isAlphaNumeric(firstChar) {
		return fmt.Errorf("container name must start with an alphanumeric character")
	}

	for _, c := range target[1:] {
		if !isAlphaNumeric(c) && c != '_' && c != '-' && c != '.' {
			return fmt.Errorf("invalid container name: contains invalid character '%c'", c)
		}
	}

	return nil
}

func validateScaleContainer(target string, params map[string]any) error {
	if err := validateContainerTarget(target, params); err != nil {
		return err
	}

	replicas, ok := params["replicas"]
	if !ok {
		return fmt.Errorf("missing required parameter: replicas")
	}

	switch value := replicas.(type) {
	case string:
		count, err := strconv.Atoi(value)
		if err != nil || count < 0 {
			return fmt.Errorf("replicas must be a non-negative integer")
		}
	case int:
		if value < 0 {
			return fmt.Errorf("replicas must be a non-negative integer")
		}
	case int32:
		if value < 0 {
			return fmt.Errorf("replicas must be a non-negative integer")
		}
	case int64:
		if value < 0 {
			return fmt.Errorf("replicas must be a non-negative integer")
		}
	case float64:
		if value < 0 || value != float64(int(value)) {
			return fmt.Errorf("replicas must be a non-negative integer")
		}
	default:
		return fmt.Errorf("replicas must be a string or number")
	}

	return nil
}

func isAlphaNumeric(r rune) bool {
	return (r >= '0' && r <= '9') ||
		(r >= 'a' && r <= 'z') ||
		(r >= 'A' && r <= 'Z')
}

func isHexChar(r rune) bool {
	return (r >= '0' && r <= '9') ||
		(r >= 'a' && r <= 'f') ||
		(r >= 'A' && r <= 'F')
}
