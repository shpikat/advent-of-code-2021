package internal

import (
	"os"
	"testing"
)

func ReadInput(t *testing.T, filename string) string {
	input, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("Error reading file %s: %v", filename, err)
	}
	return string(input)
}
