package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func addExtra(x, y int, xFunc func(x, y int) int) int {
	return x + y + xFunc(20, 30)
}

var addBy10 = func(x, y int) int {
	return 10 + x + y
}

func main() {
	// fmt.Println(add(12, 12))

	// o := addExtra(10, 20, addBy10)
	// fmt.Println(o)

	func(name, x string) {
		fmt.Printf("Hello, %s\n", name)
	}("Dhanush", "G")
}
