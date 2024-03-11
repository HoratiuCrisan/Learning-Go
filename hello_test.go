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
}
