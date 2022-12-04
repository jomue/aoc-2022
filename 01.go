package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func FindElvesCarryingMostCalories(input string, n int) int {
	scanner := bufio.NewScanner(strings.NewReader(input))

	var mostCalories []int
	currentElf := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			n, _ := strconv.Atoi(line)
			currentElf += n
		} else {
			mostCalories = append(mostCalories, currentElf)
			currentElf = 0
		}
	}
	mostCalories = append(mostCalories, currentElf)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(mostCalories)

	if n > len(mostCalories) {
		log.Fatal("n is larger than the number of elves ", len(mostCalories))
	}

	sum := 0
	for _, cal := range mostCalories[len(mostCalories)-n:] {
		sum += cal
	}

	return sum
}

func main() {
	input, err := os.ReadFile("./01_input.txt")

	if err != nil {
		log.Fatal(err)
	}

	calories := FindElvesCarryingMostCalories(string(input), 1)
	fmt.Println(calories)

	calories = FindElvesCarryingMostCalories(string(input), 3)
	fmt.Println(calories)
}
