package products

import (
	"app/config"
	"database/sql"
	"fmt"
	"log"
	"strings"
)

// GetDistributorID retrieves the distributor_id for a given user_id
func GetDistributorID(cfg *config.Config, userID int) (int, error) {
	var distributorID int

	// Open a connection to the database
	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		return 0, fmt.Errorf("error connecting to database: %v", err)
	}
	defer db.Close()

	// Prepare the SQL query to fetch the distributor_id
	query := "SELECT distributor_id FROM Users WHERE user_id = ?"

	// Execute the query
	err = db.QueryRow(query, userID).Scan(&distributorID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("no distributor found for user_id %d", userID)
		}
		return 0, fmt.Errorf("failed to query distributor_id: %v", err)
	}

	return distributorID, nil
}

// AddProduct adds a new product to the database
func AddProduct(cfg *config.Config, name, description, size string, price float64, stock, categoryID, distributorID int) {
	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	// Updated SQL query to include stock and size
	_, err = db.Exec(
		"INSERT INTO products (name, description, size, price, stock, category_id, distributor_id) VALUES (?, ?, ?, ?, ?, ?, ?)",
		name, description, size, price, stock, categoryID, distributorID,
	)
	if err != nil {
		fmt.Println("=======================================================================================")
		fmt.Println("Error adding product:", err)
	} else {
		fmt.Println("=======================================================================================")
		fmt.Println("Product added successfully")
	}

	// Call the ViewProduct function to display the updated product list
	ViewProduct(cfg, distributorID, "View Product After Addition")
}

// EditProduct updates an existing product in the database
func EditProduct(cfg *config.Config, id int, name, description, size string, price float64, stock, categoryID, distributorID int) {
	//fmt.Println("id, distributorID", id, distributorID)

	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	// Updated SQL query to include size and stock
	result, err := db.Exec("UPDATE products SET name=?, description=?, size=?, price=?, stock=?, category_id=? WHERE product_id = ? AND distributor_id = ?", name, description, size, price, stock, categoryID, id, distributorID)
	if err != nil {
		fmt.Println("Error editing product:", err)
		return
	}

	// Check the number of affected rows
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error fetching number of affected rows:", err)
		return
	}

	if rowsAffected == 0 {
		fmt.Println("No product was updated. Please check the product ID and distributor ID.")
	} else {
		fmt.Println("Product edited successfully")
	}

	// Display updated product list
	ViewProduct(cfg, distributorID, "View Product Updated")
}

// DeleteProduct removes a product from the database
func DeleteProduct(cfg *config.Config, id int, distributorID int) {

	//fmt.Println("id, distributorID", id, distributorID)

	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
		return
	}
	defer db.Close()

	query := "DELETE FROM products WHERE product_id = ? AND distributor_id = ?"

	result, err := db.Exec(query, id, distributorID)
	if err != nil {
		fmt.Println("No product found with the specified ID and distributorID")
	} else {
		rowsAffected, _ := result.RowsAffected()
		if rowsAffected > 0 {
			fmt.Println("Product deleted successfully")
		} else {
			fmt.Println("No product found with the specified ID and distributorID")
		}
	}

	ViewProduct(cfg, distributorID, "View Product Updated")
}

// ViewProduct retrieves and displays a list of products for a given distributor
func ViewProduct(cfg *config.Config, distributorID int, headerReport string) {
	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT 
			products.product_id,
			categories.category_name AS category_name,
			products.name, 
			products.description,
			products.size,
			products.stock,
			products.price
		FROM 
			products 
		LEFT JOIN 
			categories 
		ON 
			products.category_id = categories.category_id 
		WHERE 
			products.distributor_id = ?
	`, distributorID)
	if err != nil {
		fmt.Println("Error querying products:", err)
		return
	}
	defer rows.Close()

	var (
		productID    int
		categoryName sql.NullString
		name         string
		description  string
		size         sql.NullString // Handle possible null values
		stock        int
		price        float64
	)

	// Header for the report
	printHeader(headerReport)
	fmt.Printf("%-5s %-10s %-20s %-25s %-10s %-10s %-10s %-50s\n", "No.", "Product ID", "Product Name", "Category", "Size", "Stock", "Price", "Description")
	fmt.Println("------------------------------------------------------------------------------------------------------------------------")

	index := 1 // Initialize the index counter
	for rows.Next() {
		err := rows.Scan(&productID, &categoryName, &name, &description, &size, &stock, &price)
		if err != nil {
			fmt.Println("Error scanning product:", err)
			return
		}

		// Format the output with index number and truncate the description
		fmt.Printf("%-5d %-10d %-20s %-25s %-10s %-10d %-10.2f %-50s\n", index, productID, name, categoryName.String, size.String, stock, price, truncate(description, 30))
		fmt.Println("------------------------------------------------------------------------------------------------------------------------")
		index++
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over rows:", err)
		return
	}
}

// Helper function to truncate description to a specific length
func truncate(text string, length int) string {
	if len(text) > length {
		return text[:length-3] + ".."
	}
	return text
}

// printHeader dynamically prints a header for reports
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
