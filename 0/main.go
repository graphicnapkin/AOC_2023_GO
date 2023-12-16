package main

import (
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	input := string(data)
	fmt.Println("Solution part 1: ", solutionPart1(input))
	fmt.Println("Solution: part 2: ", solutionPart1(input))
}

func solutionPart1(input string) int  {
	return 0
}

func solutionPart2(input string) int {
	return 0
}