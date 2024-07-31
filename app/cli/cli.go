package cli

import (
	"app/config"
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ========================================================================
// Admin Menu
// ========================================================================

// Admin CLI
func adminCLI(db *sql.DB, cfg *config.Config) {
	fmt.Println("Admin CLI")
	fmt.Println("1. Manage Category")
	fmt.Println("2. Manage Distributor")
	fmt.Println("3. Manage Delivery")
	fmt.Println("4. Manage Coupons")
	// Add more admin options here

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		Categories(cfg)
	case 2:
		Distributor(cfg)
	case 3:
		Delivery(cfg)
	case 4:
		Coupons(cfg)
	case 0:
		os.Exit(0)
	default:
		fmt.Println("Invalid choice. Please try again.")
	}
}

// ========================================================================
// Customer Menu
// ========================================================================

// Customer CLI
func customerCLI(db *sql.DB, cfg *config.Config, userId int) {
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
		Order(db, userId)
	}
	// Add more customer options here
}

// ========================================================================
// Distributor Menu
// ========================================================================

// Distributor CLI
func distributorCLI(db *sql.DB, cfg *config.Config) {
	fmt.Println("Distributor CLI")
	fmt.Println("1. Add Products")
	fmt.Println("2. View Orders")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		Product(cfg)
	case 2:

	case 0:
		os.Exit(0)
	default:
		fmt.Println("Invalid choice. Please try again.")
	}
}
