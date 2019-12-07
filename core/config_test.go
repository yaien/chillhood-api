package core

import "testing"

import "os"

func TestAddress(t *testing.T) {
	os.Setenv("ADDRESS", "")
	config := load()
	if config.Address != ":8080" {
		t.Errorf("expected server address to be :8080, received %s", config.Address)
	}
	addr := ":5000"
	os.Setenv("ADDRESS", addr)
	config = load()
	if config.Address != addr {
		t.Errorf("expected server address to be %s, received %s", addr, config.Address)
	}
}
