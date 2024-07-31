package cli

import (
	"app/config"
	"app/handler"
	"fmt"
	"os"
)

func Delivery(cfg *config.Config) {
	var choice int
	fmt.Println("=======================================================================================")
	fmt.Println("Delivery Menu")
	fmt.Println("=======================================================================================")
	fmt.Println("1. Add Delivery")
	fmt.Println("2. Edit Delivery")
	fmt.Println("3. Delete Delivery")
	fmt.Println("0. Exit")
	fmt.Println("=======================================================================================")

	fmt.Scan(&choice)

	switch choice {
	case 1:
		handler.AddDelivery(cfg)
	case 2:
		handler.EditDelivery(cfg)
	case 3:
		handler.DeleteDelivery(cfg)
	case 0:
		os.Exit(0)
	default:
		fmt.Println("Invalid choice. Please try again.")
	}
}
