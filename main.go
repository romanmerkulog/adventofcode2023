package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Array []string

type Coords struct {
	Row int
	Col int
}

func main() {
	var input = GetInput()
	var total int
	var exitHandle = true
	var prevCoords Coords = Coords{42, 8}
	var currentPos Coords = Coords{42, 9}
	var nextPos Coords = GetNextPoint(currentPos, prevCoords, input)
	total = 1
	for exitHandle {
		prevCoords = currentPos
		currentPos = nextPos
		nextPos = GetNextPoint(currentPos, prevCoords, input)
		if nextPos.Row == -1 {
			exitHandle = false
			total++
		} else {
			prevCoords = currentPos
			total++
		}
	}

	fmt.Println(total / 2)
}

func GetInput() (result []Array) {
	input, _ := os.ReadFile("input")
	scanner := bufio.NewScanner(strings.NewReader(string(input)))
	for scanner.Scan() {
		result = append(result, rowConvert(scanner.Text()))
	}
	return result
}

func rowConvert(input string) (result []string) {
	for _, val := range input {
		result = append(result, string(val))
	}
	return result
}

func NextCoord(input []int) (result []int) {
	for i := 0; i < len(input)-1; i++ {
		diff := input[i+1] - input[i]
		result = append(result, diff)
	}
	return result
}

func GetNextPoint(current Coords, prevCoords Coords, input []Array) (next Coords) {
	log.Println(input[current.Row][current.Col])
	switch input[current.Row][current.Col] {
	case "|":
		next.Col = current.Col
		if current.Row > prevCoords.Row {
			next.Row = current.Row + 1
		} else {
			next.Row = current.Row - 1
		}
	case "-":
		next.Row = current.Row
		if current.Col > prevCoords.Col {
			next.Col = current.Col + 1
		} else {
			next.Col = current.Col - 1
		}
	case "L":
		if current.Row == prevCoords.Row {
			next.Row = current.Row - 1
			next.Col = current.Col
		} else if current.Col == prevCoords.Col {
			next.Row = current.Row
			next.Col = current.Col + 1
		}
	case "J":
		if current.Col == prevCoords.Col {
			next.Row = current.Row
			next.Col = current.Col - 1
		} else if current.Row == prevCoords.Row {
			next.Row = current.Row - 1
			next.Col = current.Col
		}
	case "7":
		if current.Row == prevCoords.Row {
			next.Row = current.Row + 1
			next.Col = current.Col
		} else if current.Col == prevCoords.Col {
			next.Row = current.Row
			next.Col = current.Col - 1
		}
	case "F":
		if current.Row == prevCoords.Row {
			next.Row = current.Row + 1
			next.Col = current.Col
		} else if current.Col == prevCoords.Col {
			next.Row = current.Row
			next.Col = current.Col + 1
		}
	case ".":
		next.Row = 0
		next.Col = 0
	case "S":
		next = Coords{-1, -1}
	}
	return next
}
