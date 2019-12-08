package main

import (
	"reflect"
	"testing"
)

func TestIntCoder(t *testing.T) {
	t.Run("addition", func(t *testing.T) {
		got, _ := IntCoder([]int{1, 0, 0, 0, 99}, 0, 0)
		want := []int{2, 0, 0, 0, 99}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("multiplication", func(t *testing.T) {
		got, _ := IntCoder([]int{2, 3, 0, 3, 99}, 3, 0)
		want := []int{2, 3, 0, 6, 99}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("halting", func(t *testing.T) {
		got, _ := IntCoder([]int{2, 4, 4, 5, 99, 0}, 4, 4)
		want := []int{2, 4, 4, 5, 99, 9801}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("halt replace", func(t *testing.T) {
		got, _ := IntCoder([]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, 1, 1)
		want := []int{30, 1, 1, 4, 2, 5, 6, 0, 99}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("with bad opcode", func(t *testing.T) {
		_, got := IntCoder([]int{5, 0, 0, 0, 99}, 0, 0)

		if got == nil {
			t.Errorf("got %s expected error", got)
		}
	})
}

func TestFindNounAndVerb(t *testing.T) {
	t.Run("with simple desired output", func(t *testing.T) {
		got, _ := FindNounAndVerb([]int{1, 0, 0, 0, 99}, 100)
		want := 4

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
