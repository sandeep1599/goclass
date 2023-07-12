package auth

import "fmt"

func Signup(email, password, phone string) {
	registered := ifAlreadyRegistered()

	if registered == true {
		fmt.Println("THe user is already registered")
		return
	}

	fmt.Println("You have successfully signed up...")
}
