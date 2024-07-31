package handler

import (
	"app/config"
	"app/internal/coupons"
	"fmt"
)

func AddCoupons(cfg *config.Config) {
	var coupon_code string
	var discount_amount float64

	fmt.Print("Enter Coupons Code ")
	fmt.Scan(&coupon_code)
	fmt.Print("Enter discount amount: ")
	fmt.Scan(&discount_amount)

	coupons.AddCoupons(cfg, coupon_code, discount_amount)
}

func EditCoupons(cfg *config.Config) {
	var id int
	var coupon_code string
	var discount_amount float64

	fmt.Print("Enter Delivery ID to edit: ")
	fmt.Scan(&id)
	fmt.Print("Enter Coupons Code ")
	fmt.Scan(&coupon_code)
	fmt.Print("Enter discount amount: ")
	fmt.Scan(&discount_amount)

	coupons.EditCoupons(cfg, id, coupon_code, discount_amount)
}

func DeleteCoupons(cfg *config.Config) {
	var id int
	fmt.Print("Enter Coupons ID to delete: ")
	fmt.Scan(&id)

	coupons.DeleteCoupons(cfg, id)
}
