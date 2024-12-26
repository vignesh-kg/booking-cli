package main

import (
	"fmt"
	"booking-cli/constants"
	"booking-cli/structs"
)

const totalTickets = 50
var availableTickets = 50

func main() {
	fmt.Printf("Welcome to %v\n", constants.ConferenceName)
	fmt.Printf("We have total %v tickets. Available %v\n", totalTickets, availableTickets)
	fmt.Println("Please enter your userName:")
	var user structs.User
	fmt.Scan(&user.Username)
	fmt.Println("Please enter number of Tickets you wish to purchase:")
	fmt.Scan(&user.PurchasedTickets)
	fmt.Printf("User %v purchased %v tickets\n", user.Username, user.PurchasedTickets)
	fmt.Printf("Remaining Tickets : %d", availableTickets-int(user.PurchasedTickets))
}