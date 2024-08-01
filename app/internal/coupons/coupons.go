package coupons

import (
	"database/sql"
	"fmt"
	"log"

	"app/config"
	"app/entity"
)

func ShowCoupons(cfg *config.Config) {
	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}

	catQuery := "SELECT coupon_id, coupon_code, discount_amount FROM Coupons"
	catRows, err := db.Query(catQuery)
	if err != nil {
		log.Fatal(err)
	}
	coupons := make([]*entity.Coupons, 0)
	for catRows.Next() {
		cat := new(entity.Coupons)
		err := catRows.Scan(&cat.ID, &cat.CouponCode, &cat.DiscountAmount)
		if err != nil {
			log.Fatal(err)
		}
		coupons = append(coupons, cat)
	}
	if err = catRows.Err(); err != nil {
		log.Fatal(err)
	}

	defer catRows.Close()
	defer db.Close()

	for _, c := range coupons {
		fmt.Printf("%d. code : %s, deduction: %s\n", c.ID, c.CouponCode, c.DiscountAmount)
	}

}

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
