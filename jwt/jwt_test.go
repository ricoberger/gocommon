package jwt

import (
	"testing"
)

func TestCreate(t *testing.T) {
	_, err := Create(nil, "mysecret")
	if err != nil {
		t.Errorf("Could not create JWT token: %s", err.Error())
	}
}

func TestParse(t *testing.T) {
	_, err := Parse("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.bnVsbA.MkUUFKpbXwAZ9N6W9LDMMwByDne5Vmd8mM-0SKrjAo0", "mysecret")
	if err != nil {
		t.Errorf("Could not parse JWT token: %s", err.Error())
	}
}
