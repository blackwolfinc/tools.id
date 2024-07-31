package products

import (
	"app/config"
	"database/sql"
	"fmt"
)

func AddProduct(cfg *config.Config, name, description string, price float64, categoryID, distributorID int) {
	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO products (name, description, price, category_id, distributor_id) VALUES (?, ?, ?, ?, ?)", name, description, price, categoryID, distributorID)
	if err != nil {
		fmt.Println("Error adding product:", err)
	} else {
		fmt.Println("Product added successfully")
	}
}

func EditProduct(cfg *config.Config, id int, name, description string, price float64, categoryID, distributorID int) {
	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	_, err = db.Exec("UPDATE products SET name=?, description=?, price=?, category_id=?, distributor_id=? WHERE id=?", name, description, price, categoryID, distributorID, id)
	if err != nil {
		fmt.Println("Error editing product:", err)
	} else {
		fmt.Println("Product edited successfully")
	}
}

func DeleteProduct(cfg *config.Config, id int) {
	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM products WHERE id=?", id)
	if err != nil {
		fmt.Println("Error deleting product:", err)
	} else {
		fmt.Println("Product deleted successfully")
	}
}
