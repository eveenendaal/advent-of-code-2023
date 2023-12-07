package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Create enum of cards
type Card string

// Create constants for cards
func (c Card) Score() int {
	// return the score of the card
	switch c {
	case "A":
		return 12
	case "K":
		return 11
	case "Q":
		return 10
	case "J":
		return 9
	case "T":
		return 8
	case "9":
	case "8":
	case "7":
	case "6":
	case "5":
	case "4":
	case "3":
	case "2":
		value, _ := strconv.Atoi(string(c))
		return value - 2
	}
	return -1
}

// Create a hand struct
type Hand struct {
	cards []Card
	bid   int
}

func (h Hand) Score() int {
	// return the score of the hand
	rankScore := 0
	handScore := 0

	// count all the cards
	cardCount := map[Card]int{}
	for _, card := range h.cards {
		cardCount[card]++
	}

	// check the max number of cards
	max := 0
	secondMax := 0
	for _, count := range cardCount {
		if count >= max {
			secondMax = max
			max = count
		}
	}

	if max == 5 {
		handScore = 6
	} else if max == 4 {
		handScore = 5
	} else if max == 3 && secondMax == 2 {
		handScore = 4
	} else if max == 3 {
		handScore = 3
	} else if max == 2 && secondMax == 2 {
		handScore = 2
	} else if max == 2 {
		handScore = 1
	}

	// check for four of a kind : are four of the five cards the same value
	rankScore += h.cards[0].Score() * int(math.Pow(12, 5))
	rankScore += h.cards[1].Score() * int(math.Pow(12, 4))
	rankScore += h.cards[2].Score() * int(math.Pow(12, 3))
	rankScore += h.cards[3].Score() * int(math.Pow(12, 2))
	rankScore += h.cards[4].Score() * int(math.Pow(12, 1))

	return (handScore * int(math.Pow(12, 6))) + rankScore
}

func Part1(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return -1
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	hands := []Hand{}
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		// split on space
		parts := strings.Split(line, " ")
		cardsString := parts[0]
		bidString := parts[1]
		// convert bid to int
		bid, _ := strconv.Atoi(bidString)
		// iterate through string characters
		cards := []Card{}
		for _, c := range cardsString {
			// convert character to card
			card := Card(string(c))
			// add card to hand
			cards = append(cards, card)
		}
		// create hand
		hand := Hand{
			cards: cards,
			bid:   bid,
		}
		// add hand to hands
		hands = append(hands, hand)
	}

	// sort hands by score
	sort.Slice(hands, func(i, j int) bool {
		return hands[i].Score() < hands[j].Score()
	})

	rank := 1
	output := 0
	for _, hand := range hands {
		fmt.Printf("Hand: %v, Score: %d, Rank: %d\n", hand, hand.Score(), rank)
		output += rank * hand.bid
		rank++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return output
}

func main() {
	result := Part1("data.txt")
	fmt.Println(result)
}
