package main

import "fmt"

// higher order function: function that takes another function as an argument
func applyOperation(x, y int, operation func(int, int) int) int {
	return operation(x, y)
}

// first class function
func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func main() {
	fmt.Println(applyOperation(3, 4, add))      // 7
	fmt.Println(applyOperation(3, 4, multiply)) // 12
}
