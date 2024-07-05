package directory

import (
	"os"
	"testing"
)

func TestGetCurrent(t *testing.T) {
	expected, err := os.Getwd()
	if err != nil {
		t.Fatalf("os.Getwd() failed: %v", err)
	}

	actual := GetCurrent()
	if actual != expected {
		t.Errorf("GetCurrent() = %v; want %v", actual, expected)
	}
}
