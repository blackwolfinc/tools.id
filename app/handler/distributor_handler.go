package handler

import (
	"app/config"
	"app/internal/distributors"
	"fmt"
)

func AddDistributor(cfg *config.Config) {
	var name, address, phone string

	fmt.Print("Enter distributor name: ")
	fmt.Scan(&name)
	fmt.Print("Enter distributor address: ")
	fmt.Scan(&address)
	fmt.Print("Enter distributor phone number: ")
	fmt.Scan(&phone)

	distributors.AddDistributor(cfg, name, address, phone)
}

func EditDistributor(cfg *config.Config) {
	var id int
	var name, address, phone string

	fmt.Print("Enter distributor ID to edit: ")
	fmt.Scan(&id)
	fmt.Print("Enter new distributor name: ")
	fmt.Scan(&name)
	fmt.Print("Enter new distributor address: ")
	fmt.Scan(&address)
	fmt.Print("Enter new distributor phone number: ")
	fmt.Scan(&phone)

	distributors.EditDistributor(cfg, id, name, address, phone)
}

func DeleteDistributor(cfg *config.Config) {
	var id int
	fmt.Print("Enter distributor ID to delete: ")
	fmt.Scan(&id)

	distributors.DeleteDistributor(cfg, id)
}
