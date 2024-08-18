package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const totalTickets uint = 100

var conferenceName string = "GO Conference"
var remainingTickets uint = 10

// var bookings = make([]map[string]string, 0) // map
var bookings = make([]UserData, 0) // struct

type UserData struct {
	firstname string
	lastname  string
	email     string
	tickets   uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	for {
		var firstname, lastname, email, userTickets = getUserInputs()
		var isValidName, isValidEmail, isValidTicketNumber = helper.ValidateUserInputs(firstname, lastname, email, userTickets, remainingTickets)

		if userTickets > remainingTickets {
			fmt.Printf("There are only %v tickets remaining, so you can't book %v tickets.", remainingTickets, userTickets)
			continue
		}

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTickets(firstname, lastname, email, userTickets)
			wg.Add(1)
			go sendTickets(userTickets, firstname, lastname, email)

			fmt.Printf("The whole bookings %v \n", bookings)
			fmt.Printf("The whole first names %v \n", getFirstnames())

			// fmt.Printf("The first bookings %v \n", bookings[0])
			// fmt.Printf("The type bookings %T \n", bookings)
			// fmt.Printf("The length bookings %v \n", len(bookings))

			if remainingTickets == 0 {
				fmt.Printf("Our conference tickets are all sold out, Come back next year! \n")
				break
			}
		} else {
			if !isValidName {
				fmt.Printf("The name you entered is too short \n")
			}
			if !isValidEmail {
				fmt.Printf("The email you entered is not correct \n")
			}
			if !isValidTicketNumber {
				fmt.Printf("The ticket number you entered is invalid \n")
			}
		}

	}
	wg.Wait()
}

func greetUsers() {
	fmt.Println("Welcome to our", conferenceName, "booking application!")
	fmt.Printf("We have total of %v tickets and %v are still availables. \n", totalTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend \n")
}

func getFirstnames() []string {

	firstnames := []string{}

	for _, booking := range bookings {
		firstnames = append(firstnames, booking.firstname)
	}

	return firstnames

}

func getUserInputs() (string, string, string, uint) {
	firstname := ""
	lastname := ""
	email := ""

	var userTickets uint = 0

	fmt.Printf("Enter your first name: \n")
	fmt.Scan(&firstname)

	fmt.Printf("Enter your last name: \n")
	fmt.Scan(&lastname)

	fmt.Printf("Enter your email: \n")
	fmt.Scan(&email)

	fmt.Printf("Enter your tickets: \n")
	fmt.Scan(&userTickets)

	return firstname, lastname, email, userTickets
}

func bookTickets(firstname string, lastname string, email string, userTickets uint) {
	remainingTickets -= userTickets
	// bookings = append(bookings, firstname+" "+lastname)

	// creating a map for user
	// var userData = make(map[string]string)
	// userData["firstname"] = firstname
	// userData["lastname"] = lastname
	// userData["tickets"] = strconv.FormatUint(uint64(userTickets), 10) // type conversion for same data types in map

	// creating struct
	var userData = UserData{
		firstname: firstname,
		lastname:  lastname,
		email:     email,
		tickets:   userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("Thank you for your booking! %v %v booked %v tickets! \n", firstname, lastname, userTickets)
	fmt.Printf("There are only %v remaining tickets now! \n", remainingTickets)

}

func sendTickets(userTickets uint, firstname string, lastname string, email string) {
	time.Sleep(10 * time.Second) // delaying 10s
	var ticket = fmt.Sprintf("%v Tickets for %v %v", userTickets, firstname, lastname)
	fmt.Printf("************* \n")
	fmt.Printf("Sending ticket ~ \n %v \n to email address %v \n", ticket, email)
	fmt.Printf("************* \n")
	wg.Done()
}
