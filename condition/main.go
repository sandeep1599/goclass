package main

import "fmt"

func main() {
	var username string
	var password string

	fmt.Printf("Please enter the username: ")
	fmt.Scanln(&username)
	fmt.Printf("Enter the password: ")
	fmt.Scanln(&password)

	// If username is "sai123" and password is Secret then login
	if username != "sai123" {
		fmt.Println("Username is incorrect")
		return
	}

	if password != "Secret" {
		fmt.Println("Password is incorrect")
		return
	}

	fmt.Printf("%v has logged in to the system\n", username)
}

func ElseIf(name string) {
	if name == "Dave" {
		fmt.Println("Dave has logged in!")
	} else if name == "Sai" {
		fmt.Println("Sai has logged in!")
	} else if name == "Raju" {
		fmt.Println("Raju has logged in!")
	} else if name == "Bob" {
		fmt.Println("Bob has logged in!")
	} else if name == "Smith" {
		fmt.Println("Smith has logged in!")
	} else {
		fmt.Println("Cannot find the valid user!")
	}
}

func SwitchCase(name string) {
	switch name {
	case "Dave":
		fmt.Println("Dave has logged in!")
	case "Sai":
		fmt.Println("Sai has logged in!")
	case "Bob":
		fmt.Println("Bob has logged in!")
	default:
		fmt.Println("Cannot find the valid user!")
	}
}
