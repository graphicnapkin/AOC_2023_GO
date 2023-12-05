package main

import (
	"os"
	"testing"
)

func TestSolution(t *testing.T) {
	data, _ := os.ReadFile("test_input.txt")
	testInput := string(data)

	tests := []struct {
		name           string
		input          string
		function 		 func(string) int
		expectedOutput int
	}{
		{
			name:           "Part 1",
			input:          testInput,
			function: 		 solution,
			expectedOutput: 13,
		},
		{
			name:           "Part 2",
			input:          testInput,
			function: 		 solutionPart2,
			expectedOutput: 30,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualOutput := test.function(test.input)
			if actualOutput != test.expectedOutput {
				t.Errorf("Failed: expected %d, got %d", test.expectedOutput, actualOutput)
			}
		})
	}
}
