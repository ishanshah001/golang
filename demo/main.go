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

	userName := "Tom" // variables only
	var userTicket = 2
	email := ""

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

}
