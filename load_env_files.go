package env

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func LoadEnvFiles(filenames ...string) {
	// Load environment
	err := godotenv.Load(filenames...)
	if err != nil {
		log.Warning("an error occurred while loading the .env files:", err)
	}
}
