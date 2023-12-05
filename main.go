package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Seed struct {
	Seed       uint64
	Soil       uint64
	Fertilizer uint64
	Water      uint64
	Light      uint64
	Temp       uint64
	Humidity   uint64
	Location   uint64
}

type ConvertMap []struct {
	Source      uint64
	Destination uint64
	Length      uint64
}

func main() {
	var allRecipies []Seed
	var seedsList, seed2soil, soil2Fertilizer, fertilizer2water, water2light, light2temp, temp2hum, hum2loc = GetInput()
	for i, seed := range seedsList {
		allRecipies = append(allRecipies, struct {
			Seed       uint64
			Soil       uint64
			Fertilizer uint64
			Water      uint64
			Light      uint64
			Temp       uint64
			Humidity   uint64
			Location   uint64
		}{Seed: uintParse(seed)})
		allRecipies[i].Soil = Convert(seed2soil, allRecipies[i].Seed)
		allRecipies[i].Fertilizer = Convert(soil2Fertilizer, allRecipies[i].Soil)
		allRecipies[i].Water = Convert(fertilizer2water, allRecipies[i].Fertilizer)
		allRecipies[i].Light = Convert(water2light, allRecipies[i].Water)
		allRecipies[i].Temp = Convert(light2temp, allRecipies[i].Light)
		allRecipies[i].Humidity = Convert(temp2hum, allRecipies[i].Temp)
		allRecipies[i].Location = Convert(hum2loc, allRecipies[i].Humidity)
	}
	var location = allRecipies[0].Location
	for _, each := range allRecipies {
		if each.Location < location {
			location = each.Location
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
		if value >= row.Source && value <= row.Source+row.Length {
			index := value - row.Source
			result = row.Destination + index
			return result
		}
	}
	return value
}
