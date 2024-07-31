package cli

import (
	"app/config"
	"app/handler"
	"fmt"
	"os"
)

func Product(cfg *config.Config) {
	var choice int
	fmt.Println("1. Add Product")
	fmt.Println("2. Edit Product")
	fmt.Println("3. Delete Product")
	fmt.Println("0. Exit")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		handler.AddProduct(cfg)
	case 2:
		handler.EditProduct(cfg)
	case 3:
		handler.DeleteProduct(cfg)
	case 0:
		os.Exit(0)
	default:
		fmt.Println("Invalid choice. Please try again.")
	}
}
