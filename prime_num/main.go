package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter the number")
	fmt.Scan(&num)
	isPrime := true
	for i := 2; i < num; i++ {
		if num%2 == 0 {
			isPrime = false
		}
	}
	if isPrime {
		fmt.Print(num, " is a prime")
	} else {
		fmt.Print(num, " is not a prime")
	}
}
