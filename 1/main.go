package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	input := string(data)
	fmt.Println("Solution: ", solution(input))
}

func solution(input string) int {
	sum := 0
	inputRows := strings.Split(input, "\n")

	for _, row := range inputRows {
		var rowNumbers string
		spelledOutNumbers := extractSpelledOutNumbers(row)

		for i, char := range row {
			// if character is a number
			if char >= '0' && char <= '9' {
				rowNumbers = string(char)
				break
			}
			// if we found a spelled out number at this index
			if spelledOutNumbers[i] != "" {
				rowNumbers = spelledOutNumbers[i]
				break
			}
		}

		for i := len(row) - 1; i >= 0; i-- {
			// if character is a number
			if row[i] >= '0' && row[i] <= '9' {
				rowNumbers += string(row[i])
				break
			}

			// if we found a spelled out number at this index
			if spelledOutNumbers[i] != "" {
				rowNumbers += string(spelledOutNumbers[i])
				break
			}
		}

		num, _ := strconv.Atoi(rowNumbers)

		sum += num
	}
	return sum
}

func extractSpelledOutNumbers(row string) []string {
	var spelledOutNumbers = make([]string, len(row))

	for i, number := range []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"} {
		index := strings.Index(row, number)
		if index == -1 {
			continue
		}

		spelledOutNumbers[index] = strconv.Itoa(i + 1)

		lastIndex := strings.LastIndex(row, number)
		if lastIndex == index {
			continue
		}

		spelledOutNumbers[lastIndex] = strconv.Itoa(i + 1)
	}
	return spelledOutNumbers
}
