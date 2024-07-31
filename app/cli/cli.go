package cli

import (
	"database/sql"
	"fmt"
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
	fmt.Println("1. View Products")
	fmt.Println("2. Place Order")
	fmt.Println("3. View Order Status")
	Order(db)
	// Add more customer options here
}

// Distributor CLI
func distributorCLI() {
	fmt.Println("Distributor CLI")
	fmt.Println("1. Add Products")
	fmt.Println("2. View Orders")
	// Add more distributor options here
}
