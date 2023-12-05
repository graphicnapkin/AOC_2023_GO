package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	input := string(data)
	fmt.Println("Solution: ", solution(input))
}

func solution(input string) int {
	sum := 0
	rawCards := strings.Split(input, "\n")

	for _, rawCard := range rawCards {
		cardValue := 0
		splitCard := strings.Split(strings.Split(rawCard, ": ")[1], " | ")

		winningNumbers := strings.Split(splitCard[0], " ")
		winningNumbersMap := map[string]bool{}
		for _, winningNumber := range winningNumbers {
			winningNumbersMap[winningNumber] = true
		}
		myNumbers := strings.Split(splitCard[1], " ")

		for _, winningNumber := range winningNumbers {
			if winningNumber == "" {
				continue
			}
			if strings.Contains(myNumbers+" ", winningNumber+" ") {
				if cardValue == 0 {
					cardValue = 1
					continue
				}

				cardValue *= 2
			}
		}
		if cardValue == 0 {
			continue
		}
		sum += cardValue
	}
	return sum
}
