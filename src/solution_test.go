package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := map[string]string{
		"test": "grfg",
		"Test": "Grfg",
	}

	for key, element := range tests {
		output := rot13(key)

		if element != output {
			t.Errorf("Test for '%s' was incorrect, expected '%s', got '%s'", key, element, output)
		}
	}
}
