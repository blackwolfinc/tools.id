package cli

import (
	"app/config"
	"app/handler"
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
	fmt.Println("=======================================================================================")
	fmt.Println("Admin Menu")
	fmt.Println("=======================================================================================")
	fmt.Println("1. Manage Category")
	fmt.Println("2. Manage Distributor")
	fmt.Println("3. Manage Delivery")
	fmt.Println("4. Manage Coupons")
	fmt.Println("0. Exit")
	fmt.Println("=======================================================================================")

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
	// fmt.Println("2. View Order History")
	reader := bufio.NewReader(os.Stdin)
	choiceInp, _ := reader.ReadString('\n')
	choiceInp = strings.TrimSpace(choiceInp)
	choice, err := strconv.Atoi(choiceInp)
	if err != nil || choice != 1 {
		fmt.Printf("Invalid input.\n")
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
func distributorCLI(db *sql.DB, cfg *config.Config, userId int) {
	fmt.Println("=======================================================================================")
	fmt.Println("Distributor Menu")
	fmt.Println("=======================================================================================")
	fmt.Println("1. Add/View/Edit/Delete Products")
	fmt.Println("2. View Order Products")
	fmt.Println("3. Report")
	fmt.Println("0. Exit")
	fmt.Println("=======================================================================================")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		Product(cfg, userId)
	case 2:
		handler.ViewOrderProducts(cfg, userId)
	case 3:
		ReportProducts(cfg, userId)
	case 0:
		os.Exit(0)
	default:
		fmt.Println("Invalid choice. Please try again.")
	}
}
