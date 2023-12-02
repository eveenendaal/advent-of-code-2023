package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read test1.txt file
	file, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	total := 0
	totalPower := 0

	maxRed := 12
	maxBlue := 14
	maxGreen := 13

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		game := processLine(line)

		possible := true
		for _, round := range game.rounds {
			if round.red > maxRed || round.blue > maxBlue || round.green > maxGreen {
				possible = false
			}
		}

		if possible {
			total += game.number
		}
		totalPower += game.power
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("Total: %d Power: %d\n", total, totalPower)
}

// Create Round struct
type Round struct {
	red   int
	blue  int
	green int
}

type Game struct {
	rounds   []Round
	number   int
	maxRed   int
	maxBlue  int
	maxGreen int
	power    int
}

func processLine(line string) Game {
	results := []Round{}

	fmt.Println(line)
	// split line on ":"
	parts := strings.Split(line, ":")
	// Just get the digits in the gameString
	gameNumberString := strings.TrimFunc(parts[0], func(r rune) bool {
		return r < '0' || r > '9'
	})
	// Convert ot int
	gameNumber, _ := strconv.Atoi(gameNumberString)
	fmt.Printf("Game Number: %d\n", gameNumber)

	maxRed := 0
	maxBlue := 0
	maxGreen := 0

	rounds := strings.Split(parts[1], ";")
	for _, round := range rounds {
		roundData := Round{}

		// fmt.Println(round)
		// split round on ","
		roundParts := strings.Split(round, ",")

		for _, next := range roundParts {
			next = strings.ToLower(next)
			// Just get the digits in the roundString
			countString := strings.TrimFunc(next, func(r rune) bool {
				return r < '0' || r > '9'
			})
			// Convert to int
			count, _ := strconv.Atoi(countString)

			// if string contains
			if strings.Contains(next, "red") {
				// create new Round struct
				roundData.red = count
				if count > maxRed {
					maxRed = count
				}
			}
			if strings.Contains(next, "blue") {
				// create new Round struct
				roundData.blue = count
				if count > maxBlue {
					maxBlue = count
				}
			}
			if strings.Contains(next, "green") {
				// create new Round struct
				roundData.green = count
				if count > maxGreen {
					maxGreen = count
				}
			}
		}

		results = append(results, roundData)
		fmt.Printf("Red: %d, Blue: %d, Green: %d\n", roundData.red, roundData.blue, roundData.green)
	}

	return Game{
		rounds:   results,
		number:   gameNumber,
		maxRed:   maxRed,
		maxBlue:  maxBlue,
		maxGreen: maxGreen,
		power:    maxRed * maxBlue * maxGreen,
	}
}
