package cli

import (
	"app/config"
	"app/handler"
	"app/internal/coupons"
	"fmt"
	"os"
)

func Coupons(cfg *config.Config) {
	var choice int
	fmt.Println("=======================================================================================")
	fmt.Println("Coupons Menu")
	fmt.Println("=======================================================================================")
	coupons.ShowCoupons(cfg)
	fmt.Println("=======================================================================================")
	fmt.Println("1. Add Coupons")
	fmt.Println("2. Edit Coupons")
	fmt.Println("3. Delete Coupons")
	fmt.Println("0. Exit")
	fmt.Println("=======================================================================================")
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
