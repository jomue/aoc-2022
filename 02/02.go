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

var DEFEATS = map[Shape]Shape{
	ROCK:     SCISSORS,
	PAPER:    ROCK,
	SCISSORS: PAPER,
}

var LOSES = map[Shape]Shape{
	ROCK:     PAPER,
	PAPER:    SCISSORS,
	SCISSORS: ROCK,
}

var SYMBOLS = map[string]Shape{
	"A": ROCK,
	"B": PAPER,
	"C": SCISSORS,
}

var RESULT_SYMBOLS = map[string]int{
	"X": LOSS,
	"Y": DRAW,
	"Z": WIN,
}

const (
	LOSS = 0
	DRAW = 3
	WIN  = 6
)

func (s Shape) against(s2 Shape) int {
	result := DRAW

	if LOSES[s] == s2 {
		result = LOSS
	} else if DEFEATS[s] == s2 {
		result = WIN
	}

	return result
}

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

	sum := 0
	for _, n := range scores {
		sum += n
	}

	return sum
}

func parseInputLine(input string) int {
	symbols := strings.Split(input, " ")

	opponent := getSymbol(symbols[0])
	result := RESULT_SYMBOLS[symbols[1]]

	me := findShape(opponent, result)

	return calculateScore(me, opponent)
}

func getSymbol(input string) Shape {
	return SYMBOLS[input]
}

func findShape(opponent Shape, result int) Shape {
	var myShape Shape

	if result == DRAW {
		myShape = opponent
	}

	if result == LOSS {
		myShape = DEFEATS[opponent]
	}

	if result == WIN {
		myShape = LOSES[opponent]
	}

	return myShape
}

func calculateScore(p1 Shape, p2 Shape) int {
	score := p1.Score

	score += p1.against(p2)

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
