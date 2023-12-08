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
	var flag bool = true
	var coordsList []Coordinates
	var results []int
	for _, coords := range allCoords {
		if strings.LastIndex(coords.Current, "A") == 2 {
			coordsList = append(coordsList, coords)
		}
	}
	var nextCoords Coordinates = currentCoords
	var moveCounter int
	for _, coord := range coordsList {
		flag = true
		moveCounter = 0
		currentCoords = coord
		nextCoords = coord
		for flag {
			for _, turn := range turnsList {
				currentCoords = nextCoords
				nextCoords = GetNextCoord(turn, currentCoords, allCoords)
				if strings.LastIndex(currentCoords.Current, "Z") == 2 {
					results = append(results, moveCounter)
					flag = false
				}
				moveCounter++
			}
		}
	}
	fmt.Println(LCM(results[0], results[1], results[2], results[3], results[4], results[5]))
}

func GetInput() (turns []string, coords []Coordinates) {
	input, _ := os.ReadFile("input")
	scanner := bufio.NewScanner(strings.NewReader(string(input)))
	var turnsList = "LRLRLRLR"
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

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
