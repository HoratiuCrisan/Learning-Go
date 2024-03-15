package main

import "fmt"

func Sum(numbers []int) (sum int) {
	/*for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}*/

	/*Refactoring with the range function */
	for _, value := range numbers { /* Since we do not need the index of the value, we use _ to ignore it */
		sum += value /* Adding each value to the sum variable */
	}
	return sum /* Returning the sum of the elements of the slice passed as argument */
}

func SumAll(slices ...[]int) (result []int) {
	// for _, i := range slices { /* Iterating over all the slices passed as arguments */
	//result = append(result, Sum(i)) /* Appending the result of the Sum function of each slice to the result variable */
	//}

	/* Refacoring the function to create a slice using make() */

	numbersOfSlices := len(slices)       /* Storing the number of slices passed as arguments */
	sums := make([]int, numbersOfSlices) /* Creating a new slice having the size equal to the number of slices passed as arguments */
	for i, numbers := range slices {     /* Iterating over the slices */
		sums[i] = Sum(numbers) /* Adding the sum of each slice to the new slice */
	}

	result = sums /* Copying the new slice into the result  slice */
	return result /* Returning the resulted slice */
}

func SumAllTails(slices ...[]int) (result []int) {
	for _, numbers := range slices {
		if len(numbers) > 1 {
			tails := numbers[1:]
			result = append(result, Sum(tails))
		} else {
			result = append(result, 0)
		}

	}
	return result /* Returning the result slice */
}

func main() {
	slice1 := []int{1, 12, 9, 8}
	slice2 := []int{0, 55, 1, 2, 8}

	fmt.Println(SumAllTails(slice1, slice2))
}
