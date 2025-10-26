package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Global constants and variables
const totalTickets int = 50
var remainingTickets int = totalTickets
var bookings []string

// 🏁 Main Entry Point
func main() {
	fmt.Println("#########################################")
	fmt.Println("🎉 Welcome to the Go Conference Booking App 🎉")
	fmt.Println("#########################################\n")

	greetUsers()

	// Infinite loop for multiple bookings
	for {
		fullName, email, numberOfTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketCount := validateInput(fullName, email, numberOfTickets)

		if !isValidName {
			fmt.Println("❌ Invalid name! Must be at least 3 characters long.")
			continue
		}
		if !isValidEmail {
			fmt.Println("❌ Invalid email! Please include '@' and a domain.")
			continue
		}
		if !isValidTicketCount {
			fmt.Printf("❌ Invalid number of tickets. You can book between 1 and %v.\n", remainingTickets)
			continue
		}

		bookTicket(fullName, email, numberOfTickets)
		displaySummary()

		if remainingTickets == 0 {
			fmt.Println("🚫 All tickets are sold out! Thank you for booking.")
			break
		}

		// ✅ Give users a way to exit gracefully
		fmt.Print("Do you want to make another booking? (yes/no): ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		answer := strings.ToLower(strings.TrimSpace(scanner.Text()))
		if answer != "yes" && answer != "y" {
			fmt.Println("\n👋 Exiting booking system. See you at the conference!")
			break
		}
	}
}

// 🗣️ Greet the user
func greetUsers() {
	fmt.Printf("We have %v tickets in total and %v tickets remaining.\n\n", totalTickets, remainingTickets)
}

// 🧍 Get user input with bufio (supports spaces)
func getUserInput() (string, string, int) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your full name: ")
	fullNameInput, _ := reader.ReadString('\n')
	fullName := strings.TrimSpace(fullNameInput)

	fmt.Print("Enter your email address: ")
	emailInput, _ := reader.ReadString('\n')
	email := strings.TrimSpace(emailInput)

	var numberOfTickets int
	fmt.Print("Enter number of tickets to book: ")
	fmt.Scan(&numberOfTickets)

	return fullName, email, numberOfTickets
}

// ✅ Validate user inputs
func validateInput(fullName string, email string, tickets int) (bool, bool, bool) {
	isValidName := len(fullName) >= 3
	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
	isValidTicketCount := tickets > 0 && tickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketCount
}

// 🎟️ Book tickets and update system
func bookTicket(fullName string, email string, numberOfTickets int) {
	remainingTickets -= numberOfTickets
	bookings = append(bookings, fullName)
	fmt.Printf("\n✅ Thank you, %v! You successfully booked %v ticket(s).\n", fullName, numberOfTickets)
	fmt.Printf("📩 A confirmation email has been sent to %v\n\n", email)
}

// 📊 Display current booking summary
func displaySummary() {
	fmt.Printf("🎟️ Tickets remaining: %v\n", remainingTickets)
	fmt.Printf("👥 Current Bookings: %v\n", bookings)
	fmt.Println("-------------------------------------------")
}
