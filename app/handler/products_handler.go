package handler

import (
	"app/config"
	"app/internal/products"
	"fmt"
)

func AddProduct(cfg *config.Config) {
	var name, description string
	var price float64
	var categoryID, distributorID int

	fmt.Print("Enter product name: ")
	fmt.Scan(&name)
	fmt.Print("Enter product description: ")
	fmt.Scan(&description)
	fmt.Print("Enter product price: ")
	fmt.Scan(&price)
	fmt.Print("Enter category ID: ")
	fmt.Scan(&categoryID)
	fmt.Print("Enter distributor ID: ")
	fmt.Scan(&distributorID)

	products.AddProduct(cfg, name, description, price, categoryID, distributorID)
}

func EditProduct(cfg *config.Config) {
	var id int
	var name, description string
	var price float64
	var categoryID, distributorID int

	fmt.Print("Enter product ID to edit: ")
	fmt.Scan(&id)
	fmt.Print("Enter new product name: ")
	fmt.Scan(&name)
	fmt.Print("Enter new product description: ")
	fmt.Scan(&description)
	fmt.Print("Enter new product price: ")
	fmt.Scan(&price)
	fmt.Print("Enter new category ID: ")
	fmt.Scan(&categoryID)
	fmt.Print("Enter new distributor ID: ")
	fmt.Scan(&distributorID)

	products.EditProduct(cfg, id, name, description, price, categoryID, distributorID)
}

func DeleteProduct(cfg *config.Config) {
	var id int
	fmt.Print("Enter product ID to delete: ")
	fmt.Scan(&id)

	products.DeleteProduct(cfg, id)
}
