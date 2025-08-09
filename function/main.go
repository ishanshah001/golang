package main

import (
	"errors"
	"fmt"
	"math"
)

// error checking
type divideError struct {
	dividend float64
}

// error checking
func (de divideError) Error() string {
	return fmt.Sprintf("Cannot divide %.2f by zero", de.dividend)
}

func sub(x int, y int) int {
	return x - y
}

func concat(x string, y string) string {
	return x + y
}

func getPoint(x int, y int) (int, int) {
	return x, y
}

func nakedReturn() (x int, y int) {
	return
}

func hypotenuse(a, b float64) (c float64) {
	c = math.Sqrt(a*a + b*b)
	return // short, clear naked return
}

// error checking
func divide(a, b float64) (result float64, err error) {
	if b == 0 {
		err = divideError{dividend: a}
		return
	}
	result = a / b
	return
}

func lengthCheck(a string) (length int, err error) {
	if len(a) < 10 {
		err = errors.New("length is too short")
		return
	}
	length = len(a)
	return
}

func helloworld() {
	fmt.Println("Hello world")
}

func main() {
	ans := sub(2, 1)
	ct := concat("hello", "world")
	x, _ := getPoint(1, 2)
	helloworld()
	fmt.Println(ans)
	fmt.Println(ct)
	fmt.Println(x)
	fmt.Println(nakedReturn())
	fmt.Println(hypotenuse(3, 4))
	fmt.Println(divide(5, 2))
	fmt.Println(divide(5, 0))
	q, err := divide(6, 0)
	if err == nil {
		fmt.Println(q)
	} else {
		fmt.Println(err)
	}

	length, err := lengthCheck("abcde")
	if err == nil {
		fmt.Printf("Length is %d", length)
	} else {
		fmt.Println(err)
	}
}
