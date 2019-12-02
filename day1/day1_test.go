package main

import "testing"

func TestCalculateFuelRequirements(t *testing.T) {
	t.Run("divisible mass", func(t *testing.T) {
		got := CalculateFuelRequirements(12)
		want := 2

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("floored mass", func(t *testing.T) {
		got := CalculateFuelRequirements(17)
		want := 3

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
