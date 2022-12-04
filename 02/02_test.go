package main

import "testing"

func TestScoreRockPaperScisscors(t *testing.T) {
	t.Run("example input", func(t *testing.T) {
		input := `A Y
B X
C Z`

		want := 15
		got := ScoreRockPaperScisscors(input)

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
