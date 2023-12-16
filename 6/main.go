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
	fmt.Println("Solution Part 1: ", solution(input))
	fmt.Println("Solution Part 2: ", solutionPart2(input))
	
}

func solution(input string) int {
	waysToWinOverall := []int{}

	rows := strings.Split(input, "\n")
	// get the first row and remove everything before a ":" and then split each number. There will be more than one space between each number
	timeRow := strings.Split(rows[0][strings.Index(rows[0], ":")+1:], " ")
	times := []int{}
	for _, time := range timeRow {
		if time != "" {
			num, _ := strconv.Atoi(time)
			times = append(times, num)
		}
	}

	// get the second row and remove everything before a ":" and then split each number. There will be more than one space between each number
	distance := strings.Split(rows[1][strings.Index(rows[1], ":")+1:], " ")
	distances := []int{}
	for _, dist := range distance {
		if dist != "" {
			num, _ := strconv.Atoi(dist)
			distances = append(distances, num)
		}
	}

	for i := 0; i < len(times); i++ {
		waysToWin := 0
		for j := 2; j < times[i]; j++ {
			if j * (times[i] - j) > distances[i] {
				waysToWin++
			}
		}
		waysToWinOverall = append(waysToWinOverall, waysToWin)
	}

	// multiply all the ways to win together and return the result
	result := 1
	for _, ways := range waysToWinOverall {
		result *= ways
	}
	return result
}

func solutionPart2(input string) int {

	rows := strings.Split(input, "\n")
	// get the first row and remove everything before a ":" and then split each number. There will be more than one space between each number
	timeString := ""
	timeRow := strings.Split(rows[0][strings.Index(rows[0], ":")+1:], " ")
	for _, time := range timeRow {
		if time != "" {
			timeString += time
		}
	}

	// get the second row and remove everything before a ":" and then split each number. There will be more than one space between each number
	distanceString := ""
	distanceRow := strings.Split(rows[1][strings.Index(rows[1], ":")+1:], " ")
	for _, dist := range distanceRow {
		if dist != "" {
			distanceString += dist
		}
	}

	distance, _ := strconv.Atoi(distanceString)
	time, _ := strconv.Atoi(timeString)
	waysToWin := 0

	for i := 2; i < time; i++ {
		if i * (time - i) > distance{
			waysToWin++
		}
	}

	// multiply all the ways to win together and return the result
	return waysToWin
}

