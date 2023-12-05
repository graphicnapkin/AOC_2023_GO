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

type Gear struct {
	connectedParts []PartNumber
}

func solution(input string) int {
	inputRows := strings.Split(input, "\n")
	gears := map[string]Gear{}

	sum := 0

	for y, row := range inputRows {
		for x := 0; x < len(row); x++ {
			char := row[x]
			if char >= '0' && char <= '9' {
				numStart := []int{x, y}
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
				numStop := []int{nextIndex - 1, y}

				potentialPart := PartNumber{
					value: value,
					start: numStart,
					stop:  numStop,
				}

				valid, gearLocation := checkForValidNumber(potentialPart, inputRows)
				if valid {
					if len(gearLocation) == 0 {
						continue
					}
					coordinates := strconv.Itoa(gearLocation[0]) + "," + strconv.Itoa(gearLocation[1])

					if gear, ok := gears[coordinates]; ok {
						gear.connectedParts = append(gear.connectedParts, potentialPart)
						gears[coordinates] = gear
					} else {
						gear := Gear{
							connectedParts: []PartNumber{potentialPart},
						}
						gears[coordinates] = gear
					}
				}
			}
		}

	}

	for _, gear := range gears {
		gearRatio := 1
		if len(gear.connectedParts) < 2 {
			continue
		}
		for _, part := range gear.connectedParts {
			gearRatio *= part.value
		}
		sum += gearRatio
	}

	return sum
}

func checkPosition(x int, y int, grid []string) (bool, []int) {
	if x < 0 || y < 0 || x >= len(grid[0]) || y >= len(grid) {
		return false, []int{}
	}

	char := grid[y][x]
	if char == '*' {
		return true, []int{x, y}
	}

	return !(char >= '0' && char <= '9') && char != '.', []int{}
}

func checkForValidNumber(partNumber PartNumber, grid []string) (bool, []int) {
	startX := partNumber.start[0]
	startY := partNumber.start[1]
	stopX := partNumber.stop[0]
	stopY := partNumber.stop[1]

	// check adjacent positions to the left of the starting point
	left, gearLocation := checkPosition(startX-1, startY, grid)
	if left {
		return true, gearLocation
	}

	upperLeft, geargearLocation := checkPosition(startX-1, startY+1, grid)
	if upperLeft {
		return true, geargearLocation
	}

	lowerLeft, geargearLocation := checkPosition(startX-1, startY-1, grid)
	if lowerLeft {
		return true, geargearLocation
	}

	// check adjacent positions to the right of the ending point
	right, geargearLocation := checkPosition(stopX+1, stopY, grid)
	if right {
		return true, geargearLocation
	}

	upperRight, geargearLocation := checkPosition(stopX+1, stopY+1, grid)
	if upperRight {
		return true, geargearLocation
	}

	lowerRight, geargearLocation := checkPosition(stopX+1, stopY-1, grid)
	if lowerRight {
		return true, geargearLocation
	}

	// check positions direclty above and below all positions bewteen start and stop
	for x := startX; x <= stopX; x++ {
		up, geargearLocation := checkPosition(x, startY+1, grid)
		if up {
			return true, geargearLocation
		}

		down, geargearLocation := checkPosition(x, startY-1, grid)
		if down {
			return true, geargearLocation
		}
	}

	return false, []int{}
}

type PartNumber struct {
	value int
	start []int
	stop  []int
}
