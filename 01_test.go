package main

import "testing"

func TestFindElfCarryingMostCalories(t *testing.T) {

	t.Run("example input", func(t *testing.T) {
		input := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

		want := 24000
		got := FindElfCarryingMostCalories(input)

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
