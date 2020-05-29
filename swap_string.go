package main

import "fmt"

var va int = 99
var vb int = 88

func swap(x, y string) (string, string) {
	return y, x
}

func print() {
	a, b, c := 123, 456, "hi"
	fmt.Println(a, b, c)
}

func main() {
	a, b := swap("Mahesh", "Kumar")
	fmt.Println(a, b)

	fmt.Println(va, vb)

	print()
}
