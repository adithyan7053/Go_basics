package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still vailable.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")

	var userName string
	var userTickets int

	fmt.Scan(userName)

	userTickets = 2
	fmt.Printf("User %v has %v tickets", userName, userTickets)
}
