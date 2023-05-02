package main

import "fmt"

// assign a function to a name
func main() {
	sub := func(a, b int) int {
		return a - b
	}
	// use the name to call the function
	fmt.Println(sub(13, 6))
}
