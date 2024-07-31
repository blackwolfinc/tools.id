package coupons

import (
	"database/sql"
	"fmt"

	"app/config"
)

// coupon_code Ascending 1
// discount_amount

func AddCoupons(cfg *config.Config, coupon_code string, discount_amount float64) {
	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO Coupons (coupon_code , discount_amount ) VALUES (?, ?, ?)", coupon_code, discount_amount)
	if err != nil {
		fmt.Println("Error adding distributor:", err)
	} else {
		fmt.Println("Distributor added successfully")
	}
}

func EditCoupons(cfg *config.Config, id int, coupon_code string, discount_amount float64) {
	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	_, err = db.Exec("UPDATE Coupons SET coupon_code=?, discount_amount=?, WHERE id=?", coupon_code, discount_amount, id)
	if err != nil {
		fmt.Println("Error editing distributor:", err)
	} else {
		fmt.Println("Distributor edited successfully")
	}
}

func DeleteCoupons(cfg *config.Config, id int) {
	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM Coupons WHERE id=?", id)
	if err != nil {
		fmt.Println("Error deleting distributor:", err)
	} else {
		fmt.Println("Distributor deleted successfully")
	}
}
