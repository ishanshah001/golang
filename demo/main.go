package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println("Ishan Shah")
	fmt.Println("24 :)")
	fmt.Println("Python")

	var conferenceName string = "Go Conference"
	const CONFERENCE_TICKETS int = 50
	var remainingTickets int = CONFERENCE_TICKETS

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("Get your tickets!")
	fmt.Printf("Total tickets: %d. There are %d tickets remaining\n", CONFERENCE_TICKETS, remainingTickets)

	// variables
	userName := "Tom" // variables only
	var userTicket = 2
	email := ""

	// input output
	fmt.Printf("%s booked %d tickets.\n", userName, userTicket)
	fmt.Printf("userName is %T and userTicket is %T\n", userName, userTicket)

	fmt.Println("Enter user name: ")
	fmt.Scan(&userName)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTicket)

	fmt.Println("Enter email id: ")
	fmt.Scan(&email)

	fmt.Printf("%s booked %d tickets. You will receive a confirmation at %s\n", userName, userTicket, email)
	remainingTickets -= userTicket

	println(remainingTickets, "remain.")

	// arrays
	var bookings [50]string
	bookings[0] = userName
	bookings[1] = "Adam"
	fmt.Println(bookings[0])
	fmt.Printf("Bookings: %v\n", bookings)
	fmt.Printf("Length of array: %d\n", len(bookings))

	// slices
	var slices []string
	slices = append(slices, "Ishan")
	slices = append(slices, "Brett")
	fmt.Printf("Slices: %v\n", slices)
	fmt.Printf("Length of slice: %d\n", len(slices))

	// loops
	var numUsers int
	fmt.Println("Enter number of users")
	fmt.Scan(&numUsers)

	// standard for loop
	var name string
	for i := 0; i < numUsers; i++ {
		fmt.Println("Enter name")
		fmt.Scan(&name)
		slices = append(slices, name)
	}

	fmt.Println(slices)
}
