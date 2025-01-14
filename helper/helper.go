package helper

import "strings"

func ValidateUserInputs(firstname string, lastname string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstname) >= 2 && len(lastname) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	// We can return multiple values in func in GO
	return isValidName, isValidEmail, isValidTicketNumber
}
