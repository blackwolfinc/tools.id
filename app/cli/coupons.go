package cli

import (
	"app/config"
	"app/handler"
	"fmt"
	"os"
)

func Coupons(cfg *config.Config) {
	var choice int
	fmt.Println("1. Add coupons")
	fmt.Println("2. Edit coupons")
	fmt.Println("3. Delete coupons")
	fmt.Println("0. Exit")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		handler.AddCoupons(cfg)
	case 2:
		handler.EditCoupons(cfg)
	case 3:
		handler.DeleteCoupons(cfg)
	case 0:
		os.Exit(0)
	default:
		fmt.Println("Invalid choice. Please try again.")
	}
}
