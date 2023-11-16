package main

import "fmt"

func findMax(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func main() {
	// User input first number
	fmt.Println("Enter the first number: ")
	var firstNum int
	fmt.Scanln(&firstNum)

	// User input first number
	fmt.Println("Enter the second number: ")
	var secondNum int
	fmt.Scanln(&secondNum)

	// User input first number
	fmt.Println("Enter the third number: ")
	var thirdNum int
	fmt.Scanln(&thirdNum)

	// insert to the list
	var numbers []int
	numbers = append(numbers, firstNum)
	numbers = append(numbers, secondNum)
	numbers = append(numbers, thirdNum)
	maxNum := findMax(numbers)
	fmt.Printf("the maximum number is: %d \n", maxNum)

}
