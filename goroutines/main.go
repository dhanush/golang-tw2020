package main

func add(x, y int) int {
	return x + y
}

func addChan(x, y int, ch chan int) {
	o := x + y
	ch <- o
}

func addChan10(x, y int, ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- x + y
	}
	close(ch)
}

func main() {
	//normal call
	// o := add(1, 2)
	// fmt.Println(o)
	//goroutine
	// go add(3, 4)

	//goroutine in a loop
	// for i := 0; i < 10; i++ {
	// 	go add(i, i+1)
	// }
	// time.Sleep(1 * time.Second)

	//unbuffered channel

	// ch := make(chan int)
	// go addChan(1, 1, ch)
	// o := <-ch
	// fmt.Println(o)

	//buffered channels

	// ch2 := make(chan int, 2)
	// go addChan(1, 1, ch2)
	// go addChan(1, 1, ch2)
	// go addChan(1, 1, ch2)
	// go addChan(1, 1, ch2)
	// a1 := <-ch2
	// fmt.Println(a1)
	// a1 = <-ch2
	// fmt.Println(a1)
	// a1 = <-ch2
	// fmt.Println(a1)
	// a1 = <-ch2
	// fmt.Println(a1)

	//for-range
	// ch := make(chan int, 10)

	// go addChan10(10, 20, ch)

	// for i := range ch {
	// 	fmt.Println(i)
	// }
}
