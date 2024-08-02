package reportproducts

import (
	"app/config"
	"database/sql"
	"fmt"
	"strings"
)

func ViewMostStock(cfg *config.Config, distributorID int) {
	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT 
			Products.product_id,
			COALESCE(Categories.category_name, 'N/A') AS category_name,
			Products.name, 
			COALESCE(Products.description, 'N/A') AS description,
			COALESCE(Products.size, 'N/A') AS size,
			COALESCE(Products.stock, 0) AS stock,
			Products.price
		FROM 
			Products 
		LEFT JOIN 
			Categories 
		ON 
			Products.category_id = Categories.category_id  
		WHERE 
			distributor_id = ? 
		ORDER BY 
			stock DESC 
		LIMIT 5
	`, distributorID)
	if err != nil {
		fmt.Println("Error querying most stocked products:", err)
		return
	}
	defer rows.Close()

	printHeader("Most Stocked Products")
	fmt.Printf("%-5s %-10s %-20s %-25s %-10s %-10s %-10s %-50s\n", "No.", "Product ID", "Product Name", "Category", "Size", "Stock", "Price", "Description")
	fmt.Println(strings.Repeat("-", 110))

	index := 1
	for rows.Next() {
		var (
			productID    int
			categoryName string
			name         string
			description  string
			size         string
			stock        int
			price        float64
		)
		if err := rows.Scan(&productID, &categoryName, &name, &description, &size, &stock, &price); err != nil {
			fmt.Println("Error scanning product:", err)
			return
		}

		// Display the data
		fmt.Printf("%-5d %-10d %-20s %-25s %-10s %-10d %-10.2f %-50s\n", index, productID, name, categoryName, size, stock, price, truncate(description, 30))
		fmt.Println(strings.Repeat("-", 110))
		index++
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over rows:", err)
	}
}

func ViewLeastStock(cfg *config.Config, distributorID int) {
	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT 
			Products.product_id,
			COALESCE(Categories.category_name, 'N/A') AS category_name,
			Products.name, 
			COALESCE(Products.description, 'N/A') AS description,
			COALESCE(Products.size, 'N/A') AS size,
			COALESCE(Products.stock, 0) AS stock,
			Products.price
		FROM 
			Products 
		LEFT JOIN 
			Categories 
		ON 
			Products.category_id = Categories.category_id  
		WHERE 
			distributor_id = ? 
		ORDER BY 
			stock ASC 
		LIMIT 5
	`, distributorID)
	if err != nil {
		fmt.Println("Error querying least stocked products:", err)
		return
	}
	defer rows.Close()

	printHeader("Least Stocked Products")
	fmt.Printf("%-5s %-10s %-20s %-25s %-10s %-10s %-10s %-50s\n", "No.", "Product ID", "Product Name", "Category", "Size", "Stock", "Price", "Description")
	fmt.Println(strings.Repeat("-", 110))

	index := 1
	for rows.Next() {
		var (
			productID    int
			categoryName string
			name         string
			description  string
			size         string
			stock        int
			price        float64
		)
		if err := rows.Scan(&productID, &categoryName, &name, &description, &size, &stock, &price); err != nil {
			fmt.Println("Error scanning product:", err)
			return
		}

		// Display the data
		fmt.Printf("%-5d %-10d %-20s %-25s %-10s %-10d %-10.2f %-50s\n", index, productID, name, categoryName, size, stock, price, truncate(description, 30))
		fmt.Println(strings.Repeat("-", 110))
		index++
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over rows:", err)
	}
}

func ViewMostExpensiveProducts(cfg *config.Config, distributorID int) {
	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT 
			Products.product_id,
			Categories.category_name AS category_name,
			Products.name, 
			Products.description,
			Products.size,
			Products.stock,
			Products.price
		FROM 
			Products 
		LEFT JOIN 
			Categories 
		ON 
			Products.category_id = Categories.category_id  
		WHERE 
			distributor_id = ? 
		ORDER BY 
			price DESC 
		LIMIT 5
	`, distributorID)
	if err != nil {
		fmt.Println("Error querying most expensive products:", err)
		return
	}
	defer rows.Close()

	printHeader("Most Expensive Products")
	fmt.Printf("%-5s %-10s %-20s %-25s %-10s %-10s %-10s %-50s\n", "No.", "Product ID", "Product Name", "Category", "Size", "Stock", "Price", "Description")
	fmt.Println(strings.Repeat("-", 110))

	index := 1
	for rows.Next() {
		var (
			productID    int
			categoryName sql.NullString
			name         string
			description  string
			size         sql.NullString
			stock        sql.NullInt64
			price        float64
		)
		if err := rows.Scan(&productID, &categoryName, &name, &description, &size, &stock, &price); err != nil {
			fmt.Println("Error scanning product:", err)
			return
		}
		// Handle NULL values by using the Valid field
		category := categoryName.String
		if !categoryName.Valid {
			category = "N/A"
		}
		productSize := size.String
		if !size.Valid {
			productSize = "N/A"
		}
		productStock := int(stock.Int64)
		if !stock.Valid {
			productStock = 0 // Or handle as needed
		}

		fmt.Printf("%-5d %-10d %-20s %-25s %-10s %-10d %-10.2f %-50s\n", index, productID, name, category, productSize, productStock, price, truncate(description, 30))
		fmt.Println(strings.Repeat("-", 110))
		index++
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over rows:", err)
	}
}

func ViewCheapestProducts(cfg *config.Config, distributorID int) {
	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT 
			Products.product_id,
			Categories.category_name AS category_name,
			Products.name, 
			Products.description,
			Products.size,
			Products.stock,
			Products.price
		FROM 
			Products 
		LEFT JOIN 
			Categories 
		ON 
			Products.category_id = Categories.category_id  
		WHERE 
			distributor_id = ? 
		ORDER BY 
			price ASC 
		LIMIT 5
	`, distributorID)
	if err != nil {
		fmt.Println("Error querying cheapest products:", err)
		return
	}
	defer rows.Close()

	printHeader("Cheapest Products")
	fmt.Printf("%-5s %-10s %-20s %-25s %-10s %-10s %-10s %-50s\n", "No.", "Product ID", "Product Name", "Category", "Size", "Stock", "Price", "Description")
	fmt.Println(strings.Repeat("-", 110))

	index := 1
	for rows.Next() {
		var (
			productID    int
			categoryName sql.NullString
			name         string
			description  string
			size         sql.NullString
			stock        sql.NullInt64
			price        float64
		)
		if err := rows.Scan(&productID, &categoryName, &name, &description, &size, &stock, &price); err != nil {
			fmt.Println("Error scanning product:", err)
			return
		}
		// Handle NULL values by using the Valid field
		category := categoryName.String
		if !categoryName.Valid {
			category = "N/A"
		}
		productSize := size.String
		if !size.Valid {
			productSize = "N/A"
		}
		productStock := int(stock.Int64)
		if !stock.Valid {
			productStock = 0 // Or handle as needed
		}

		fmt.Printf("%-5d %-10d %-20s %-25s %-10s %-10d %-10.2f %-50s\n", index, productID, name, category, productSize, productStock, price, truncate(description, 30))
		fmt.Println(strings.Repeat("-", 110))
		index++
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over rows:", err)
	}
}

func ViewLargestProducts(cfg *config.Config, distributorID int) {
	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT 
			Products.product_id,
			Categories.category_name AS category_name,
			Products.name, 
			Products.description,
			Products.size,
			Products.stock,
			Products.price
		FROM 
			Products 
		LEFT JOIN 
			Categories 
		ON 
			Products.category_id = Categories.category_id  
		WHERE 
			distributor_id = ? 
		ORDER BY 
			size DESC 
		LIMIT 5
	`, distributorID)
	if err != nil {
		fmt.Println("Error querying largest Products:", err)
		return
	}
	defer rows.Close()

	printHeader("Largest Products")
	fmt.Printf("%-5s %-10s %-20s %-25s %-10s %-10s %-10s %-50s\n", "No.", "Product ID", "Product Name", "Category", "Size", "Stock", "Price", "Description")
	fmt.Println(strings.Repeat("-", 110))

	index := 1
	for rows.Next() {
		var (
			productID    int
			categoryName sql.NullString
			name         string
			description  string
			size         sql.NullString
			stock        sql.NullInt64
			price        float64
		)
		if err := rows.Scan(&productID, &categoryName, &name, &description, &size, &stock, &price); err != nil {
			fmt.Println("Error scanning product:", err)
			return
		}
		// Handle NULL values by using the Valid field
		category := categoryName.String
		if !categoryName.Valid {
			category = "N/A"
		}
		productSize := size.String
		if !size.Valid {
			productSize = "N/A"
		}
		productStock := int(stock.Int64)
		if !stock.Valid {
			productStock = 0 // Or handle as needed
		}

		fmt.Printf("%-5d %-10d %-20s %-25s %-10s %-10d %-10.2f %-50s\n", index, productID, name, category, productSize, productStock, price, truncate(description, 30))
		fmt.Println(strings.Repeat("-", 110))
		index++
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over rows:", err)
	}
}

func ViewSmallestProducts(cfg *config.Config, distributorID int) {
	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT 
			Products.product_id,
			Categories.category_name AS category_name,
			Products.name, 
			Products.description,
			Products.size,
			Products.stock,
			Products.price
		FROM 
			Products 
		LEFT JOIN 
			Categories 
		ON 
			Products.category_id = Categories.category_id  
		WHERE 
			distributor_id = ? 
		ORDER BY 
			size ASC 
		LIMIT 5
	`, distributorID)
	if err != nil {
		fmt.Println("Error querying smallest Products:", err)
		return
	}
	defer rows.Close()

	printHeader("Smallest Products")
	fmt.Printf("%-5s %-10s %-20s %-25s %-10s %-10s %-10s %-50s\n", "No.", "Product ID", "Product Name", "Category", "Size", "Stock", "Price", "Description")
	fmt.Println(strings.Repeat("-", 110))

	index := 1
	for rows.Next() {
		var (
			productID    int
			categoryName sql.NullString
			name         string
			description  string
			size         sql.NullString
			stock        sql.NullInt64
			price        float64
		)
		if err := rows.Scan(&productID, &categoryName, &name, &description, &size, &stock, &price); err != nil {
			fmt.Println("Error scanning product:", err)
			return
		}
		// Handle NULL values by using the Valid field
		category := categoryName.String
		if !categoryName.Valid {
			category = "N/A"
		}
		productSize := size.String
		if !size.Valid {
			productSize = "N/A"
		}
		productStock := int(stock.Int64)
		if !stock.Valid {
			productStock = 0 // Or handle as needed
		}

		fmt.Printf("%-5d %-10d %-20s %-25s %-10s %-10d %-10.2f %-50s\n", index, productID, name, category, productSize, productStock, price, truncate(description, 30))
		fmt.Println(strings.Repeat("-", 110))
		index++
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over rows:", err)
	}
}

func ViewBestSellingProducts(cfg *config.Config, distributorID int) {
	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT 
			p.product_id, 
			p.name, 
			SUM(od.quantity) AS total_sold 
		FROM 
			Products p 
		JOIN 
			OrderDetails od 
		ON 
			p.product_id = od.product_id 
		JOIN 
			Orders o 
		ON 
			od.order_id = o.order_id 
		WHERE 
			p.distributor_id = ? 
		GROUP BY 
			p.product_id 
		ORDER BY 
			total_sold DESC 
		LIMIT 5
	`, distributorID)
	if err != nil {
		fmt.Println("Error querying best selling products:", err)
		return
	}
	defer rows.Close()

	printHeader("Best Selling Products")
	fmt.Printf("%-5s %-10s %-30s %-15s\n", "No.", "Product ID", "Product Name", "Total Sold")
	fmt.Println(strings.Repeat("-", 70))

	index := 1
	for rows.Next() {
		var productID, totalSold int
		var name string
		if err := rows.Scan(&productID, &name, &totalSold); err != nil {
			fmt.Println("Error scanning product:", err)
			return
		}
		fmt.Printf("%-5d %-10d %-30s %-15d\n", index, productID, name, totalSold)
		index++
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over rows:", err)
	}
}

func ViewMostUnsoldProducts(cfg *config.Config, distributorID int) {
	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT 
			p.product_id, 
			p.name, 
			IFNULL(SUM(od.quantity), 0) AS total_sold 
		FROM 
			Products p 
		LEFT JOIN 
			OrderDetails od 
		ON 
			p.product_id = od.product_id 
		LEFT JOIN 
			Orders o 
		ON 
			od.order_id = o.order_id 
		WHERE 
			p.distributor_id = ? 
		GROUP BY 
			p.product_id 
		HAVING 
			total_sold = 0
	`, distributorID)
	if err != nil {
		fmt.Println("Error querying most unsold products:", err)
		return
	}
	defer rows.Close()

	printHeader("Most Unsold Products")
	fmt.Printf("%-5s %-10s %-30s %-15s\n", "No.", "Product ID", "Product Name", "Total Sold")
	fmt.Println(strings.Repeat("-", 70))

	index := 1
	for rows.Next() {
		var productID, totalSold int
		var name string
		if err := rows.Scan(&productID, &name, &totalSold); err != nil {
			fmt.Println("Error scanning product:", err)
			return
		}
		fmt.Printf("%-5d %-10d %-30s %-15d\n", index, productID, name, totalSold)
		index++
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over rows:", err)
	}
}

func ViewOrderProducts(cfg *config.Config, distributorID int) {
	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT 
			p.product_id, 
			p.name, 
			SUM(od.quantity) AS total_sold 
		FROM 
			Products p 
		JOIN 
			OrderDetails od 
		ON 
			p.product_id = od.product_id 
		JOIN 
			Orders o 
		ON 
			od.order_id = o.order_id 
		WHERE 
			p.distributor_id = ? 
		GROUP BY 
			p.product_id 
	`, distributorID)
	if err != nil {
		fmt.Println("Error querying order products:", err)
		return
	}
	defer rows.Close()

	printHeader("View All Order Products")
	fmt.Printf("%-5s %-10s %-30s %-15s\n", "No.", "Product ID", "Product Name", "Total Sold")
	fmt.Println(strings.Repeat("-", 70))

	index := 1
	for rows.Next() {
		var productID, totalSold int
		var name string
		if err := rows.Scan(&productID, &name, &totalSold); err != nil {
			fmt.Println("Error scanning product:", err)
			return
		}
		fmt.Printf("%-5d %-10d %-30s %-15d\n", index, productID, name, totalSold)
		index++
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over rows:", err)
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
