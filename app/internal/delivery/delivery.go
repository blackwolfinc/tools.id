package delivery

import (
	"database/sql"
	"fmt"

	"app/config"
)

func AddDelivery(cfg *config.Config, name, size, cost, estimated_date string) {
	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO Delivery (name,size, cost, estimated_date ) VALUES (?,?, ?, ?)", name, size, cost, estimated_date)
	if err != nil {
		fmt.Println("Error adding distributor:", err)
	} else {
		fmt.Println("Distributor added successfully")
	}
}

func EditDelivery(cfg *config.Config, id int, size, name, cost, estimated_date string) {
	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	_, err = db.Exec("UPDATE Delivery SET name=?,  size=?, cost=?, estimated_date=? WHERE id=?", name, size, cost, estimated_date, id)
	if err != nil {
		fmt.Println("Error editing distributor:", err)
	} else {
		fmt.Println("Distributor edited successfully")
	}
}

func DeleteDelivery(cfg *config.Config, id int) {
	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM Delivery WHERE id=?", id)
	if err != nil {
		fmt.Println("Error deleting distributor:", err)
	} else {
		fmt.Println("Distributor deleted successfully")
	}
}
