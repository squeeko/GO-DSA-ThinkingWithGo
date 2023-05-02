package main

import "fmt"

func returnMultiple() (n int, s string) {
	n = 19
	s = "value 1"

	return
}

func main() {
	fmt.Println(returnMultiple())
}
