package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Numbers []struct {
	Number     int
	IndexStart int
	IndexEnd   int
}

func main() {
	var gameNumbersSum int
	var rawStrings = GetInput()
	for i := range rawStrings {
		var allSymbols []int
		if i > 0 {
			allSymbols = append(allSymbols, ParseSymbol(rawStrings[i-1])...)
		}
		allSymbols = append(allSymbols, ParseSymbol(rawStrings[i])...)
		if i < len(rawStrings)-1 {
			allSymbols = append(allSymbols, ParseSymbol(rawStrings[i+1])...)
		}

		for _, number := range ParseNumber(rawStrings[i]) {
			for _, symbol := range allSymbols {
				if symbol >= number.IndexStart-1 && symbol <= number.IndexEnd+1 {
					gameNumbersSum += number.Number
					break
				}
			}
		}
	}
	fmt.Printf("Game numbers sum is %d", gameNumbersSum)
}

func ParseNumber(input string) (result Numbers) {
	re := regexp.MustCompile(`\d{1,}`)
	numbers := re.FindAllString(input, -1)
	for _, num := range numbers {
		index := strings.Index(input, num)
		number, _ := strconv.Atoi(num)
		result = append(result, struct {
			Number     int
			IndexStart int
			IndexEnd   int
		}{Number: number, IndexStart: index, IndexEnd: index + len(num) - 1})
		input = strings.Replace(input, num, DummyReplacer(len(num)), 1) // remove added number from line to avoid index duplicate
	}
	return result
}

func ParseSymbol(input string) (result []int) {
	for pos, char := range input {
		isSymbol, _ := regexp.MatchString(`[^.\d\s]`, fmt.Sprintf("%c", char))
		if isSymbol {
			result = append(result, pos)
		}
	}
	return result
}

func DummyReplacer(len int) (result string) { // generate dummy string to replace tthe number
	for i := 0; i < len; i++ {
		result += "."
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
