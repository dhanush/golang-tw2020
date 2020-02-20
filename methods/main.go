package main

import "fmt"

type numbers struct {
	one   int
	two   int
	three int
}

func (n numbers) add() int {
	// fmt.Printf("%p\n", &n)
	return n.one + n.two + n.three
}

func (n *numbers) increaseOne(a int) {
	n.one = n.one + a
	fmt.Println("n.one in increaseOne", n.one)
}

func main() {
	n := numbers{1, 2, 3}
	// fmt.Printf("%p\n", &n)
	o := n.add()
	fmt.Println(o)

	n.increaseOne(10)
	fmt.Println(n.one)

	// o = n.add()
	// fmt.Println(o)
}
