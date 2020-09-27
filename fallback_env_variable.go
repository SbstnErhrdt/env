package env

import (
	"log"
	"os"
)

func FallbackEnvVariable(envKey, fallbackValue string) string {
	result := os.Getenv(envKey)
	if len(result) == 0 {
		log.Println("Using fallback value", fallbackValue, "for env variable", envKey)
		return fallbackValue
	}
	return result
}
