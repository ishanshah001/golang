package main

import (
	"fmt"
	"os"
)

// panic("division by zero") starts the panic process.
// The defer runs before the function exits, calling our recovery logic.
// recover() stops the panic and returns the value passed to panic.
// The function continues after the deferred block — but note that if you panic before returning a value,
// you’ll get the zero value for the type.

func safeDivide(a, b int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	if b == 0 {
		panic("division by zero")
	}
	println("This wont be printed")
	return a / b
}

// function-> defer -> panic -> return
// recover would help program not crash if multiple call stacks.

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Start")

	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")

	fmt.Println(safeDivide(10, 0))

	fmt.Println("End")

	file, err := os.Open("nonexistent.txt")
	if err != nil {
		panic(err) // panic with the error object
	}
	defer file.Close()

	fmt.Println("File opened successfully")
}
