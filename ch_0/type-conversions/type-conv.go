package main

import (
	"fmt"
	"math"
)

func main() {
	var i int = 19
	var f float64 = float64(i)
	// var u int = uint(f) - CANNOT backward convert

	fmt.Println(i, f)

	// easier declaration

	j := 20
	k := float64(j)
	fmt.Println(j, k)

	x, y := 2, 3
	f = math.Sqrt(float64(x*x + y*y))
	// z uint = uint(f) - CANNOT backward convert
	fmt.Println(x, y, f)

}
