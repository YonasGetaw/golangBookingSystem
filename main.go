package main

import (
	"fmt"
	"strings"
)

func main() {
	const totalTickets int = 50
	var remainingTickets int = 50
	var bookings []string

	fmt.Println("##### Welcome to Go Conference! #####")
	fmt.Println("We have", totalTickets, "tickets in total.")
	fmt.Println("Currently, there are", remainingTickets, "tickets remaining.\n")

	for {
		var fullName string
		var email string
		var numberOfTickets int

		// 🧍 Get user input
		fmt.Print("Enter your full name: ")
		fmt.Scan(&fullName)

		fmt.Print("Enter your email: ")
		fmt.Scan(&email)

		fmt.Print("Enter number of tickets: ")
		fmt.Scan(&numberOfTickets)

		// ✅ VALIDATIONS
		isValidName := len(fullName) >= 3
		isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
		isValidTicketCount := numberOfTickets > 0 && numberOfTickets <= remainingTickets

		// ⚙️ Error handling
		if !isValidName {
			fmt.Println("❌ Invalid name! Your full name must be at least 3 characters long.")
			continue
		}

		if !isValidEmail {
			fmt.Println("❌ Invalid email address! It must contain '@' and a domain like '.com'.")
			continue
		}

		if !isValidTicketCount {
			fmt.Printf("❌ Invalid ticket number! You must book between 1 and %v tickets.\n", remainingTickets)
			continue
		}

		// ✅ If all validations pass
		bookings = append(bookings, fullName)
		remainingTickets -= numberOfTickets

		fmt.Printf("✅ Hello %v! You successfully booked %v ticket(s).\n", fullName, numberOfTickets)
		fmt.Printf("📩 Confirmation will be sent to %v\n", email)
		fmt.Printf("🎟️ Tickets remaining: %v\n", remainingTickets)
		fmt.Println("Current bookings list:", bookings)
		fmt.Println("---------------------------------------")

		// 🚫 Stop if all tickets sold
		if remainingTickets == 0 {
			fmt.Println("🚫 All tickets for the Go Conference are sold out!")
			break
		}
	}
}
