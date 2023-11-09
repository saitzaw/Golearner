package main

import "fmt"

func SimpleLogin() {
	// User input - User name
	fmt.Println("Enter the user name: ")
	var name string
	fmt.Scanln(&name)

	// User input - Password
	fmt.Println("Enter the password: ")
	var password string
	fmt.Scanln(&password)

	// Conditional check here

	if name == "admin" && password == "password123" {
		fmt.Println("Login successful")
	} else {
		fmt.Println("Login fail")
	}
}

func main() {
	SimpleLogin()
}
