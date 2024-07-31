package auth

import (
	"app/config"
	"app/pkg/models"
	"app/pkg/utils"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"
)

func Login(cfg *config.Config, Password, Email string) []string {
	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return []string{"", ""}
	}
	defer db.Close()

	var user models.User

	err = db.QueryRow("SELECT user_id, email, password_hash, role  FROM Users WHERE email = ?", Email).Scan(&user.ID, &user.Email, &user.Password, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Invalid email or password.")
			return []string{"", ""}
		}
		log.Fatalf("Failed to query user: %v", err)
	}
	err = utils.CheckPasswordHash(user.Password, Password)

	if err != nil {
		fmt.Println("Invalid email or password.")
		return []string{"", ""}
	}

	fmt.Printf("Log in successful")
	fmt.Printf("User Id: %s", strconv.Itoa(user.ID))
	return []string{user.Role, strconv.Itoa(user.ID)}
}

func Register(cfg *config.Config, email, passwordHash, address, role string) {

	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO Users (email, password_hash, address, role, created_at) VALUES (?, ?, ?, ?, ?)", email, passwordHash, address, role, time.Now())
	if err != nil {
		log.Fatalf("Failed to insert user: %v", err)
	}

	fmt.Println("Sign up successful. Please log in.")
	if err != nil {
		fmt.Println("Invalid email or password.")
		return
	}

	Login(cfg, passwordHash, email)
}
