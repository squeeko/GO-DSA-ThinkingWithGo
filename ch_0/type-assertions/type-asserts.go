package main

import "fmt"

func main() {
	var i interface{} = "Go Lang type asssertions"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64)
	fmt.Println(f)
}

/*
Go Lang type asssertions

Go Lang type asssertions true

0 false

panic: interface conversion: interface {} is string, not float64

goroutine 1 [running]:
main.main()
        /Users/jallgood/GO-DSA-ThinkingWithGo/ch_0/type-assertions/type-asserts.go:17 +0x14c
exit status 2
*/
