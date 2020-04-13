package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	tests := map[string]string{
		"test": "grfg",
		"Test": "Grfg",
	}

	for key, element := range tests {
		output := ""

		if element != output {
			t.Errorf("Test for '%d' was incorrect, expected '%t', got '%t'", key, element, output)
		}
	}
}
