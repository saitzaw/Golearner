package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		if i%2 != 0 && i%3 != 0 {
			fmt.Println(i)
		}
	}

}
