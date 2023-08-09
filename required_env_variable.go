package env

import (
	"log/slog"
	"os"
)

// RequiredEnvVariable returns the value of the environment variable with the given key.
func RequiredEnvVariable(name string) string {
	variable := os.Getenv(name)
	if len(variable) == 0 {
		slog.With("env_variable", name).Error("required environment variable is not set")
		panic("required environment variable is not set")
	}
	return variable
}
