package core

import "testing"

import "os"

func TestAddress(t *testing.T) {
	os.Setenv("ADDRESS", "")
	addr := address()
	if addr != ":8080" {
		t.Errorf("expected server address to be :8080, received %s", addr)
	}
	envAddr := ":5000"
	os.Setenv("ADDRESS", envAddr)
	addr = address()
	if addr != envAddr {
		t.Errorf("expected server address to be %s, received %s", envAddr, addr)
	}
}
