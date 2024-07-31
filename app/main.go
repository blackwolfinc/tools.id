package main

import (
	"app/cli"
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
		log.Fatalf("Database is not reachable: %v", err)
	}

	fmt.Println("Database connection successful.")

	var choice int
	fmt.Println("Welcome to the CLI program")
	fmt.Println("1. Sign Up")
	fmt.Println("2. Log In")
	fmt.Print("Choose an option: ")

	fmt.Scanln(&choice)

	switch choice {
	case 1:
		handler.Register(db, cfg)
	case 2:
		var Role = handler.Login(db, cfg)
		if Role != "" {
			cli.HandleUserRole(Role, db, cfg)
		}
	default:
		fmt.Println("Invalid choice")
	}
}
