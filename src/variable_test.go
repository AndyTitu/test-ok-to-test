package src

import (
	"os"
	"testing"
)

func TestSecretAccess(t *testing.T) {
	val := os.Getenv("PERSONAL_ACCESS_TOKEN")
	if val == "" {
		panic("no access to secrets")
	}
}
