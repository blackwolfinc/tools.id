package cli

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"strconv"
)

// HandleUserRole function to handle user roles
func HandleUserRole(role string, db *sql.DB) {
	switch role {
	case "Admin":
		adminCLI()
	case "Customer":
		customerCLI(db)
	case "Distributor":
		distributorCLI()
	default:
		fmt.Println("Unknown role")
	}
}

// Admin CLI
func adminCLI() {
	fmt.Println("Admin CLI")
	fmt.Println("1. Manage Users")
	fmt.Println("2. Manage Products")
	fmt.Println("3. Manage Orders")
	// Add more admin options here
}

// Customer CLI
func customerCLI(db *sql.DB) {
	fmt.Println("Customer CLI")
	fmt.Println("1. Place Order")
	fmt.Println("2. View Order History")
	reader := bufio.NewReader(os.Stdin)
	choiceInp, _ := reader.ReadString('\n')
	choiceInp = strings.TrimSpace(choiceInp)
	choice, err := strconv.Atoi(choiceInp)
	if err != nil || choice < 1 || choice > 2 {
		fmt.Printf("Invalid input. Please enter a number between 1 to 2\n")
		return
	}
	if choice == 1 {
		Order(db)
	}
	// Add more customer options here
}

// Distributor CLI
func distributorCLI() {
	fmt.Println("Distributor CLI")
	fmt.Println("1. Add Products")
	fmt.Println("2. View Orders")
	// Add more distributor options here
}
