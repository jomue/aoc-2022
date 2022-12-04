package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Shape struct {
	Name  string
	Score int
}

var ROCK = Shape{"Rock", 1}
var PAPER = Shape{"Paper", 2}
var SCISSORS = Shape{"Scissors", 3}

var SYMBOLS = map[string]Shape{
	"A": ROCK,
	"B": PAPER,
	"C": SCISSORS,
	"X": ROCK,
	"Y": PAPER,
	"Z": SCISSORS,
}

const (
	LOSS = 0
	DRAW = 3
	WIN  = 6
)

func ScoreRockPaperScisscors(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))

	var scores []int

	for scanner.Scan() {
		line := scanner.Text()
		scores = append(scores, parseInputLine(line))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(scores)

	sum := 0
	for _, n := range scores {
		sum += n
	}

	return sum
}

func parseInputLine(input string) int {
	symbols := strings.Split(input, " ")

	opponent := getSymbol(symbols[0])
	me := getSymbol(symbols[1])

	return calculateScore(me, opponent)
}

func getSymbol(input string) Shape {
	return SYMBOLS[input]
}

func calculateScore(p1 Shape, p2 Shape) int {
	score := p1.Score

	if p1 == p2 {
		score += DRAW
	}

	if p1 == ROCK {
		if p2 == PAPER {
			score += LOSS
		}
		if p2 == SCISSORS {
			score += WIN
		}
	}

	if p1 == PAPER {
		if p2 == ROCK {
			score += WIN
		}
		if p2 == SCISSORS {
			score += LOSS
		}
	}

	if p1 == SCISSORS {
		if p2 == PAPER {
			score += WIN
		}
		if p2 == ROCK {
			score += LOSS
		}
	}

	fmt.Println(p1, p2, score)

	return score
}

func main() {
	input, err := os.ReadFile("./02_input.txt")

	if err != nil {
		log.Fatal(err)
	}

	score := ScoreRockPaperScisscors(string(input))
	fmt.Println(score)
}
