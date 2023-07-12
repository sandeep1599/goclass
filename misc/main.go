package main

import "fmt"

type Rectangle struct {
	length float64
	width  float64
}

type Car struct {
	Name        string
	Brand       string
	IsAutomatic bool
	IsElectric  bool
	Capacity    []string // no of people can sit in the car
	Color       string
}

func main() {
	c1 := Car{
		Name:        "Range Rover",
		Brand:       "Land Rover",
		IsAutomatic: true,
		IsElectric:  false,
		Color:       "black",
		Capacity:    []string{"Dave"},
	}

	isStarted := c1.Start()
	if isStarted == false {
		return
	}

	c1.GoStraight()
	c1.TurnLeft()
}

func (c Car) Start() bool {
	if len(c.Capacity) == 0 {
		fmt.Println("The car cannot be started as there is no one inside!")
		return false
	}

	fmt.Printf("%v car started\n", c.Name)
	return true
}

func (c Car) GoStraight() {
	fmt.Printf("%v car is moving straight\n", c.Name)
}

func (c Car) TurnLeft() {
	fmt.Printf("%v car turned left\n", c.Name)
}

func (c Car) TurnRight() {
	fmt.Printf("%v car turned right\n", c.Name)
}
