package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ConvertMap []struct {
	Source      uint64
	Destination uint64
	Length      uint64
}

func main() {
	var location uint64
	var seedsList, seed2soil, soil2Fertilizer, fertilizer2water, water2light, light2temp, temp2hum, hum2loc = GetInput()
	location = Convert(hum2loc, Convert(temp2hum, Convert(light2temp, Convert(water2light, Convert(fertilizer2water, Convert(soil2Fertilizer, Convert(seed2soil, uintParse(seedsList[0]))))))))
	for i := 0; i < len(seedsList); i = i + 2 {
		startItem := uintParse(seedsList[i])
		length := uintParse(seedsList[i+1])
		for y := startItem; y <= startItem+length; y++ {
			currentLoc := Convert(hum2loc, Convert(temp2hum, Convert(light2temp, Convert(water2light, Convert(fertilizer2water, Convert(soil2Fertilizer, Convert(seed2soil, y)))))))
			if currentLoc < location {
				location = currentLoc
			}
		}
	}
	fmt.Println(location)
}

func GetInput() (seedsList []string, seed2soil, soil2Fertilizer, fertilizer2water, water2light, light2temp, temp2hum, hum2loc ConvertMap) {
	input, _ := os.ReadFile("input")
	splittedInput := strings.Split(string(input), "\n\n")
	seedsList = strings.Split(strings.ReplaceAll(splittedInput[0], "seeds: ", ""), " ")
	seed2soil = MakeMap(strings.ReplaceAll(splittedInput[1], "seed-to-soil map:\n", ""))
	soil2Fertilizer = MakeMap(strings.ReplaceAll(splittedInput[2], "soil-to-fertilizer map:\n", ""))
	fertilizer2water = MakeMap(strings.ReplaceAll(splittedInput[3], "fertilizer-to-water map:\n", ""))
	water2light = MakeMap(strings.ReplaceAll(splittedInput[4], "water-to-light map:\n", ""))
	light2temp = MakeMap(strings.ReplaceAll(splittedInput[5], "light-to-temperature map:\n", ""))
	temp2hum = MakeMap(strings.ReplaceAll(splittedInput[6], "temperature-to-humidity map:\n", ""))
	hum2loc = MakeMap(strings.ReplaceAll(splittedInput[7], "humidity-to-location map:\n", ""))
	return seedsList, seed2soil, soil2Fertilizer, fertilizer2water, water2light, light2temp, temp2hum, hum2loc
}

func uintParse(input string) uint64 {
	var result, _ = strconv.ParseUint(input, 10, 64)
	return result
}

func MakeMap(input string) (result ConvertMap) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), " ")
		result = append(result, struct {
			Source      uint64
			Destination uint64
			Length      uint64
		}{
			Source:      uintParse(values[1]),
			Destination: uintParse(values[0]),
			Length:      uintParse(values[2]),
		})
	}
	return result
}

func Convert(input ConvertMap, value uint64) (result uint64) {
	for _, row := range input {
		if value >= row.Source && value < row.Source+row.Length {
			index := value - row.Source
			result = row.Destination + index
			return result
		}
	}
	return value
}
