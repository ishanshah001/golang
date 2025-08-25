package main

import (
	"fmt"
	"time"
)

func display(str string) {
	for i := 0; i < 3; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(str)
	}
}

func main() {
	go display("Hello, Goroutine!") // Runs concurrently
	display("Hello, Main!")
}
