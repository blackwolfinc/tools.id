package cli

import (
	"app/config"
	"database/sql"
	"fmt"
)

// HandleUserRole function to handle user roles
func HandleUserRole(role string, db *sql.DB, cfg *config.Config) {
	switch role {
	case "Admin":
		adminCLI(db, cfg)
	case "Customer":
		customerCLI(db, cfg)
	case "Distributor":
		distributorCLI(db, cfg)
	default:
		fmt.Println("Unknown role")
	}
}
