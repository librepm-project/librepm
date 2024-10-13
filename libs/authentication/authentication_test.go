package authentication

import (
	"testing"
)

func TestAuthentication(t *testing.T) {
	result := Authentication("works")
	if result != "Authentication works" {
		t.Error("Expected Authentication to append 'works'")
	}
}
