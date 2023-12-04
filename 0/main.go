package main

import (
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	input := string(data)
	fmt.Println("Solution: ", solution(input))
}

func solution(input string) int {
	return 0
}
