package cli

import (
	"app/config"
	"database/sql"
	"fmt"
)

// HandleUserRole function to handle user roles
func HandleUserRole(role string, userID int, db *sql.DB, cfg *config.Config) {
	switch role {
	case "Admin":
		adminCLI(db, cfg)
	case "Customer":
		customerCLI(db, cfg, userID)
	case "Distributor":
		distributorCLI(db, cfg, userID)
	default:
		fmt.Println("Unknown role")
	}
}
