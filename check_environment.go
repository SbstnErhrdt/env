package env

import (
	"errors"
	"log"
	"os"
)

// Checks if environment variables are present
func checkEnvironment(vars []string) (errs []error) {
	// iterate over necessary os vars
	for _, v := range vars {
		if len(os.Getenv(v)) == 0 {
			msg := "Environment variable " + v + " is not present."
			errs = append(errs, errors.New(msg))
		}
	}
	return
}

// Checks if environment variables are present
// panics if one is missing
func CheckRequiredEnvironmentVariables(requiredEnvironmentVariables []string) {
	errs := checkEnvironment(requiredEnvironmentVariables)
	// if necessary env variable: panic
	if len(errs) != 0 {
		for _, e := range errs {
			log.Println("[FATAL]", e, "Please add the variable to startup the system")
		}
		panic("Missing required environment variables")
	}
	return
}

// Checks if
func CheckOptionalEnvironmentVariables(possibleEnvironmentVariables []string) {
	errs := checkEnvironment(possibleEnvironmentVariables)
	// if necessary env variable: panic
	if len(errs) != 0 {
		for _, e := range errs {
			log.Println("[WARNING]", e)
		}
	}
	return
}
