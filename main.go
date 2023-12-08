package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Coordinates struct {
	Current string
	Left    string
	Right   string
}

func main() {
	var turnsList, allCoords = GetInput()
	var currentCoords Coordinates
	for _, coords := range allCoords {
		if coords.Current == "AAA" {
			currentCoords = coords
		}
	}
	var nextCoords Coordinates = currentCoords
	var moveCounter int
	for {
		for _, turn := range turnsList {
			currentCoords = nextCoords
			nextCoords = GetNextCoord(turn, currentCoords, allCoords)
			moveCounter++
			if nextCoords.Current == "ZZZ" {
				fmt.Println(moveCounter)
				os.Exit(0)
			}
		}
	}
}

func GetInput() (turns []string, coords []Coordinates) {
	input, _ := os.ReadFile("input")
	scanner := bufio.NewScanner(strings.NewReader(string(input)))
	var turnsList = ""
	for scanner.Scan() {
		var currentString = scanner.Text()
		if strings.Contains(currentString, "=") {
			currentString = strings.ReplaceAll(currentString, ",", "")
			splitted := strings.Split(strings.ReplaceAll(strings.ReplaceAll(currentString, "(", ""), ")", ""), " ")
			coords = append(coords, struct {
				Current string
				Left    string
				Right   string
			}{
				Current: splitted[0],
				Left:    splitted[2],
				Right:   splitted[3]})

		}
	}
	for _, turn := range turnsList {
		turns = append(turns, string(turn))
	}
	return turns, coords
}

func GetNextCoord(turn string, position Coordinates, allCoords []Coordinates) (nextCoord Coordinates) {
	var nextSpot string
	if turn == "L" {
		nextSpot = position.Left
	} else {
		nextSpot = position.Right
	}
	for _, coords := range allCoords {
		if nextSpot == coords.Current {
			nextCoord = coords
		}
	}
	return nextCoord
}
