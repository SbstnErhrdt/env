package env

import (
	"os"
	"testing"
)

func TestRequiredEnvVariable(t *testing.T) {

	err := os.Setenv("city", "london")
	if err != nil {
		t.Error(err)
		return
	}
	variable := RequiredEnvVariable("city")
	if variable != "london" {
		t.Error("variable is not london")
		return
	}
}

func TestRequiredEnvVariablePanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	// unset_env_variable is not set in the environment
	// so this should panic
	_ = RequiredEnvVariable("unset_env_variable")
}
