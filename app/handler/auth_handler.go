package handler

import (
	"app/config"
	"app/internal/auth"
	"app/pkg/utils"
	"database/sql"
	"fmt"
	"log"
)

func Register(db *sql.DB, cfg *config.Config) {
	var email, password, address, role string
	fmt.Print("Enter email: ")
	fmt.Scanln(&email)
	fmt.Print("Enter password: ")
	fmt.Scanln(&password)
	fmt.Print("Enter address: ")
	fmt.Scanln(&address)

	fmt.Println("Choose role:")
	fmt.Println("1. Customer")
	fmt.Println("2. Distributor")
	var roleChoice int
	fmt.Print("Enter role choice: ")
	fmt.Scanln(&roleChoice)

	switch roleChoice {
	case 1:
		role = "Customer"
	case 2:
		role = "Distributor"
	default:
		fmt.Println("Invalid role choice")
		return
	}

	passwordHash, err := utils.HashPassword(password)

	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}
	auth.Register(cfg, email, passwordHash, address, role)
}

func Login(db *sql.DB, cfg *config.Config) []string {
	var email, password string
	fmt.Print("Enter email: ")
	fmt.Scanln(&email)
	fmt.Print("Enter password: ")
	fmt.Scanln(&password)
	result := auth.Login(cfg, password, email)

	if result[0] != "" {
		return []string{result[0], result[1]}
	} else {
		return []string{"", ""}
	}
}
