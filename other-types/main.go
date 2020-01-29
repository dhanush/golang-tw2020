package main

import "fmt"

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

	var m map[string]string

	fmt.Println(m)

	m = make(map[string]string)

	m["hello"] = "world"
	m["c"] = "language"
	m["hello"] = "world"
	fmt.Println(m)
	fmt.Println(m["hello"])

	var m2 = map[string]string{
		"hello boss": "world",
	}
	fmt.Println(m2)

	for k, v := range m {
		fmt.Printf("Key %s, Value %s\n", k, v)
	}
	//Arrays

	var threePrimes [3]int
	fmt.Println("Length of threePrimes ", len(threePrimes))
	threePrimes[0] = 1
	threePrimes[1] = 2
	threePrimes[2] = 3
	fmt.Println(threePrimes)
	fmt.Println("Length of threePrimes ", len(threePrimes))

	//slices
	var primes []int
	fmt.Println("Length of primes ", len(primes))
	primes = make([]int, 3)
	fmt.Println("Length of primes ", len(primes))
	primes[0] = 1
	primes[1] = 2
	primes[2] = 3
	fmt.Println(primes)
	// primes[4] = 5

	names := []string{"Dhanush", "Luka", "Krishnan"}
	fmt.Println("Length of names ", len(names))
	fmt.Println(names)

	fmt.Println(names[1:3])

	names = append(names, "Sneha", "Clair")
	fmt.Println(names)

	fmt.Println("Length of names ", len(names))

}
