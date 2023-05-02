/*
What is a variadic function?
Functions in general accept only a fixed number of arguments. A variadic function is a function that accepts a variable number of arguments. If the last parameter of a function definition is prefixed by ellipsis ..., then the function can accept any number of arguments for that parameter.
*/

package main

import "fmt"

func main() {
	fmt.Println(adder(1, 2, 3))
	fmt.Println(adder(9, 11))

	nums := []int{10, 20, 30}
	// fmt.Println(adder(nums)) - without the trailing ..., it will NOT compile..the
	fmt.Println(adder(nums...))
}

// By using ... before the type name of the last parameter we can indicate that
// it takes zero or more of those parameters. The function is invoked like any
// other function except we can pass as many arguments as we want.
func adder(args ...int) int {
	total := 0
	for _, v := range args {
		total += v

	}
	return total

	// return total
}
