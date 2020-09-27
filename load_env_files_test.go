package env

import (
	"os"
	"testing"
)

func TestLoadEnvFiles(t *testing.T) {
	LoadEnvFiles("test0.env", "test1.env")
	if os.Getenv("alan") != "turing" {
		t.Error("Could not read test var")
	}
	if os.Getenv("city") != "munich" {
		t.Error("Could not read test var")
	}
}