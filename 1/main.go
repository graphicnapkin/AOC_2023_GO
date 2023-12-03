package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//testPart1()
	//testPart2()
	data, _ := os.ReadFile("input.txt")
	input := string(data)
	fmt.Println("solution", solution(input))
}

func solution(input string) int {
	sum := 0
	inputRows := strings.Split(input, "\n")

	for _, row := range inputRows {
		var rowNumbers []string
		var spelledOutNumbers = make([]string, len(row))

		for i, number := range []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"} {
			index := strings.Index(row, number)
			if index == -1 {
				continue
			}

			spelledOutNumbers[index] = string(i + '1')

			lastIndex := strings.LastIndex(row, number)
			if lastIndex == index {
				continue
			}

			spelledOutNumbers[lastIndex] = string(i + '1')
		}

		for i, char := range row {
			// if character is a number
			if char >= '0' && char <= '9' {
				rowNumbers = append(rowNumbers, string(char))
				// if we found a spelled out number at this index
			} else if spelledOutNumbers[i] != "" {
				rowNumbers = append(rowNumbers, spelledOutNumbers[i])
			}
		}

		if len(rowNumbers) == 0 {
			continue
		}

		if len(rowNumbers) == 1 {
			v, _ := strconv.Atoi(rowNumbers[0])
			sum += v * 2
			continue
		}
		num, _ := strconv.Atoi(rowNumbers[0] + rowNumbers[len(rowNumbers)-1])

		sum += num
		fmt.Println(rowNumbers, spelledOutNumbers, sum)
	}
	return sum
}

func testPart1() {
	testInput := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

	expectedOutput := 142
	actualOutput := solution(testInput)

	if expectedOutput != actualOutput {
		fmt.Println("Failed expected", expectedOutput, "got", actualOutput)
	} else {
		fmt.Println("Passed", actualOutput)
	}
}

func testPart2() {
	testInput := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`
	expectedOutput := 281
	actualOutput := solution(testInput)

	if expectedOutput != actualOutput {
		fmt.Println("Failed expected", expectedOutput, "got", actualOutput)
	} else {
		fmt.Println("Passed", actualOutput)
	}
}
