package src

import (
	"os"
	"testing"
)

func TestSecretAccess(t *testing.T) {
	val := os.Getenv("OP_SERVICE_ACCOUNT_TOKEN")
	if val == "" {
		panic("no access to secrets")
	}
}
