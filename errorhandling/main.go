package main

import (
	"errors"
	"fmt"
)

type numbers struct {
	one   int
	two   int
	three int
}

func (n numbers) add() (int, error) {
	if n.one < 0 || n.two < 0 || n.three < 0 {
		return 0, errors.New("One of the numbers is negative")
	}
	return n.one + n.two + n.three, nil
}

func main() {
	n := numbers{-11, 2, 3}
	if o, err := n.add(); err != nil {
		fmt.Println(err)
		// panic(err)
	} else {
		fmt.Println(o)
	}

	n1 := numbers{1, 2, 3}
	o, err := n1.add()
	fmt.Println(o)
	fmt.Println("Error is ", err)

}
