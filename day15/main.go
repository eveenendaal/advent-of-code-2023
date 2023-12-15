package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func hash(input string) int {
	currentValue := 0

	// iterate over each character
	for _, char := range input {
		// convert to int
		charValue := int(char)
		// add to current value
		currentValue += charValue
		// multiply by 17
		currentValue *= 17
		// mod by 256
		currentValue %= 256
	}

	return currentValue
}

type Lens struct {
	label       string
	focalLength int
}

func parseStep(input string) (int, bool, Lens) {
	addLens := true
	label := ""
	if strings.Contains(input, "-") {
		label = strings.Split(input, "-")[0]
		addLens = false
	}

	// only box additions have a focal length
	focalLength := 0
	if addLens {
		parts := strings.Split(input, "=")
		if len(parts) == 2 {
			label = parts[0]
			focalLength, _ = strconv.Atoi(parts[1])
		}
	}

	lens := Lens{
		label:       label,
		focalLength: focalLength,
	}
	boxId := hash(lens.label)
	return boxId, addLens, lens
}

// remove removes the given element at index i from the slice
// remaining elements are shifted to the left
func remove[E any](slice []E, s int) []E {
	return append(slice[:s], slice[s+1:]...)
}

func Solve(filePath string, part2 bool) int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	total := 0
	boxes := make([][]Lens, 256)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		// split string on ,
		blocks := strings.Split(scanner.Text(), ",")
		for _, block := range blocks {
			if part2 {
				boxId, addLens, lens := parseStep(block)
				existing := slices.IndexFunc(boxes[boxId], func(l Lens) bool {
					return l.label == lens.label
				})

				if addLens {
					if existing > -1 {
						// replace exiting lens with same label
						boxes[boxId][existing] = lens
					} else {
						// add to end of box
						boxes[boxId] = append(boxes[boxId], lens)
					}
				} else {
					// remove from box if exists
					if existing > -1 {
						boxes[boxId] = remove(boxes[boxId], existing)
					}
				}

			} else {
				total += hash(block)
			}
		}
	}

	err = file.Close()
	if err != nil {
		log.Fatalf("failed to close file: %s", err)
	}

	if part2 {
		total = 0
		for i, box := range boxes {
			for j, lens := range box {
				total += (i + 1) * (j + 1) * lens.focalLength
			}
		}
		return total
	} else {
		return total
	}
}

func main() {
	// fmt.Printf("Part 1: %d\n", Solve("input.txt", false))
	fmt.Printf("Part 2: %d\n", Solve("input.txt", true))
}
