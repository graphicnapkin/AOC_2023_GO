package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	test()
	data, _ := os.ReadFile("input.txt")
	input := string(data)
	fmt.Println("Solution:", solution(input))
}

func solution(input string) int {
	inputRows := strings.Split(input, "\n")
	games := make([]Game, len(inputRows))
	sumOfGameIDs := 0
	power := 0
	for i, row := range inputRows {
		var game Game
		game.id = i + 1
		row = strings.Split(row, ":")[1]
		unprocessedGames := strings.Split(row, ";")
		for _, gameRow := range unprocessedGames {
			boxes := strings.Split(gameRow, ",")
			for _, box := range boxes {
				box = strings.TrimSpace(box)

				if strings.Contains(box, "blue") {
					blueBoxes, _ := strconv.Atoi(strings.Split(box, " ")[0])
					if game.blue < blueBoxes {
						game.blue = blueBoxes
					}
				}

				if strings.Contains(box, "red") {
					redBoxes, _ := strconv.Atoi(strings.Split(box, " ")[0])
					if game.red < redBoxes {
						game.red = redBoxes
					}
				}

				if strings.Contains(box, "green") {
					greenBoxes, _ := strconv.Atoi(strings.Split(box, " ")[0])
					if game.green < greenBoxes {
						game.green = greenBoxes
					}
				}
			}
		}
		games[i] = game
	}

	//solution to part 1
	for _, game := range games {
		if game.red <= 12 && game.green <= 13 && game.blue <= 14 {
			sumOfGameIDs += game.id
		}
	}

	//solution to part 2
	for _, game := range games {
		power += game.red * game.green * game.blue
	}

	return power

}

type Game struct {
	id    int
	blue  int
	red   int
	green int
}

func test() {
	testInput :=
		`Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

	expectedOutput := 281
	actualOutput := solution(testInput)

	if expectedOutput != actualOutput {
		fmt.Println("Failed expected", expectedOutput, "got", actualOutput)
	} else {
		fmt.Println("Passed", actualOutput)
	}
}
