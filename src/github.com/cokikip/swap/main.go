package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

// Return more than one valuec
func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
