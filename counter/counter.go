package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Println("Enter a string:")
	fmt.Scanln(&input)

	freq := make(map[rune]int)

	for _, ch := range input {
		freq[ch]++
	}

	fmt.Println("Character frequencies:")
	for k, v := range freq {
		fmt.Printf("%q : %d\n", k, v)
	}
}
