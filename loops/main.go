package main

import "fmt"

func main() {

	// standard loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// while loop
	count := 0
	for count < 5 {
		fmt.Println(count)
		count++
	}

	// infinite loop
	for {
		fmt.Println("Forever...")
	}

}
