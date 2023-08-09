package env

import (
	"log/slog"
)

func LoadEnvFiles(filenames ...string) {
	// Load environment
	err := Load(filenames...)
	if err != nil {
		slog.With("err", err).Warn("an error occurred while loading the .env files")
	}
}
