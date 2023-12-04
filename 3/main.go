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
	inputRows := strings.Split(input, "\n")
	parts := make([]PartNumber, 0)

	sum := 0

	for y, row := range inputRows{
		for x := 0; x < len(row); x++ {
			char := row[x]
			if char >= '0' && char <= '9' {
				numStart := []int{x,y}
				var currentNumber = string(char)
				nextIndex := x + 1
				stop := nextIndex >= len(row)

				for !stop {
					if nextIndex >= len(row) {
						stop = true
						break
					}

					if row[nextIndex] >= '0' && row[nextIndex] <= '9' {
						currentNumber += string(row[nextIndex])
						nextIndex++
					} else {
						x = nextIndex
						stop = true
					}
				}

				value, _ := strconv.Atoi(currentNumber)
				numStop := []int{nextIndex -1, y}

				potentialPart := PartNumber{
					value: value,
					start: numStart,
					stop: numStop,
				}

				valid := checkForValidNumber(potentialPart, inputRows)
				if valid {
					parts = append(parts, potentialPart)
					sum += value
				}

				if !valid {
					fmt.Println("Invalid part: ", potentialPart)
				}
			}
		}

	}

	return sum
}

func checkPosition(x int, y int, grid []string) bool {
	if x < 0 { return false }
	if y < 0 { return false }
	if x >= len(grid[0]) { return false }
	if y >= len(grid) { return false }

	char := grid[y][x]
	return !(char >= '0' && char <= '9') && char != '.'
}

func checkForValidNumber(partNumber PartNumber, grid []string) bool {
	startX := partNumber.start[0]
	startY := partNumber.start[1]
	stopX := partNumber.stop[0]
	stopY := partNumber.stop[1]

	// check adjacent positions to the left of the starting point
	if checkPosition(startX - 1, startY, grid) { return true }
	if checkPosition(startX -	1, startY + 1, grid) { return true }
	if checkPosition(startX - 1, startY - 1, grid) { return true }

	// check adjacent positions to the right of the ending point
	if checkPosition(stopX + 1, stopY, grid) { return true }
	if checkPosition(stopX + 1, stopY + 1, grid) { return true }
	if checkPosition(stopX + 1, stopY - 1, grid) { return true }

	// check positions direclty above and below all positions bewteen start and stop
	for x := startX; x <= stopX; x++ {
		if checkPosition(x, startY - 1, grid) { return true }
		if checkPosition(x, stopY + 1, grid) { return true }
	}

	return false
}

type PartNumber struct {
	value int
	start []int
	stop  []int
}

// Instructions: --- Day 3: Gear Ratios ---
// You and the Elf eventually reach a gondola lift station; he says the gondola lift will take you up to the water source, but this is as far as he can bring you. You go inside.

// It doesn't take long to find the gondolas, but there seems to be a problem: they're not moving.

// "Aaah!"

// You turn around to see a slightly-greasy Elf with a wrench and a look of surprise. "Sorry, I wasn't expecting anyone! The gondola lift isn't working right now; it'll still be a while before I can fix it." You offer to help.

// The engineer explains that an engine part seems to be missing from the engine, but nobody can figure out which one. If you can add up all the part numbers in the engine schematic, it should be easy to work out which part is missing.

// The engine schematic (your puzzle input) consists of a visual representation of the engine. There are lots of numbers and symbols you don't really understand, but apparently any number adjacent to a symbol, even diagonally, is a "part number" and should be included in your sum. (Periods (.) do not count as a symbol.)

// Here is an example engine schematic:

// 467..114..
// ...*......
// ..35..633.
// ......#...
// 617*......
// .....+.58.
// ..592.....
// ......755.
// ...$.*....
// .664.598..
// In this schematic, two numbers are not part numbers because they are not adjacent to a symbol: 114 (top right) and 58 (middle right). Every other number is adjacent to a symbol and so is a part number; their sum is 4361.

// Of course, the actual engine schematic is much larger. What is the sum of all of the part numbers in the engine schematic?
