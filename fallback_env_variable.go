package env

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func FallbackEnvVariable(envKey, fallbackValue string) string {
	result := os.Getenv(envKey)
	if len(result) == 0 {
		log.WithField("key", envKey).
			WithField("fallbackValue", fallbackValue).
			Warning("Using fallback")
		err := os.Setenv(envKey, fallbackValue)
		if err != nil {
			log.WithField("key", envKey).
				WithField("fallbackValue", fallbackValue).
				Fatal("can not set fallback variable")
		}
		return fallbackValue
	}
	return result
}
