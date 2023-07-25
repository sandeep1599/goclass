package main

import "fmt"

func main() {

	// print 9 to 0 in descending order
	for i := 9; i >= 0; i-- {
		fmt.Println(i)
	}

	loop1()
	loop2()
}

// Print Hello world 5 times
func loop1() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Hello World!")
	}
}

// Print Hello world 5 times
func loop2() {
	var i int
	i = 1

	for i <= 5 {
		fmt.Println("Hello World! by loop 2")
		i++
	}
}
