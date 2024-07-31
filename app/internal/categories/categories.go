package categories

import (
	"app/config"
	"app/entity"
	"database/sql"
	"fmt"
	"log"
)

func ShowCategory(cfg *config.Config) {
	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}

	catQuery := "SELECT category_id, category_name, description FROM Categories"
	catRows, err := db.Query(catQuery)
	if err != nil {
		log.Fatal(err)
	}

	categories := make([]*entity.Category, 0)
	for catRows.Next() {
		cat := new(entity.Category)
		err := catRows.Scan(&cat.ID, &cat.Name, &cat.Description)
		if err != nil {
			log.Fatal(err)
		}
		categories = append(categories, cat)
	}
	if err = catRows.Err(); err != nil {
		log.Fatal(err)
	}

	defer catRows.Close()
	defer db.Close()

	for _, c := range categories {
		fmt.Printf("%d. %s, Description: %s\n", c.ID, c.Name, c.Description)
	}

}

func AddCategory(cfg *config.Config, name, description string) {
	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO categories (category_name, description) VALUES (?, ?)", name, description)
	if err != nil {
		fmt.Println("Error adding category:", err)
	} else {
		fmt.Println("Category added successfully")
	}
}

func EditCategory(cfg *config.Config, id int, name, description string) {
	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	_, err = db.Exec("UPDATE categories SET category_name=?, description=? WHERE id=?", name, description, id)
	if err != nil {
		fmt.Println("Error editing category:", err)
	} else {
		fmt.Println("Category edited successfully")
	}
}

func DeleteCategory(cfg *config.Config, id int) {
	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM categories WHERE id=?", id)
	if err != nil {
		fmt.Println("Error deleting category:", err)
	} else {
		fmt.Println("Category deleted successfully")
	}
}
