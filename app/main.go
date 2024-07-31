package main

import (
	"fmt"
	"log"
	"os"

	"database/sql"

	"app/config"
	"app/handler"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Welcome to Tools.ID CLI Online Shop")

	// Load configuration
	cfg := config.LoadConfig()

	// Check database connection
	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Database is not reachable: %v", err)
	}

	fmt.Println("Database connection successful.")

	// CLI Menu
	for {
		fmt.Println("1. Add Product")
		fmt.Println("2. Edit Product")
		fmt.Println("3. Delete Product")
		fmt.Println("4. Add Category")
		fmt.Println("5. Edit Category")
		fmt.Println("6. Delete Category")
		fmt.Println("7. Add Distributor")
		fmt.Println("8. Edit Distributor")
		fmt.Println("9. Delete Distributor")
		fmt.Println("0. Exit")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			handler.AddProduct(cfg)
		case 2:
			handler.EditProduct(cfg)
		case 3:
			handler.DeleteProduct(cfg)
		case 4:
			handler.AddCategory(cfg)
		case 5:
			handler.EditCategory(cfg)
		case 6:
			handler.DeleteCategory(cfg)
		case 7:
			handler.AddDistributor(cfg)
		case 8:
			handler.EditDistributor(cfg)
		case 9:
			handler.DeleteDistributor(cfg)
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
