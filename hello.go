package main

import (
	"fmt"
)

func Hello(name string) string { /* keeping our domain from the outside world (side-effects) */
	return "Hello, " + name + "!"
}

func main() { /* The fmt.Println() function is a side effect since it prints to the stdout */
	fmt.Println(Hello("world"))
}
