package handler

import (
	"app/config"
	"app/internal/products"
	reportproducts "app/internal/report_product"
	"log"
)

func ViewOrderProducts(cfg *config.Config, userID int) {

	// Retrieve distributor ID using userID
	distributorID, err := products.GetDistributorID(cfg, userID)
	if err != nil {
		log.Fatalf("Error retrieving distributor ID: %v", err)
		return
	}

	reportproducts.ViewOrderProducts(cfg, distributorID)

}
