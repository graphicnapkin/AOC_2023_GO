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

		winningNumbersMap := map[string]bool{}
		for i:=0; i < len(splitCard[0]); i+= 3 {
			winningNumber := splitCard[0][i:i+2]
			winningNumbersMap[winningNumber] = true
		}

		for i:=0; i < len(splitCard[1]); i+= 3 {
			myNumber := splitCard[1][i:i+2]
			if _, ok := winningNumbersMap[myNumber]; ok {
				if cardValue == 0 {
					cardValue = 1
					continue
				}
				cardValue *= 2
			}
		}

		if cardValue > 0 {
			sum += cardValue
		}
	}
	return sum
}


func solutionPart2(input string) int {
	sum := 0
	cardCount := map[int]int{}

	rawCards := strings.Split(input, "\n")
	for cardNumber, rawCard := range rawCards {
		cardCount[cardNumber] = cardCount[cardNumber] + 1

		splitCard := strings.Split(strings.Split(rawCard, ": ")[1], " | ")

		winningNumbersMap := map[string]bool{}
		// both winning numbers and your numbers are alwyas 2 digits with a space after
		// so we can just grab the first 2 characters of each 3 character chunk
		for i:=0; i < len(splitCard[0]); i+= 3 {
			winningNumber := splitCard[0][i:i+2]
			winningNumbersMap[winningNumber] = true
		}

		matches := 0
		for i:=0; i < len(splitCard[1]); i+= 3 {
			myNumber := splitCard[1][i:i+2]
			if _, ok := winningNumbersMap[myNumber]; ok {
				matches++
			}
		}

		for j := 0; j < cardCount[cardNumber]; j++ {
			for i:= cardNumber + 1; i < cardNumber + matches + 1; i++ {
				cardCount[i] = cardCount[i] + 1
			}
		}
	}

	for _, count := range cardCount {
		sum += count
	}
	return sum
}
