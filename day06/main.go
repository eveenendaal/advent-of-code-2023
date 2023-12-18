package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Race struct {
	time     int
	distance int
}

func Part1(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return -1
	}
	defer file.Close()

	races := make([]Race, 0)
	// create regex for spaces
	spacesRegex := regexp.MustCompile(`\s+`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())

		// If lines starts with "time:"
		if line[0:5] == "time:" {
			// Split line by spaces
			splitLine := spacesRegex.Split(line, -1)
			// storage the times to the Race struct
			times := splitLine[1:]
			for _, time := range times {
				time, _ := strconv.Atoi(time)
				races = append(races, Race{time: time})
			}
		}
		// If lines starts with "distance:"
		if line[0:9] == "distance:" {
			// split line by spaces
			splitLine := spacesRegex.Split(line, -1)
			// add the distances to the Race struct
			distances := splitLine[1:]
			for i, distance := range distances {
				distance, _ := strconv.Atoi(distance)
				races[i].distance = distance
			}
		}
	}

	output := 1

	// iterate over races
	for _, race := range races {
		fmt.Println(race)
		winningSpeeds := 0
		// calculate the speed
		for i := 0; i <= race.time; i++ {
			speed := i
			time := race.time - speed
			// calculate the distance
			distance := speed * time
			if distance > race.distance {
				winningSpeeds++
			}
		}
		output *= winningSpeeds
	}

	//

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return output
}

func Part2(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return -1
	}
	defer file.Close()

	race := Race{}
	// create regex for spaces
	digitRe := regexp.MustCompile(`[^\d]+`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())

		// If lines starts with "time:"
		if line[0:5] == "time:" {
			// just get the digits
			digits := string(digitRe.ReplaceAll([]byte(line), []byte("")))
			time, _ := strconv.Atoi(digits)
			race.time = time
		}
		// If lines starts with "distance:"
		if line[0:9] == "distance:" {
			// just get the digits
			digits := string(digitRe.ReplaceAll([]byte(line), []byte("")))
			distance, _ := strconv.Atoi(digits)
			race.distance = distance
		}
	}
	fmt.Println(race)
	winningSpeeds := 0

	// calculate the speed
	for i := 0; i <= race.time; i++ {
		speed := i
		time := race.time - speed
		// calculate the distance
		distance := speed * time
		if distance > race.distance {
			winningSpeeds++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return winningSpeeds
}

func main() {
	// result := Part1("data.txt")
	result := Part2("data.txt")
	fmt.Println(result)
}
