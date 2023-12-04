package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var gameNumbersSum float64
	var rawStrings = GetInput()
	for _, string := range rawStrings {
		var matches []int
		rows := strings.Split(strings.Split(string, ":")[1], "|")
		for _, winNumber := range ParseNumber(rows[0]) {
			for _, yourNumber := range ParseNumber(rows[1]) {
				if winNumber == yourNumber {
					matches = append(matches, winNumber)
				}
			}
		}
		if len(matches) > 0 {
			gameNumbersSum += math.Pow(2, float64(len(matches)-1))
		}
	}
	fmt.Printf("Game numbers sum is %s", strconv.FormatFloat(gameNumbersSum, 'f', -1, 64))
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
