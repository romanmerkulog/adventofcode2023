package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	Number int
	Sets   []Set
}

type Set struct {
	R int
	G int
	B int
}

func main() {
	var gameNumbersSum int
	input, _ := os.ReadFile("input")
	scanner := bufio.NewScanner(strings.NewReader(string(input)))
	for scanner.Scan() {
		game := GameParse(scanner.Text())
		var totalRed int
		var totalGreen int
		var totalBlue int
		for _, gameSet := range game.Sets {
			totalRed = totalRed + gameSet.R
			totalGreen = totalGreen + gameSet.G
			totalBlue = totalBlue + gameSet.B
		}
		if totalRed <= 12 && totalGreen <= 13 && totalBlue <= 14 {
			gameNumbersSum = gameNumbersSum + game.Number
		}
	}
	fmt.Printf("Game numbers sum is %d", gameNumbersSum)
}

func GameParse(input string) (game Game) {
	cutOne := strings.Split(input, ":")
	game.Number, _ = strconv.Atoi(strings.Fields(cutOne[0])[1])
	setsBody := strings.Split(strings.ReplaceAll(cutOne[1], " ", ""), ";")
	for _, set := range setsBody {
		setParsed := SetParse(set)
		game.Sets = append(game.Sets, struct {
			R int
			G int
			B int
		}{
			R: setParsed.R,
			G: setParsed.G,
			B: setParsed.B,
		})
	}
	return game
}

func SetParse(input string) (set Set) {
	cutOne := strings.Split(input, ",")
	for _, cube := range cutOne {
		if strings.Contains(cube, "red") {
			set.R, _ = strconv.Atoi(strings.ReplaceAll(cube, "red", ""))
		} else if strings.Contains(cube, "green") {
			set.G, _ = strconv.Atoi(strings.ReplaceAll(cube, "green", ""))
		} else if strings.Contains(cube, "blue") {
			set.B, _ = strconv.Atoi(strings.ReplaceAll(cube, "blue", ""))
		}
	}
	return set
}
