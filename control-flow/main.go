package main

import "fmt"

var primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199}

func main() {

	defer func() {
		fmt.Println("bye bye")
	}()

	// defer os.Remove("file")

	for i := 0; i < len(primes); i++ {
		if i == 23 {
			fmt.Println("I found 23")
		} else if i == 29 {
			fmt.Println("I found 29")
		} else {
			//
		}
	}

	for i, prime := range primes {
		fmt.Println(i)
		if prime == 23 {
			fmt.Println("I found 23 again")
			break
		} else if prime == 29 {
			fmt.Println("I found 29 but will i continue?")
			continue
		}

	}

	x := 0

	switch x {
	case 0:
		fmt.Printf("Printing 0\n")
	case 10:
		fmt.Printf("Printing 10\n")
	default:
		fmt.Printf("Printing default\n")
	}

}
