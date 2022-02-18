package main

import (
	"fmt"
	"strings"
)

func main(){

    conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	//create an array slice
	bookings := []string{}

	//placeholder to reference the type of our variable
	//fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T \n\n", conferenceName, conferenceTickets, remainingTickets) 

	//placeholders to reference the value of our variable
	fmt.Printf("Welcome to our %v booking application \n", conferenceName)
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
		 

		//get remainingg tickets
		remainingTickets =  remainingTickets - userTickets	
		//store the values in a slice
		bookings = append(bookings, firstName + " " + lastName)


		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)

		firstNames := []string{}
		for index, booking := range bookings{
			var names = strings.Fields(booking)
			firstNames = append(firstNames, name[0])
		}
		fmt.Printf("This are all the bookings in our application : %v \n", bookings)

	}

	

}




