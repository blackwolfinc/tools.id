package cli

import (
	"app/config"
	"app/handler"
	"fmt"
	"os"
)

func Product(cfg *config.Config, userID int) {
	var choice int
	fmt.Println("=======================================================================================")
	fmt.Println("Product Menu")
	fmt.Println("=======================================================================================")
	fmt.Println("1. Add Product")
	fmt.Println("2. View Product")
	fmt.Println("3. Edit Product")
	fmt.Println("4. Delete Product")
	fmt.Println("0. Exit")
	fmt.Println("=======================================================================================")

	fmt.Scan(&choice)

	switch choice {
	case 1:
		handler.AddProduct(cfg, userID)
	case 2:
		handler.ViewProduct(cfg, userID)
	case 3:
		handler.EditProduct(cfg, userID)
	case 4:
		handler.DeleteProduct(cfg, userID)
	case 0:
		os.Exit(0)
	default:
		fmt.Println("Invalid choice. Please try again.")
	}
}
