package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	Kind  int
	Label string
	Bid   int
	Set   []struct {
		Card  string
		Count int
	}
	Card []int
}

func main() {
	var handsAll []Hand
	var totalSum int
	var multiplier int = 1
	for _, hand := range GetInput() {
		handsAll = append(handsAll, ParseHand(hand))
	}
	sort.SliceStable(handsAll, func(i, j int) bool {
		return handsAll[i].Kind < handsAll[j].Kind
	})
	for i := 1; i <= 7; i++ {
		var temp []Hand
		for _, hand := range handsAll {
			if hand.Kind == i {
				temp = append(temp, hand)
			}
		}
		sort.SliceStable(temp, func(i, j int) bool {
			if temp[i].Card[0] == temp[j].Card[0] {
				if temp[i].Card[1] == temp[j].Card[1] {
					if temp[i].Card[2] == temp[j].Card[2] {
						if temp[i].Card[3] == temp[j].Card[3] {
							if temp[i].Card[4] == temp[j].Card[4] {
								return temp[i].Card[5] < temp[j].Card[5]
							}
							return temp[i].Card[4] < temp[j].Card[4]
						}
						return temp[i].Card[3] < temp[j].Card[3]
					}
					return temp[i].Card[2] < temp[j].Card[2]
				}
				return temp[i].Card[1] < temp[j].Card[1]
			}
			return temp[i].Card[0] < temp[j].Card[0]

		})
		for _, each := range temp {
			totalSum += each.Bid * multiplier
			fmt.Println(each.Label, each.Bid)
			multiplier++
		}
	}

	fmt.Println(totalSum)
}

func GetInput() (result []string) {
	input, _ := os.ReadFile("input")
	scanner := bufio.NewScanner(strings.NewReader(string(input)))
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result
}

func ParseHand(input string) (hand Hand) {
	handString := strings.Split(input, " ")[0]
	hand.Bid, _ = strconv.Atoi(strings.Split(input, " ")[1])
	var Joker = strings.Count(handString, "J")
	if Joker == 0 {
		for _, card := range []string{"A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2"} {
			if strings.Count(handString, card) > 0 {
				hand.Set = append(hand.Set, struct {
					Card  string
					Count int
				}{Card: card, Count: strings.Count(handString, card)})
			}
		}
		if len(hand.Set) == 5 {
			hand.Kind = 1
		} else if len(hand.Set) == 4 {
			hand.Kind = 2
		} else if len(hand.Set) == 3 {
			if hand.Set[0].Count == 3 || hand.Set[1].Count == 3 || hand.Set[2].Count == 3 {
				hand.Kind = 4
			} else {
				hand.Kind = 3
			}
		} else if len(hand.Set) == 2 {
			if hand.Set[0].Count == 4 || hand.Set[1].Count == 4 {
				hand.Kind = 6
			} else {
				hand.Kind = 5
			}
		} else if len(hand.Set) == 1 {
			hand.Kind = 7
		}
	} else {
		for _, card := range []string{"A", "K", "Q", "T", "9", "8", "7", "6", "5", "4", "3", "2"} {
			if strings.Count(handString, card) > 0 {
				hand.Set = append(hand.Set, struct {
					Card  string
					Count int
				}{Card: card, Count: strings.Count(handString, card)})
			}
		}
		hand.Kind = JokerFix(hand, Joker)
	}

	for _, char := range handString {
		switch string(char) {
		case "A":
			hand.Card = append(hand.Card, 62)
		case "K":
			hand.Card = append(hand.Card, 61)
		case "Q":
			hand.Card = append(hand.Card, 60)
		case "J":
			hand.Card = append(hand.Card, 10)
		case "T":
			hand.Card = append(hand.Card, 58)
		default:
			hand.Card = append(hand.Card, int(char))
		}
	}
	hand.Label = handString
	return hand
}

func JokerFix(hand Hand, joker int) (kind int) {
	if joker == 5 {
		kind = 7
	} else if joker == 4 {
		kind = 7
	} else if joker == 3 {
		if len(hand.Set) == 2 {
			kind = 6
		} else if len(hand.Set) == 1 {
			kind = 7
		}
	} else if joker == 2 {
		if len(hand.Set) == 1 {
			kind = 7
		} else if len(hand.Set) == 2 {
			kind = 6
		} else if len(hand.Set) == 3 {
			kind = 4
		}
	} else if joker == 1 {
		if len(hand.Set) == 1 {
			kind = 7
		} else if len(hand.Set) == 2 {
			if hand.Set[0].Count == 3 || hand.Set[1].Count == 3 {
				kind = 6
			} else {
				kind = 5
			}
		} else if len(hand.Set) == 3 {
			kind = 4
		} else if len(hand.Set) == 4 {
			kind = 2
		}
	}
	return kind
}
