package main

import (
	"testing"
)

func TestHello(t *testing.T) { /* Running a test to check if the returning value of the `Hello() is `Hello, world!`*/
	got := Hello("Chris")
	want := "Hello, Chris!"

	if got != want { /* If the result is different, we are printing the error message */
		t.Errorf("got %v, want %v", got, want)
	}

	/*
		Creating a failing function
		This is a subset of tests that allow tests to be groupped around "a thing"
		A subset describes different scenarios
	*/
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world!"

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
