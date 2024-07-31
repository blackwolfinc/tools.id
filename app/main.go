package main

import (
	"app/cli"
	"app/config"
	"app/handler"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"bufio"
	"strings"
	"os"

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
	for {
		fmt.Println("Welcome to the CLI program")
		fmt.Println("1. Sign Up")
		fmt.Println("2. Log In")
		fmt.Print("Choose an option: ")
		reader := bufio.NewReader(os.Stdin)
		choiceInp, _ := reader.ReadString('\n')
		choiceInp = strings.TrimSpace(choiceInp)
		choice, err = strconv.Atoi(choiceInp)
		if err != nil || choice < 1 || choice > 2 {
			fmt.Println("Invalid input.")
			continue
		}
		break
	}


	switch choice {
	case 1:
		handler.Register(db, cfg)
	case 2:
		result := handler.Login(db, cfg)
		userId, _ := strconv.Atoi(result[1])
		if result[0] != "" {
			cli.HandleUserRole(result[0], userId, db, cfg)
		}
	default:
		fmt.Println("Invalid choice")
	}
}
