package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Array []int

func main() {
	var input = GetInput()
	var flag bool
	var totalSum int
	for _, val := range input {
		var temp []Array
		temp = append(temp, val)
		nextDiff := val
		flag = true
		for flag {
			temp = append(temp, MakeDiff(nextDiff))
			if isAllZero(MakeDiff(nextDiff)) {
				flag = false
			} else {
				flag = true
				nextDiff = MakeDiff(nextDiff)
			}
		}
		for i := len(temp) - 1; i > 0; i-- {
			firstElem := temp[i-1][0] - temp[i][0]
			temp[i-1] = append([]int{firstElem}, temp[i-1]...)
		}
		totalSum += temp[0][0]
	}
	fmt.Println(totalSum)
}

func GetInput() (result []Array) {
	input, _ := os.ReadFile("input")
	scanner := bufio.NewScanner(strings.NewReader(string(input)))
	for scanner.Scan() {
		result = append(result, rowConvert(scanner.Text()))
	}
	return result
}

func rowConvert(input string) (result []int) {
	var splitted = strings.Split(input, " ")
	for _, val := range splitted {
		num, _ := strconv.Atoi(val)
		result = append(result, num)
	}
	return result
}

func isAllZero(input []int) (result bool) {
	result = true
	for _, val := range input {
		if val != 0 {
			result = false
		}
	}
	return result
}

func MakeDiff(input []int) (result []int) {
	for i := 0; i < len(input)-1; i++ {
		diff := input[i+1] - input[i]
		result = append(result, diff)
	}
	return result
}
