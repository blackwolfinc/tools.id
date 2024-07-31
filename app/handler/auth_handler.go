package handler

import (
	"app/cli"
	"app/pkg/models"
	"database/sql"
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func Register(db *sql.DB) {
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

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	_, err = db.Exec("INSERT INTO Users (email, password_hash, address, role, created_at) VALUES (?, ?, ?, ?, ?)", email, passwordHash, address, role, time.Now())
	if err != nil {
		log.Fatalf("Failed to insert user: %v", err)
	}

	fmt.Println("Sign up successful. Please log in.")
	Login(db)
}

func Login(db *sql.DB) {
	var email, password string
	fmt.Print("Enter email: ")
	fmt.Scanln(&email)
	fmt.Print("Enter password: ")
	fmt.Scanln(&password)

	var user models.User
	err := db.QueryRow("SELECT user_id, email, password_hash, role  FROM Users WHERE email = ?", email).Scan(&user.ID, &user.Email, &user.Password, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Invalid email or password.")
			return
		}
		log.Fatalf("Failed to query user: %v", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		fmt.Println("Invalid email or password.")
		return
	}

	fmt.Printf("Log in successful")
	cli.HandleUserRole(user.Role)
}
