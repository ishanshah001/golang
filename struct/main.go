package main

import "fmt"

type User struct {
	firstName  string
	lastName   string
	tickets    uint
	newsletter bool
}

var bookings []User

func main() {
	var fname string
	var lname string
	var tickets uint
	var num int
	fmt.Println("Enter num of users ")
	fmt.Scanf("%d", &num)

	for i := 0; i < num; i++ {

		fmt.Println("Enter fname")
		fmt.Scan(&fname)
		fmt.Println("Enter lname")
		fmt.Scan(&lname)
		fmt.Println("Enter tickets")
		fmt.Scanf("\n%d", &tickets)
		var userData = User{
			firstName:  fname,
			lastName:   lname,
			tickets:    tickets,
			newsletter: true,
		}
		bookings = append(bookings, userData)
	}

	fmt.Println(bookings)

	for _, users := range bookings {
		fmt.Printf("First name: %s\nLast name: %s\nNumber of tickets booked: %d\nOpted for newsletter: %v\n", users.firstName, users.lastName, users.tickets, users.newsletter)
	}

}
