package env

import (
	"os"
	"testing"
)

func TestFallbackEnvVariable(t *testing.T) {
	res := FallbackEnvVariable("dog", "cat")
	if res != "cat" {
		t.Error("Wrong env variable")
	}
}

func TestFallbackEnvVariableSet(t *testing.T) {
	err := os.Setenv("dog", "dog")
	if err != nil {
		t.Error("can not set env variable")
	}
	res := FallbackEnvVariable("dog", "cat")
	if res != "dog" {
		t.Error("Wrong env variable")
	}
}
