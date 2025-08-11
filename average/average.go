package main

import "fmt"

func main() {
	var nums []float64

	var len int
	fmt.Println("Enter length of nums")
	fmt.Scanf("%d", &len)

	for i := 0; i < len; i++ {
		fmt.Println("Enter number: ")
		var num float64
		fmt.Scanf("\n%f", &num)
		nums = append(nums, num)
	}

	sum := 0.0
	for i := range len {
		sum += nums[i]
	}

	fmt.Printf("Average is %.2f", sum/float64(len))
}
