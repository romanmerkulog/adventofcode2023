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
	input, _ := os.ReadFile("input")
	scanner := bufio.NewScanner(strings.NewReader(string(input)))
	for scanner.Scan() {
		game := GameParse(scanner.Text())
		fmt.Println(game)
	}

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
	fmt.Println(cutOne)
	return set
}
