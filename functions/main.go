package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func addExtra(x, y int, xFunc func(x int) int) int {
	return x + y + xFunc(20)
}

var addBy10 = func(x int) int {
	return 10 + x
}

func main() {
	fmt.Println(add(12, 12))

	o := addExtra(10, 20, addBy10)
	fmt.Println(o)

	func(name string) {
		fmt.Printf("Hello, %s\n", name)
	}("Dhanush")
}
