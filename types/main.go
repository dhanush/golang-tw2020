package main

import "fmt"

var name string = "ThoughtWorks"

func main() {
	x := 20
	var y = 20
	z := x + y
	fmt.Println(z)
	a := 23.0
	b := 34.7
	f := add(a, b)
	fmt.Printf("Sum is %0.2f", f)
	name := "Dhanush"
	// var name = "Dhanush"
	fmt.Println("Name is ", name)

}

//Add adds two numbers
func add(a, b float64) float64 {
	fmt.Println("Name is ", name)
	return a + b
}
