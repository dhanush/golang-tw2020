package main

import "fmt"

type shape interface {
	area() float32
}

type rectange struct {
	length  float32
	breadth float32
}

type circle struct {
	radius float32
}

func (r rectange) area() float32 {
	return r.length * r.breadth
}

func (c circle) area() float32 {
	return 3.14 * c.radius * c.radius
}

func areaCalculation(shapes ...shape) {
	for _, s := range shapes {
		fmt.Println(s.area())
	}
}

func emptyInterface(x interface{}) {
	fmt.Println("Printing x", x)
	// str := x.(string)
	// fmt.Println("Printing x value as string", str)

	switch x.(type) {
	case string:
		fmt.Println("Printing x value as string", x.(string))
	case int:
		fmt.Println("Printing x value as int", x.(int))
	case float32:
		fmt.Println("Printing x value as float", x.(float32))
	}
}

func main() {

	var r, c shape
	r = rectange{length: 1.2, breadth: 2.3}
	fmt.Println(r.area())

	c = circle{4.5}
	fmt.Println(c.area())

	areaCalculation(c, r)

	// var sq shape
	// fmt.Println(sq.area())

	emptyInterface("dhanush")
	emptyInterface("1")
	emptyInterface(1)
	emptyInterface(3)
}
