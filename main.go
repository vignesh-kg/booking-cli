package main

import (
	"fmt"
	"booking-cli/constants"
)

const totalTickets = 50
var availableTickets = 50

func main() {
	fmt.Printf("Welcome to %v\n", constants.ConferenceName)
	fmt.Printf("We have total %v tickets. Available %v\n", totalTickets, availableTickets)
}