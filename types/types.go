package main

import "fmt"

func main() {
	x := 20
	var y = 20
	z := x + y
	fmt.Println(z)
	a := 23.0
	b := 34.7
	f := add(a, b)
	fmt.Printf("Sum is %0.2f", f)
}

func add(a, b float64) float64 {
	return a + b
}
