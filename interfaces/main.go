package main

import "fmt"

type shape interface {
	area() float32
}

type rectangle struct {
	length  float32
	breadth float32
}

type circle struct {
	radius float32
}

func (r rectangle) area() float32 {
	fmt.Printf("Calculating area of rectangle for length %f breadth %f \n", r.length, r.breadth)
	return r.length * r.breadth
}

func (c circle) area() float32 {
	fmt.Println("Calculating area od circle for radius ", c.radius)
	return 3.14 * c.radius * c.radius
}

func areaCalculation(shapes ...shape) {
	for _, s := range shapes {
		fmt.Println(s.area())
	}
}

type ABC struct{}

func emptyInterface(x interface{}) {
	// fmt.Println("Printing x", x)
	// str := x.(string)
	// fmt.Println("Printing x value as string", str)

	switch x.(type) {
	case string:
		fmt.Println("Printing x value as string", x.(string))
	case int:
		fmt.Println("Printing x value as int", x.(int))
	case float32:
		fmt.Println("Printing x value as float", x.(float32))
	case ABC:
		fmt.Println("Printing x value as ABC", x.(ABC))
	case map[string]string:
		fmt.Println("Printing x value as Map", x.(map[string]string))
	}
}

func main() {

	// var r, c shape
	// r = rectangle{length: 1.2, breadth: 2.3}
	// fmt.Println(r.area())

	// c = circle{4.5}
	// fmt.Println(c.area())

	// areaCalculation(c, r)

	// var sq shape
	// fmt.Println(sq.area())

	emptyInterface("dhanush")
	emptyInterface("1")
	emptyInterface(1)
	emptyInterface(float32(3.3))
	emptyInterface(map[string]string{"a": "b"})
}
