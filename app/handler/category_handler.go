package handler

import (
	"app/config"
	"app/internal/categories"
	"fmt"
)

func AddCategory(cfg *config.Config) {
	var name, description string

	fmt.Print("Enter category name: ")
	fmt.Scan(&name)
	fmt.Print("Enter category description: ")
	fmt.Scan(&description)

	categories.AddCategory(cfg, name, description)
}

func EditCategory(cfg *config.Config) {
	var id int
	var name, description string

	fmt.Print("Enter category ID to edit: ")
	fmt.Scan(&id)
	fmt.Print("Enter new category name: ")
	fmt.Scan(&name)
	fmt.Print("Enter new category description: ")
	fmt.Scan(&description)

	categories.EditCategory(cfg, id, name, description)
}

func DeleteCategory(cfg *config.Config) {
	var id int
	fmt.Print("Enter category ID to delete: ")
	fmt.Scan(&id)

	categories.DeleteCategory(cfg, id)
}
