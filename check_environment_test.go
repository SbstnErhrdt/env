package env

import (
	"log"
	"testing"
)

func TestCheckOptionalEnvironmentVariables(t *testing.T) {
	LoadEnvFiles("test0.env", "test1.env")
	CheckOptionalEnvironmentVariables([]string{"city", "country"})
}

func TestCheckRequiredEnvironmentVariables(t *testing.T) {
	LoadEnvFiles("test0.env", "test1.env")
	CheckRequiredEnvironmentVariables([]string{"city", "alan"})
}

func TestCheckRequiredEnvironmentVariables2(t *testing.T) {
	defer func() {
		// recover from panic if one occurred.
		if recover() != nil {
			log.Println("its fine")
			return
		}
	}()
	CheckRequiredEnvironmentVariables([]string{"turing"})
}
