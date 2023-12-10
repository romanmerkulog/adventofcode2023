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

type Result []int

func main() {
	var input = GetInput()
	var total int
	var exitHandle = true
	var prevCoords Coords = Coords{0, 4}
	var currentPos Coords = Coords{1, 4}
	var nextPos = GetNextPoint(currentPos, prevCoords, input)
	var resultsMap = [10][20]int{}
	resultsMap[currentPos.Row][currentPos.Col] = 1
	for exitHandle {
		prevCoords = currentPos
		currentPos = nextPos
		resultsMap[currentPos.Row][currentPos.Col] = 1
		nextPos = GetNextPoint(currentPos, prevCoords, input)
		if nextPos.Row == -1 {
			exitHandle = false
		} else if nextPos.Row == -2 {
			log.Fatalln("error!")
		}
	}

	for i := 0; i < 9; i++ {
		for y := 0; y < 11; y++ {
			if resultsMap[i][y] == 0 {
				resultsMap[i][y] = 2
			} else {
				break
			}
		}
	}

	for i := 0; i < len(resultsMap); i++ {
		for y := 0; y < len(resultsMap); y++ {
			if resultsMap[i][y] == 0 {
				if Laser(i, y, resultsMap) {
					total++
				}
			}
		}
	}
	fmt.Println(total)
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

func Laser(x int, y int, matrix [10][20]int) (result bool) {
	var counter int
	for i := x; i > 0; i-- {
		if matrix[i][y] == 1 {
			counter++
		}
	}

	if counter != 0 && counter%2 > 0 {
		result = true
	} else {
		result = false
	}
	return result
}

func GetNextPoint(current Coords, prevCoords Coords, input []Array) (next Coords) {
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
		next = Coords{-2, -2}
	case "S":
		next = Coords{-1, -1}
	}
	return next
}
