package main

import (
	"fmt"
	"strings"
  "strconv"
  "io/ioutil"
)

func main() {
	testSolution()
  data, _ := ioutil.ReadFile("input.txt")
  input := string(data)
  fmt.Println(solution(input))
}

func solution(input string) int {

  sum := 0
	inputRows := strings.Split(input, "\n")

	for _, row := range inputRows {
		var rowNumbers string
		// for each character
		for _, char := range row {
			// if character is a number
			if char >= '0' && char <= '9' {
        if len(rowNumbers) <= 1 {
          rowNumbers += string(char)
        }
        if len(rowNumbers) > 1 {
          rowNumbers= string(rowNumbers[0]) + string(char)
        }
			}
		}

    if len(rowNumbers) == 1 {
      rowNumbers = string(rowNumbers) + string(rowNumbers)
    }

    i, _:= strconv.Atoi(rowNumbers)
    sum += i
	}
  return sum

}

func testSolution() {
	testInput := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

	expectedOutput := 142
	actualOutput := solution(testInput)

	if expectedOutput != actualOutput {
		fmt.Println("Failed expected", expectedOutput, "got", actualOutput)
	} else {
  fmt.Println("Passed", actualOutput)}
}
