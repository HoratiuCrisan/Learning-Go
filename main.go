package main

import (
	"fmt"
)

func main() {
	slice1 := []int{1, 3, 9}
	slice2 := []int{0, 12, 4, 6}

	fmt.Println(SumAllTails(slice1, slice2))

}
