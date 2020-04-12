package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	tests := map[uint64]string{}

	for key, element := range tests {
		output := ""

		if element != output {
			t.Errorf("Test for '%d' was incorrect, expected '%t', got '%t'", key, element, output)
		}
	}
}
