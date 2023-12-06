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
	rawInstructions := strings.Split(input, "\n")
	seeds := strings.Split(strings.Split(rawInstructions[0], "seeds: ")[1], " ")
	lowestSeed := 0

	i := 3
	i, seedToSoilMap := makeMapFunction(rawInstructions, i)
	i, soilToFertilierMap := makeMapFunction(rawInstructions, i)
	i, fertalizerToWaterMap := makeMapFunction(rawInstructions, i)
	i, waterToLightMap := makeMapFunction(rawInstructions, i)
	i, lightToTemperatureMap := makeMapFunction(rawInstructions, i)
	i, temperatureToHumidityMap := makeMapFunction(rawInstructions, i)
	_, humidityToLocationMap := makeMapFunction(rawInstructions, i)

	for i, seed := range seeds {
		seedInt, _ := strconv.Atoi(seed)
		soil := seedToSoilMap(seedInt)
		fertilizer := soilToFertilierMap(soil)
		water := fertalizerToWaterMap(fertilizer)
		light := waterToLightMap(water)
		temperature := lightToTemperatureMap(light)
		humidity := temperatureToHumidityMap(temperature)
		location := humidityToLocationMap(humidity)
		if i == 0 {
			lowestSeed = location
		} else {
			if location < lowestSeed {
				lowestSeed = location
			}
		}
	}

	return lowestSeed
}

func makeMapFunction(rawInstructions []string, start int) (int, func (int)int) {
	sourceDestinationMap := map[int]int{}
	end := 0
	for i := start; i < len(rawInstructions); i++ {
		if rawInstructions[i] != "" {
			row := strings.Split(rawInstructions[i], " ")
			numbers := []int{}
			for _, number := range row {
				num, _ := strconv.Atoi(number)
				numbers = append(numbers, num)
			}
			destinationRanngeStart := numbers[0]
			sourceRangeStart := numbers[1]
			length := numbers[2]

			for k := 0; k < length; k++ {
				sourceDestinationMap[sourceRangeStart+k] = destinationRanngeStart + k
			}
		}
		if rawInstructions[i] == "" {
			end = i + 2
			break
		}
	}

	var mapFunction = func (source int) int {
		if destination, ok := sourceDestinationMap[source]; ok {
			return destination
		}
		return source
	}

	return end, mapFunction
}

func solutionPart2(input string) int {
	return 0
}
