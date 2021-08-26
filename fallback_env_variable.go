package env

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func FallbackEnvVariable(envKey, fallbackValue string) string {
	result := os.Getenv(envKey)
	if len(result) == 0 {
		log.Warning("Using fallback value", fallbackValue, "for env variable", envKey)
		err := os.Setenv(envKey, fallbackValue)
		if err != nil {
			log.Fatal("can not set fallback variable", envKey, fallbackValue)
		}
		return fallbackValue
	}
	return result
}
