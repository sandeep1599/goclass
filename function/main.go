package main

import (
	"errors"
	"fmt"
)

func main() {

	isLoggedIn, err := UserAccess("bob", "1234")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("User can log in successfully", isLoggedIn)
}

func GreetORPrint(val string) {
	fmt.Println(val)
}

// bob, 1234
func UserAccess(username string, password string) (bool, error) {
	if username != "bob" {
		return false, errors.New("username is not bob")
	}

	if password != "1234" {
		return false, errors.New("password didn't match")
	}

	return true, nil
}

// Hello World!
// PQR
// ABC
