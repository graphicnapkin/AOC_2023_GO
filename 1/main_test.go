package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput int
	}{
		{
			name:           "Part 1",
			input:          "1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet",
			expectedOutput: 142,
		},
		{
			name:           "Part 2",
			input:          "two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen",
			expectedOutput: 281,
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
