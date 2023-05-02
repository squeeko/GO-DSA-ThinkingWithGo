package main

import "fmt"

func main() {
	// var Z []int               //declare a slice - smiliar to an array but length is not specified
	// var Y = []int{1, 2, 3, 4} // declare and initialize a slice (backed by the array given implicitly)
	X := []int{1, 3, 5, 7, 9}
	var max int = len(X)
	// chars := []string{0: "a", 1: "b", 2: "c", 3: "d", 4: "e", 5: "f"}
	var T = X[2:max] //slices of the array var T = X[2:max], X[0:], X[:max], X[max:] etc.... also the range is exclusive....

	fmt.Println(T, max)

}
