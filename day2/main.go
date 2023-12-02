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
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("Total: %d\n", total)
}

// Create Round struct
type Round struct {
	red   int
	blue  int
	green int
}

type Game struct {
	rounds []Round
	number int
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
			}
			if strings.Contains(next, "blue") {
				// create new Round struct
				roundData.blue = count
			}
			if strings.Contains(next, "green") {
				// create new Round struct
				roundData.green = count
			}
		}

		results = append(results, roundData)
		fmt.Printf("Red: %d, Blue: %d, Green: %d\n", roundData.red, roundData.blue, roundData.green)
	}

	return Game{
		rounds: results,
		number: gameNumber,
	}
}
