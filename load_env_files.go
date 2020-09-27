package env

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvFiles(filenames ...string) {
	// Load environment
	err := godotenv.Load(filenames...)
	if err != nil {
		log.Println("[WARNING] an error occurred while loading the .env files:", err)
	}
}
