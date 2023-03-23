package helper //So thsi file can contain some of the functions and code that can act as the helper functions for the main.go or any other go file present in the folder

import "strings" // which makes the     GO LANGUAGE    to share and use multiple files all at ones and use the function been present in one of the file to be used by the multiple

func ValidateUserInput(firstName string, lastName string, emailAddress string, userTickets uint, remainingTickets uint) (bool, bool, bool) { //files thus making the process much easier and faster

	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(emailAddress, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets //So by this we can make multiple files of the same code taht belongs / reference to the same package &
	return isValidName, isValidEmail, isValidTicketNumber                     // and make the process much more efficient
}
