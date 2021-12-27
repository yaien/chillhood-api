package infrastructure

import (
	"os"
	"testing"
)

func TestAddress(t *testing.T) {
	os.Setenv("PORT", "")
	addr := address()
	if addr != ":8080" {
		t.Errorf("expected server address to be :8080, received %s", addr)
	}
	port := "5000"
	os.Setenv("PORT", port)
	addr = address()
	if addr != ":"+port {
		t.Errorf("expected server address to be %s, received %s", port, addr)
	}
}
