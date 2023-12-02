package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Set []struct {
	IndexFirst int
	IndexLast  int
	Num        int
}

func main() {
	input, _ := os.ReadFile("input")
	var sum int
	scanner := bufio.NewScanner(strings.NewReader(string(input)))
	for scanner.Scan() {
		inputString := scanner.Text()
		var Index Set
		for _, num := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9} {
			tempIndexFirst := strings.Index(inputString, strconv.Itoa(num))
			tempIndexLast := strings.LastIndex(inputString, strconv.Itoa(num))
			if tempIndexFirst > -1 || tempIndexLast > -1 {
				Index = append(Index, struct {
					IndexFirst int
					IndexLast  int
					Num        int
				}{
					IndexFirst: tempIndexFirst,
					IndexLast:  tempIndexLast,
					Num:        num,
				})
			}
		}
		for _, num := range []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"} {
			tempIndexFirst := strings.Index(inputString, num)
			tempIndexLast := strings.LastIndex(inputString, num)
			if tempIndexFirst > -1 || tempIndexLast > -1 {
				Index = append(Index, struct {
					IndexFirst int
					IndexLast  int
					Num        int
				}{
					IndexFirst: tempIndexFirst,
					IndexLast:  tempIndexLast,
					Num:        Replacer(num),
				})
			}
		}
		first, last := index2num(FindFirstIndex(Index), FindLastIndex(Index), Index)
		if first == last {
			log.Println("sdfsdf")
		}
		total, _ := strconv.Atoi(fmt.Sprintf("%s%s", first, last))

		sum = sum + total
	}
	log.Printf("total %d", sum)
}

func Replacer(input string) int {
	switch input {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	default:
		return 0
	}
}

func index2num(indexFirst int, indexLast int, List Set) (first, last string) {
	var firstItem int
	var lastItem int
	for _, item := range List {
		if indexFirst == item.IndexFirst {
			firstItem = item.Num
		}
		if indexLast == item.IndexLast {
			lastItem = item.Num
		}
	}

	return strconv.Itoa(firstItem), strconv.Itoa(lastItem)
}

func FindFirstIndex(Index Set) (first int) {
	first = 999
	for _, item := range Index {
		if item.IndexFirst < first {
			first = item.IndexFirst
		}
	}
	return first
}

func FindLastIndex(Index Set) (last int) {
	last = 0
	for _, item := range Index {
		if item.IndexLast > last {
			last = item.IndexLast
		}
	}
	return last
}
