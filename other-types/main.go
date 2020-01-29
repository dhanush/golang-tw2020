package main

import "fmt"

type numbers struct {
	one   int
	two   int
	three int
}

func main() {

	// n := numbers{
	// 	one:   1,
	// 	two:   2,
	// 	three: 3,
	// }
	// fmt.Printf("The numbers are %d %d %d\n", n.one, n.two, n.three)

	// var m map[string]string

	// fmt.Println(m)

	// m = make(map[string]string)

	// m["hello"] = "world"
	// fmt.Println(m)
	// fmt.Println(m["hello"])

	// var m2 = map[string]string{
	// 	"hello boss": "world",
	// }
	// fmt.Println(m2)

	//Arrays

	var threePrimes [3]int
	threePrimes[0] = 1
	threePrimes[1] = 2
	threePrimes[2] = 3
	fmt.Println(threePrimes)

	//slices
	var primes []int
	primes = make([]int, 3)
	primes[0] = 1
	primes[1] = 2
	primes[2] = 3
	fmt.Println(primes)
	// primes[4] = 5

	names := []string{"Dhanush", "Luka", "Krishnan"}
	fmt.Println(names)

	fmt.Println(names[1:3])

	names = append(names, "Sneha", "Clair")
	fmt.Println(names)

}
