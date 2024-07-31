package cli

import (
	"app/config"
	"app/handler"
	"app/internal/categories"
	"fmt"
	"os"
)

func Categories(cfg *config.Config) {
	var choice int
	fmt.Println("=======================================================================================")
	fmt.Println("Categories Menu")
	fmt.Println("=======================================================================================")
	categories.ShowCategory(cfg)
	fmt.Println("=======================================================================================")
	fmt.Println("1. Add Categories")
	fmt.Println("2. Edit Categories")
	fmt.Println("3. Delete Categories")
	fmt.Println("0. Exit")
	fmt.Println("=======================================================================================")
	fmt.Scan(&choice)
	fmt.Println("=======================================================================================")
	switch choice {
	case 1:
		handler.AddCategory(cfg)
	case 2:
		handler.EditCategory(cfg)
	case 3:
		handler.DeleteCategory(cfg)
	case 0:
		os.Exit(0)
	default:
		fmt.Println("Invalid choice. Please try again.")
	}
}
