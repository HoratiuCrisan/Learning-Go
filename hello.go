package main

import (
	"fmt"
)

const englishHelloPrefix = "Hello, "

func Hello(name string) string { /* keeping our domain from the outside world (side-effects) */
	if name == "" {
		name = "world"
	}
	return englishHelloPrefix + name + "!"
}

func main() { /* The fmt.Println() function is a side effect since it prints to the stdout */
	fmt.Println(Hello("world"))
}
