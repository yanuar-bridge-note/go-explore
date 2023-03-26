package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	remainingTickets := conferenceTickets

	bookings := []string{}
	firstNames := []string{}

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for remainingTickets > 0 {
		var firstName string
		var lastName string
		var emailAddress string
		var userTickets uint

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&emailAddress)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(emailAddress, "@")
		isValidUserTickets := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidUserTickets {

			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets.\nYou will recieve confirmation email at %v\n", firstName, lastName, userTickets, emailAddress)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames = append(firstNames, firstName)

		} else if !isValidName {
			fmt.Println("First name and Last name should be more than 2 characters, please try it again")
			continue

		} else if !isValidEmail {
			fmt.Println("Your email format seems incorrect, please try it again")
			continue
		} else if !isValidUserTickets {
			fmt.Printf("Sorry tickets that you bought is over capacity, please enter below or same as %v. Thank you\n", remainingTickets)
			continue
		}
	}

	fmt.Printf("These first name of bookings for %v are %v\n", conferenceName, firstNames)
}
