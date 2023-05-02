package main

import "fmt"

func main() {
	var A [10]int //declares an array of 10 "zeroed integers"
	A[5] = 20     // set elements of the array

	i := A[5] // get/read elements from the array

	var B = [3]int{100, 200, 300}             // declare and initialize an array
	C := [5]int{1000, 2000, 3000, 4000, 5000} // shorthand for declaring and initializing an array
	D := [...]int{5, 5, 5, 5, 5}              // ellipses for a variable number of elements

	fmt.Println(A, i, B, C, D)
}

// Output - [0 0 0 0 0 20 0 0 0 0] 20 [100 200 300] [1000 2000 3000 4000 5000] [5 5 5 5 5]
