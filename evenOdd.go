package main

import "fmt"

func main() {
	var sum int
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			sum += i
		} else if i%2 != 0 {
			sum -= i
		}
	}
	fmt.Printf("the difference of even and odd is: %d \n", sum)

}
