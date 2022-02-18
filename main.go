package main

import (
	"fmt"
	"strings"
)

func main() {

	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	//create an array slice
	bookings := []string{}

	//placeholder to reference the type of our variable
	//fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T \n\n", conferenceName, conferenceTickets, remainingTickets)

	//placeholders to reference the value of our variable
	fmt.Printf("\nWelcome to our %v booking application \n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still availaible \n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend \n\n")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		//get user input
		fmt.Println("Enter Your First Name : ")
		fmt.Scan(&firstName)

		fmt.Println("Enter Your Last Name : ")
		fmt.Scan(&lastName)

		fmt.Println("Enter Your Email: ")
		fmt.Scan(&email)

		fmt.Println("How Many Tickets Do You Need : ")
		fmt.Scan(&userTickets)

		//validate user input
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidUserTickets := userTickets > 0 && userTickets <= remainingTickets



		//check if the user books more than the available ticket
		if isValidName && isValidEmail && isValidUserTickets { 
			//get remainingg tickets
			remainingTickets = remainingTickets - userTickets
			//store the values in a slice
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets rmaining for %v \n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of booking are : %v \n", firstNames)

			if remainingTickets == 0 {
				//end program
				fmt.Println("Our Conference is booked out. Come back next year")
				break
			}

		} else {
			if !isValidName {
				fmt.Printf("The First or Last name is too short\n")
			}
			if !isValidEmail {
				fmt.Printf("The email you entered does not have an @ sign\n")
			}
			
			if !isValidUserTickets {
				fmt.Printf("Number of ticket you entered is invalid\n")
			}
		
	}

  }
}
