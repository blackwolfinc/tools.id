package handler

import (
	"app/config"
	"app/internal/categories"
	"app/internal/products"
	"fmt"
	"log"
	"strings"
)

// AddProduct adds a new product for a distributor
func AddProduct(cfg *config.Config, userID int) {
	// Retrieve distributor ID using userID
	distributorID, err := products.GetDistributorID(cfg, userID)
	if err != nil {
		log.Fatalf("Error retrieving distributor ID: %v", err)
		return
	}

	// Show available categories before inputting the product
	fmt.Println("=======================================================================================")
	fmt.Println("Show Category")
	categories.ShowCategory(cfg)
	fmt.Println("=======================================================================================")

	var name, description, size string
	var price float64
	var categoryID, stock int

	fmt.Print("Enter category ID: ")
	fmt.Scan(&categoryID)
	fmt.Print("Enter product name: ")
	fmt.Scan(&name)
	fmt.Print("Enter product description: ")
	fmt.Scan(&description)

	// Input validation for product size
	for {
		fmt.Print("Enter product size (S, M, L): ")
		fmt.Scan(&size)
		size = strings.ToUpper(size) // Ensure the size is in uppercase

		if size == "S" || size == "M" || size == "L" {
			break // Exit the loop if the size is valid
		} else {
			fmt.Println("Invalid size. Please enter 'S', 'M', or 'L'.")
		}
	}

	fmt.Print("Enter product stock: ")
	fmt.Scan(&stock)
	fmt.Print("Enter product price: ")
	fmt.Scan(&price)

	// Add the product using the distributorID
	products.AddProduct(cfg, name, description, size, price, stock, categoryID, distributorID)
}

func EditProduct(cfg *config.Config, userID int) {
	// Retrieve distributor ID using userID
	distributorID, err := products.GetDistributorID(cfg, userID)
	if err != nil {
		log.Fatalf("Error retrieving distributor ID: %v", err)
		return
	}

	var id int
	var name, description, size string
	var price float64
	var categoryID, stock int

	printHeader("Show Category")
	categories.ShowCategory(cfg)
	fmt.Println("========================================================================================================================")
	products.ViewProduct(cfg, distributorID, "View Product Before Editing")

	fmt.Print("Enter product ID to edit: ")
	if _, err := fmt.Scan(&id); err != nil {
		fmt.Println("Invalid input for product ID. Please enter a valid number.")
		return
	}

	fmt.Print("Enter new category ID: ")
	if _, err := fmt.Scan(&categoryID); err != nil {
		fmt.Println("Invalid input for category ID. Please enter a valid number.")
		return
	}

	fmt.Print("Enter new product name: ")
	if _, err := fmt.Scan(&name); err != nil {
		fmt.Println("Invalid input for product name. Please try again.")
		return
	}

	fmt.Print("Enter new product description: ")
	if _, err := fmt.Scan(&description); err != nil {
		fmt.Println("Invalid input for product description. Please try again.")
		return
	}

	for {
		fmt.Print("Enter new product size (S, M, L): ")
		if _, err := fmt.Scan(&size); err != nil {
			fmt.Println("Invalid input for size. Please enter 'S', 'M', or 'L'.")
			continue
		}
		size = strings.ToUpper(size) // Ensure the size is in uppercase

		if size == "S" || size == "M" || size == "L" {
			break // Exit the loop if the size is valid
		} else {
			fmt.Println("Invalid size. Please enter 'S', 'M', or 'L'.")
		}
	}

	fmt.Print("Enter new product stock: ")
	if _, err := fmt.Scan(&stock); err != nil {
		fmt.Println("Invalid input for stock. Please enter a valid number.")
		return
	}

	fmt.Print("Enter new product price: ")
	if _, err := fmt.Scan(&price); err != nil {
		fmt.Println("Invalid input for price. Please enter a valid number.")
		return
	}

	fmt.Println("=======================================================================================")
	products.EditProduct(cfg, id, name, description, size, price, stock, categoryID, distributorID)
}
func DeleteProduct(cfg *config.Config, userID int) {

	//fmt.Println("userID", userID)

	// Retrieve distributor ID using userID
	distributorID, err := products.GetDistributorID(cfg, userID)
	if err != nil {
		log.Fatalf("Error retrieving distributor ID: %v", err)
		return
	}

	products.ViewProduct(cfg, distributorID, "View Product")

	var id int
	fmt.Print("Enter product ID for delete: ")
	fmt.Scan(&id)

	fmt.Println("=======================================================================================")
	products.DeleteProduct(cfg, id, distributorID)
}

// ViewProduct handles the process of viewing a product's details
func ViewProduct(cfg *config.Config, userID int) {

	// Retrieve distributor ID using userID
	distributorID, err := products.GetDistributorID(cfg, userID)
	if err != nil {
		log.Fatalf("Error retrieving distributor ID: %v", err)
		return
	}

	products.ViewProduct(cfg, distributorID, "View Product")
}

func printHeader(title string) {
	fmt.Println(strings.Repeat("=", 120))
	fmt.Printf("%s\n", centerText(title, 120))
	fmt.Println(strings.Repeat("=", 120))
}

// centerText centers the text in a given width
func centerText(text string, width int) string {
	if len(text) >= width {
		return text
	}
	pad := (width - len(text)) / 2
	return fmt.Sprintf("%s%s%s", strings.Repeat(" ", pad), text, strings.Repeat(" ", pad))
}
