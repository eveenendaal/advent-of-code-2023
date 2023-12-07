package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Create enum of cards
type Card string

// Create constants for cards
func (c Card) Score(jokers bool) int {
	// return the score of the card
	switch c {
	case "A":
		return 14
	case "K":
		return 13
	case "Q":
		return 12
	case "J":
		if jokers {
			return 1
		} else {
			return 11
		}
	case "T":
		return 10
	default:
		value, _ := strconv.Atoi(string(c))
		return value
	}
}

// Create a hand struct
type Hand struct {
	cards  []Card
	bid    int
	source string
}

type Score struct {
	values      []int
	max         int
	secondMax   int
	counts      map[Card]int
	bid         int
	cards       []Card
	cardsString string
}

func (h Hand) Score(jokers bool) Score {
	// return the score of the hand
	score := []int{}

	handScore := 0

	// count all the cards
	cardCount := make(map[Card]int, 13)
	for _, card := range h.cards {
		cardCount[card]++
	}
	// Get joker count
	jokerCount := cardCount[Card("J")]
	// Set jokers to 0
	if jokers {
		cardCount[Card("J")] = 0
	}

	// Get card count values
	counts := []int{}
	for _, count := range cardCount {
		counts = append(counts, count)
	}
	sort.Ints(counts)
	// Get Maxes
	max := counts[len(counts)-1]
	secondMax := max
	if len(counts) > 1 {
		secondMax = counts[len(counts)-2]
	}

	if jokers {
		if (max + jokerCount) == 5 { // 5 of a kind
			handScore = 6
		} else if (max + jokerCount) == 4 { // 4 of a kind
			handScore = 5
		} else if (max+jokerCount) == 3 && secondMax == 2 { // Full House
			handScore = 4
		} else if (max + jokerCount) == 3 { // 3 of a kind
			handScore = 3
		} else if max == 2 && (jokerCount+secondMax) == 2 { // 2 pair
			handScore = 2
		} else if (max + jokerCount) == 2 { // 1 pair
			handScore = 1
		} else { // High Card
			handScore = 0
		}
	} else {
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
	}

	score = append(score, handScore)
	score = append(score, h.cards[0].Score(jokers))
	score = append(score, h.cards[1].Score(jokers))
	score = append(score, h.cards[2].Score(jokers))
	score = append(score, h.cards[3].Score(jokers))
	score = append(score, h.cards[4].Score(jokers))

	return Score{
		bid:         h.bid,
		values:      score,
		max:         max,
		secondMax:   secondMax,
		counts:      cardCount,
		cards:       h.cards,
		cardsString: h.source,
	}
}

func Solve(filePath string, jokers bool) int {
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
			cards:  cards,
			bid:    bid,
			source: cardsString,
		}
		// add hand to hands
		hands = append(hands, hand)
	}

	scores := []Score{}
	for _, hand := range hands {
		scores = append(scores, hand.Score(jokers))
	}

	// sort hands by score
	sort.Slice(scores, func(i, j int) bool {
		for k := range scores[i].values {
			if scores[i].values[k] != scores[j].values[k] {
				return scores[i].values[k] < scores[j].values[k]
			}
		}
		return false
	})

	// write to output.txt for pat 1
	f, _ := os.Create("output1.txt")
	defer f.Close()

	output := 0
	for i, score := range scores {
		f.Write([]byte(fmt.Sprintf("%d: %d, %s\n", i+1, score.bid, score.cardsString)))
		fmt.Printf("%d: %d, %v, %v\n", i+1, score.bid, score, score.cardsString)
		output += (i + 1) * score.bid
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return output
}

func main() {
	// result := Solve("data.txt", false) // Part 1
	result := Solve("data.txt", true) // Part 2
	fmt.Println(result)
}
