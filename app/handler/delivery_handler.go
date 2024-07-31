package handler

import (
	"app/config"
	"app/internal/delivery"
	"app/internal/distributors"
	"fmt"
)

func AddDelivery(cfg *config.Config) {
	var name, size, cost, estimated_date string

	fmt.Print("Enter Delivery Name: ")
	fmt.Scan(&name)
	fmt.Print("Enter Delivery size S/M/L: ")
	fmt.Scan(&size)
	fmt.Print("Enter Delivery cost: ")
	fmt.Scan(&cost)
	fmt.Print("Enter Delivery estimated_date: ")
	fmt.Scan(&estimated_date)

	delivery.AddDelivery(cfg, name, size, cost, estimated_date)
}

func EditDelivery(cfg *config.Config) {
	var id int
	var name, size, cost, estimated_date string

	fmt.Print("Enter Delivery ID to edit: ")
	fmt.Scan(&id)
	fmt.Print("Enter Delivery Name: ")
	fmt.Scan(&name)
	fmt.Print("Enter new Delivery size: ")
	fmt.Scan(&size)
	fmt.Print("Enter new Delivery cost: ")
	fmt.Scan(&cost)
	fmt.Print("Enter new Delivery estimated_date: ")
	fmt.Scan(&estimated_date)

	delivery.EditDelivery(cfg, id, name, size, cost, estimated_date)
}

func DeleteDelivery(cfg *config.Config) {
	var id int
	fmt.Print("Enter Delivery ID to delete: ")
	fmt.Scan(&id)

	distributors.DeleteDistributor(cfg, id)
}
