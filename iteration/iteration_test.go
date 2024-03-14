package iteration

import (
	"testing"
)

/* A test function to send a character and get 5 of the same characters in return */
func TestRepeat(t *testing.T) {
	t.Run("5 iterations", func(t *testing.T) {
		repeated := Repeat('a', 5)
		expected := "aaaaa"

		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("10 iterations", func(t *testing.T) {
		repeated := Repeat('a', 10)
		expected := "aaaaaaaaaa"

		assertCorrectMessage(t, repeated, expected)
	})
}

func assertCorrectMessage(t *testing.T, repeated, expected string) {
	t.Helper()
	if repeated != expected {
		t.Errorf("expected %v got %v", expected, repeated)
	}
}

/* Creating a benchmark for the Repeat function */
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat('a', 10)
	}
}
