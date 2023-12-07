package main

import (
	"bufio"
	"fmt"
	"os"
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
		return 13
	case "K":
		return 12
	case "Q":
		return 11
	case "J":
		return 10
	case "10":
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
	score := 0

	return score
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

	for _, hand := range hands {
		fmt.Printf("Hand: %v, Score: %d\n", hand, hand.Score())
	}

	output := 0

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return output
}

func main() {
	Part1("data.txt")
}
