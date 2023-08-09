package env

import (
	"log/slog"
	"os"
)

// FallbackEnvVariable returns the value of the environment variable with the given key.
func FallbackEnvVariable(envKey, fallbackValue string) string {
	result := os.Getenv(envKey)
	if len(result) == 0 {
		slog.With("key", envKey).
			With("fallbackValue", fallbackValue).
			Warn("using fallback")
		err := os.Setenv(envKey, fallbackValue)
		if err != nil {
			slog.With("key", envKey).
				With("fallbackValue", fallbackValue).
				Error("can not set fallback variable")
			panic(err)
		}
		return fallbackValue
	}
	return result
}
