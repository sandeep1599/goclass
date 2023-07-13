package main

import "fmt"

type Rectangle struct {
	length float64
	width  float64
}

type Car struct {
	Name                string
	Brand               string
	IsAutomatic         bool
	IsElectric          bool
	Capacity            []string // no of people can sit in the car
	Color               string
	IsStarted           bool
	CurrentDrivingSpeed int
}

func main() {
	// I will be back in a moment
	c1 := &Car{
		Name:                "Range Rover",
		Brand:               "Land Rover",
		IsAutomatic:         true,
		IsElectric:          false,
		Color:               "black",
		Capacity:            []string{},
		CurrentDrivingSpeed: 0,
	}

	for {
		fmt.Println("\t\t\t\tEnter your choice\n")
		fmt.Println("\t\t\t\tgetin \t start \t drive")
		fmt.Println("\t\t\t\tleft \t right \t stop")
		fmt.Println("\t\t\t\tspeed-up \t slow-down \t park")

		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "getin":
			var person string
			fmt.Println("Enter the name of the person")
			fmt.Scanln(&person)
			c1.GetIntoTheCar(person)
		case "start":
			c1.Start()
		case "drive":
			c1.Drive()
		case "left":
			c1.Left()
		case "right":
			if c1.CurrentDrivingSpeed == 0 {
				fmt.Println("The car is stopped")
				continue
			}
			fmt.Println("We took right")
		case "stop":
			if c1.CurrentDrivingSpeed == 0 {
				fmt.Println("The car is already stopped")
				continue
			}
			c1.CurrentDrivingSpeed = 0
			fmt.Println("The car is stopped")
			c1.IsStarted = false
		case "speed-up":
			c1.SpeedUp(10)
		case "slowdown":
			c1.SlowDown(10)
		case "park":
			fmt.Println("parked!")
			return
		default:
			fmt.Println("The car doesn't have that feature!")
		}

	}
}

func (c *Car) Left() {
	if c.CurrentDrivingSpeed == 0 {
		fmt.Println("The car is stopped")
		return
	}
	fmt.Println("We took left")
}

func (c *Car) GetIntoTheCar(name string) {
	c.Capacity = append(c.Capacity, name)
	fmt.Printf("c.Capacity: %v\n", c.Capacity)
}

func (c *Car) Start() {
	if len(c.Capacity) > 0 {
		c.IsStarted = true
		fmt.Println("The car is started!")
		return
	}

	fmt.Println("There is no one in the car!")
}

func (c *Car) Drive() {
	if c.IsStarted == true {
		c.CurrentDrivingSpeed = 20
		fmt.Println("The car is driving")
		return
	}

	fmt.Println("The car isn;t started yet!")
}

func (c *Car) SpeedUp(kmph int) {
	c.CurrentDrivingSpeed = c.CurrentDrivingSpeed + kmph
	fmt.Printf("Current driving speed is %v", c.CurrentDrivingSpeed)
}

func (c *Car) SlowDown(kmph int) {
	c.CurrentDrivingSpeed = c.CurrentDrivingSpeed - kmph
	fmt.Printf("Current driving speed is %v", c.CurrentDrivingSpeed)
}
