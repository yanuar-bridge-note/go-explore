package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
)

const conferenceTickets uint = 50

var remainingTickets = conferenceTickets
var conferenceName = "Go Conference"
var bookings = make([]map[string]string, 0)
var firstNames = []string{}

func main() {
	greetUser()
	for remainingTickets > 0 {
		firstName, lastName, emailAddress, userTickets := setUserInput()
		isValidName, isValidEmail, isValidUserTickets := helper.ValidateUserInput(firstName, lastName, emailAddress, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidUserTickets {
			saveBooking(userTickets, firstName, lastName, emailAddress)
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
	closingMessage()
}

func greetUser() {
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func setUserInput() (string, string, string, uint) {
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

	return firstName, lastName, emailAddress, userTickets
}

func saveBooking(userTickets uint, firstName string, lastName string, emailAddress string) {
	remainingTickets = remainingTickets - userTickets

	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["emailAddress"] = emailAddress
	userData["userTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v tickets.\nYou will recieve confirmation email at %v\n", firstName, lastName, userTickets, emailAddress)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	firstNames = append(firstNames, firstName)
}

func closingMessage() {
	fmt.Println("Thank you for booking your tickets with us")
	fmt.Println("We hope you enjoy the conference")
	fmt.Println("See you there")

	for _, firstName := range firstNames {
		fmt.Println("Thank you", firstName)
	}
}
