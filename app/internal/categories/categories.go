package categories

import (
    "database/sql"
    "fmt"
    "app/config"
)

func AddCategory(cfg *config.Config, name, description string) {
    db, err := sql.Open("mysql", cfg.DSN())
    if err != nil {
        fmt.Println("Error connecting to database:", err)
        return
    }
    defer db.Close()

    _, err = db.Exec("INSERT INTO categories (name, description) VALUES (?, ?)", name, description)
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

    _, err = db.Exec("UPDATE categories SET name=?, description=? WHERE id=?", name, description, id)
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
