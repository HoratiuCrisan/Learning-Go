package main

/* Creating tests for slices and arrays
to understand the difference between them and how they work */

import (
	"reflect"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sum of 5 numbers from an array", func(t *testing.T) {
		//numbers := [5]int{1, 2, 3, 4, 5} Refactoring form using an array to using a slice
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers) /* Passing an aray of 5 integers to a function to get the sum of the numbers */
		expected := 15

		assertFuncion(t, got, expected)
	})

	/* Arrays need to have aa specidief length */

	/* Slices allow collections of any size */

	// Run go test -cover to check if the the test covers all cases

	t.Run("Returning a slice of with the total all slices passed as arguments", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		expected := []int{3, 9}

		/* to compare the slices we can use the !reflect.DeepEqual to check if each element is equal */
		/* However, reflect.DeepEqual is not type safe */
		/* You can also use slices.equal to compare the slices */
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %#v, got %#v", expected, got)
		}
	})

	t.Run("Returning the sum of the tails of each n slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}

		assertSlices(t, got, expected)
	})

	t.Run("Testing the tail for an empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 2})
		expected := []int{0, 2}

		assertSlices(t, got, expected)
	})

	t.Run("Testing the tail function for an empty slice and a slice with 1 element", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1})
		expected := []int{0, 0}

		assertSlices(t, got, expected)
	})

	t.Run("Testing for one slice only", func(t *testing.T) {
		got := SumAllTails([]int{})
		expected := []int{0}

		assertSlices(t, got, expected)
	})
}

func assertSlices(t *testing.T, got, expected []int) {
	t.Helper()

	if !slices.Equal(got, expected) {
		t.Errorf("got %v, expected %v", got, expected)
	}
}

func assertFuncion(t *testing.T, got, expected int) {
	t.Helper()

	if got != expected {
		t.Errorf("got %d, expected %d", got, expected)
	}
}
