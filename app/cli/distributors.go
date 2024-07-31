package cli

import (
	"app/config"
	"app/handler"
	"fmt"
	"os"
)

func Distributor(cfg *config.Config) {
	var choice int
	fmt.Println("1. Add Distributor")
	fmt.Println("2. Edit Distributor")
	fmt.Println("3. Delete Distributor")
	fmt.Println("0. Exit")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		handler.AddDistributor(cfg)
	case 2:
		handler.EditDistributor(cfg)
	case 3:
		handler.DeleteDistributor(cfg)
	case 0:
		os.Exit(0)
	default:
		fmt.Println("Invalid choice. Please try again.")
	}
}