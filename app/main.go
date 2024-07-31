package main

import (
	"app/config"
	"app/handler"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Initialize CLI application
func main() {

	cfg := config.LoadConfig()

	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	fmt.Println("Successfully connected to the database!")

	var choice int
	fmt.Println("Welcome to the CLI program")
	fmt.Println("1. Sign Up")
	fmt.Println("2. Log In")
	fmt.Print("Choose an option: ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		handler.Register(db)
	case 2:
		handler.Login(db)
	default:
		fmt.Println("Invalid choice")
	}
}
