package main

import (
	"reflect"
	"testing"
)

func TestIntCoder(t *testing.T) {
	t.Run("addition", func(t *testing.T) {
		got := IntCoder([]int{1, 0, 0, 0, 99})
		want := []int{2, 0, 0, 0, 99}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("multiplication", func(t *testing.T) {
		got := IntCoder([]int{2, 3, 0, 3, 99})
		want := []int{2, 3, 0, 6, 99}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("halting", func(t *testing.T) {
		got := IntCoder([]int{2, 4, 4, 5, 99, 0})
		want := []int{2, 4, 4, 5, 99, 9801}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("halt replace", func(t *testing.T) {
		got := IntCoder([]int{1, 1, 1, 4, 99, 5, 6, 0, 99})
		want := []int{30, 1, 1, 4, 2, 5, 6, 0, 99}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
