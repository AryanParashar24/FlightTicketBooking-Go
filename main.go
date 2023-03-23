package main

import ( //Now as we can see that we have moved hem from the main function to the package main and thus the whole program can get access to this function and not only one or two function because of ehihc we aren't required to define them again and again
	"booking-app/helper"
	"fmt"
	"sync"

	// "strconv"     as we are nt using the map anymorem so we don't need the strconv anymore and struct can be used instead alone
	"time"
)

// PACKAGE LEVEL VARIABLES ()
var conferenceName = "Go conference"
var remainingTickets uint = 50 //In order to keep the remaining ticket alwasy positive as u can also test by entering the negetive no. of ticket s remainning in the app and it will gove u an error
const conferenceTickets = 50

// var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	emailAddress    string
	NumberOftickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers() // As now here we have made the package variables so now the whole program has access to them and there is no need to define them again and again here or in any other function

	//for remainingTickets > 0 && len(bookings) < 50 {

	firstName, lastName, emailAddress, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, emailAddress, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTicket(userTickets, firstName, lastName, emailAddress)

		wg.Add(1)
		//wg.Add(2)
		go sendTicket(userTickets, firstName, lastName, emailAddress)
		//go doSmth()

		// fmt.Printf("The first value: %v\n", bookings[0])
		// fmt.Printf("SLices type: %T\n", bookings)
		// fmt.Printf("Slices.length: %v\n", len(bookings))

		firstNames := getFirstNames()
		fmt.Printf("the first name of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			//end Program
			fmt.Println("Our all tickets sold out so sry better luck nxt time")
			//break   not required if the for loop isn't been added in the code otherwise we don't need the wg function and this break function would ve been necessary as always
		}
	} else {
		if !isValidName {
			fmt.Println("first name or ur last name entered is wrong please check it out")
		}
		if !isValidEmail {
			fmt.Println("your entered email address is wrong please check")
		}
		if !isValidTicketNumber {
			fmt.Println("your entered ticet umber is wrong please check that out")
		}
	}
	//}
	wg.Wait()
}

func greetUsers() { // there is nothing in the inner paranthesis as the package level variables contain them all so we don;t require them to introduce any more and increase our code anybit
	fmt.Printf("Welcome to %v booking applicatrion\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available .\n", conferenceTickets, remainingTickets)
	fmt.Printf("get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		// booking["firstNames"]  or use it directly as in the code below as:--
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) { //Here we do not need anhy input from the us as the input is gping top come from the sueer and nt fropm us
	var firstName string
	var lastName string //BUT BUT BUT WE DO NEEED THE USER'S INPUTS AS WELL AS WE R GOING TO FUNCTIONONLY WITH TAHT INPUT BEEN PROVIDEDD BY THE USER THATSWHY WE HAVE TO RETRUN IT TO THE FUYNCTION MAIN
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

func bookTicket(userTickets uint, firstName string, lastName string, emailAddress string) {
	remainingTickets = remainingTickets - userTickets

	//craete a struct for a user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		emailAddress:    emailAddress,
		NumberOftickets: userTickets,
	}
	//craete a maop for a user
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["emailAddress"] = emailAddress
	// userData["NumberOftickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is : %v\n", bookings)

	fmt.Println("Get your tickets here to attend " + firstName + "" + lastName)
	fmt.Printf("thank you for booking %v Tickets with us\n", userTickets)
	fmt.Printf("An email of your tickets will be sended to your emailaddress %v\n", emailAddress)
	fmt.Printf("We have total of %v tickets remaining with us now", remainingTickets)

}

func sendTicket(userTickets uint, firstName string, lastName string, emailAddress string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###########")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, emailAddress)
	fmt.Println("###########")
	wg.Done()
}
