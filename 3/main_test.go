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
		expectedOutput int
	}{
		// {
		// 	name:           "Part 1",
		// 	input:          testInput,
		// 	expectedOutput: 4361,
		// },
		{
			name:           "Part 2",
			input:          testInput,
			expectedOutput: 467835,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualOutput := solution(test.input)
			if actualOutput != test.expectedOutput {
				t.Errorf("Failed: expected %d, got %d", test.expectedOutput, actualOutput)
			}
		})
	}
}
