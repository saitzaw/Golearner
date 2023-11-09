package main

import (
	"errors"
	"fmt"
)

type EmptyStringError struct {
	Message string
}

func (e *EmptyStringError) Error() string {
	return e.Message
}

func simpleLogin() error {

	// User input - User name
	fmt.Println("Enter the user name: ")
	var name string
	fmt.Scanln(&name)

	if len(name) == 0 {
		return &EmptyStringError{"Empty String Error!"}
	}

	// User input - Password
	fmt.Println("Enter the password: ")
	var password string
	fmt.Scanln(&password)

	if len(password) == 0 {
		return &EmptyStringError{"Empty String Error!"}
	}

	// Conditional check here
	if name == "admin" && password == "password123" {
		fmt.Println("Login successful")
		return nil
	} else {
		fmt.Println("Login fail")
		return errors.New("Log in Fail")
	}
}

func main() {
	simpleLogin()
}
