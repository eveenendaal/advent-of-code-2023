package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Read test1.txt file
	file, err := os.Open("test1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
