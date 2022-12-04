package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func FindElfCarryingMostCalories(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))

	mostCalories := 0
	currentElf := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			n, _ := strconv.Atoi(line)
			currentElf += n
		} else {
			if currentElf > mostCalories {
				mostCalories = currentElf
			}
			currentElf = 0
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return mostCalories
}

func main() {
	input, err := os.ReadFile("./01_input.txt")

	if err != nil {
		log.Fatal(err)
	}

	calories := FindElfCarryingMostCalories(string(input))
	fmt.Println(calories)
}
