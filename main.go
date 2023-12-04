package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cards []struct {
	Number int
	Count  int
}

func main() {
	var gameNumbersSum int
	var rawStrings = GetInput()
	var allCards Cards
	for i := range rawStrings {
		allCards = append(allCards, struct {
			Number int
			Count  int
		}{Number: i + 1, Count: 1})
	}
	for i, string := range rawStrings {
		var matches []int
		rawNumber := strings.ReplaceAll(strings.Split(string, ":")[0], "Card ", "")
		allCards[i].Number, _ = strconv.Atoi(strings.ReplaceAll(rawNumber, " ", ""))
		rows := strings.Split(strings.Split(string, ":")[1], "|")
		for _, winNumber := range ParseNumber(rows[0]) {
			for _, yourNumber := range ParseNumber(rows[1]) {
				if winNumber == yourNumber {
					matches = append(matches, winNumber)
				}
			}
		}
		if len(matches) > 0 {
			for y := i + 1; y <= i+len(matches); y++ {
				allCards[y].Count += allCards[i].Count
			}
		}
	}
	for _, card := range allCards {
		gameNumbersSum += card.Count
	}
	fmt.Printf("Game numbers sum is %d", gameNumbersSum)
}

func ParseNumber(input string) (result []int) {
	for _, value := range strings.Split(input, " ") {
		value = strings.ReplaceAll(value, " ", "")
		number, err := strconv.Atoi(value)
		if err == nil {
			result = append(result, number)
		}
	}
	return result
}

func GetInput() (result []string) {
	input, _ := os.ReadFile("input")
	scanner := bufio.NewScanner(strings.NewReader(string(input)))
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result
}
