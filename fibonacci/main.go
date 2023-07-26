package main

import (
	"fmt"
	"package_learn/fibonacci"
)

func main() {
	fibSeries := fibonacci.FibonacciSeries(10)
	fmt.Printf("fibSeries: %v\n", fibSeries)
}
