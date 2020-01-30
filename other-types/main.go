package main

import "fmt"

// type age int

type numbers struct {
	one   int
	two   int
	three int
}

func main() {

	n := numbers{
		one:   1,
		two:   2,
		three: 3,
	}
	fmt.Printf("The numbers are %d %d %d\n", n.one, n.two, n.three)

	n2 := numbers{4, 12, 23}
	fmt.Printf("The numbers are %d %d %d\n", n2.one, n2.two, n2.three)

	var m map[string]string
	m = make(map[string]string)

	fmt.Println(m)

	m["hello"] = "world"
	m["c"] = "language"
	m["hello123"] = "world123"
	fmt.Println(m)
	fmt.Println(m["hello"])

	// var m2 = map[string]string{
	// 	"hello boss": "world",
	// }
	// fmt.Println(m2)

	for k, v := range m {
		fmt.Printf("Key %s, Value %s\n", k, v)
	}
	//Arrays

	var threePrimes [3]int
	fmt.Println("Length of threePrimes ", len(threePrimes))
	threePrimes[0] = 1
	threePrimes[1] = 2
	threePrimes[2] = 3
	// threePrimes[3] = 3
	fmt.Println(threePrimes)
	fmt.Println("Length of threePrimes ", len(threePrimes))

	//slices
	var primes []int
	fmt.Println("Length of primes ", len(primes))
	primes = make([]int, 0)
	fmt.Println("Length of primes ", len(primes))
	primes[0] = 1
	primes[1] = 2
	primes[2] = 3
	primes = append(primes, 4, 5, 6, 7)
	fmt.Println(primes)

	names := []string{"Dhanush", "Luka", "Krishnan"}
	fmt.Println("Length of names ", len(names))
	fmt.Println(names)

	fmt.Println(names[1:3])

	names = append(names, "Sneha", "Clair")
	fmt.Println(names)

	fmt.Println("Length of names ", len(names))

}
