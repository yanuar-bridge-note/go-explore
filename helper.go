package main

import "strings"

func validateUserInput(firstName string, lastName string, emailAddress string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(emailAddress, "@")
	isValidUserTickets := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidUserTickets
}
