package main

import "fmt"

type numbers struct {
	one   int
	two   int
	three int
}

func (n numbers) add() int {
	return n.one + n.two + n.three
}

func (n *numbers) increaseOne(a int) {
	n.one = n.one + a
}

func main() {
	n := numbers{1, 2, 3}
	o := n.add()
	fmt.Println(o)

	n.increaseOne(10)
	fmt.Println(n.one)

	o = n.add()
	fmt.Println(o)
}
