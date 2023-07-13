package main

import "fmt"

func main() {
	var msg string           // ""
	var age int              // 0
	var isOld bool           // false
	var cashInPocket float64 // 0.0

	msg = "Hello Dave!"
	age = 25
	isOld = false
	cashInPocket = 100.50

	fmt.Println(age)

	age = 30

	fmt.Println(msg)
	fmt.Println(age)
	fmt.Println(isOld)
	fmt.Println(cashInPocket)
}
