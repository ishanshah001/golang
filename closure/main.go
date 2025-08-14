package main

import "fmt"

func counter() func() int {
	count := 0 // captured variable

	return func() int {
		count++ // modifies the captured variable
		return count
	}
}

func concatenator() func(string) string {
	result := "" // captured variable

	return func(s string) string {
		result += s
		return result
	}
}

func main() {
	c1 := counter()
	fmt.Println(c1()) // 1
	fmt.Println(c1()) // 2
	fmt.Println(c1()) // 3

	c2 := counter()   // separate closure with its own count
	fmt.Println(c2()) // 1

	concat := concatenator()

	fmt.Println(concat("Hello"))  // "Hello"
	fmt.Println(concat(" World")) // "Hello World"
	fmt.Println(concat("!!!"))    // "Hello World!!!"

	concat2 := concatenator()    // New independent concatenator
	fmt.Println(concat2("Go"))   // "Go"
	fmt.Println(concat2("Lang")) // "GoLang"
}
