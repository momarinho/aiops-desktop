package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port             string
	LogLevel         string
	Environment      string
	AIProvider       string
	AITimeoutSeconds int
}

func Load() *Config {
	return &Config{
		Port:             getEnv("PORT", "8080"),
		LogLevel:         getEnv("LOG_LEVEL", "info"),
		Environment:      getEnv("ENVIRONMENT", "development"),
		AIProvider:       getEnv("AI_PROVIDER", "rule-based"),
		AITimeoutSeconds: getEnvAsInt("AI_TIMEOUT_SECONDS", 8),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}
