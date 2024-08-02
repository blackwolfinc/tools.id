package handler

import (
	"app/config"
	"app/internal/products"
	reportproducts "app/internal/report_product"
	"log"
)

func MostAndLeastStock(cfg *config.Config, userID int) {

	// Retrieve distributor ID using userID
	distributorID, err := products.GetDistributorID(cfg, userID)
	if err != nil {
		log.Fatalf("Error retrieving distributor ID: %v", err)
		return
	}

	reportproducts.ViewMostStock(cfg, distributorID)
	reportproducts.ViewLeastStock(cfg, distributorID)

}

func MostExpensiveAndCheapestPrices(cfg *config.Config, userID int) {

	// Retrieve distributor ID using userID
	distributorID, err := products.GetDistributorID(cfg, userID)
	if err != nil {
		log.Fatalf("Error retrieving distributor ID: %v", err)
		return
	}

	reportproducts.ViewMostExpensiveProducts(cfg, distributorID)
	reportproducts.ViewCheapestProducts(cfg, distributorID)
}

func LargestAndSmallest(cfg *config.Config, userID int) {

	// Retrieve distributor ID using userID
	distributorID, err := products.GetDistributorID(cfg, userID)
	if err != nil {
		log.Fatalf("Error retrieving distributor ID: %v", err)
		return
	}

	reportproducts.ViewLargestProducts(cfg, distributorID)
	reportproducts.ViewSmallestProducts(cfg, distributorID)

}

func BestSellingProducts(cfg *config.Config, userID int) {

	// Retrieve distributor ID using userID
	distributorID, err := products.GetDistributorID(cfg, userID)
	if err != nil {
		log.Fatalf("Error retrieving distributor ID: %v", err)
		return
	}

	reportproducts.ViewBestSellingProducts(cfg, distributorID)

}

func MostUnsoldProducts(cfg *config.Config, userID int) {

	// Retrieve distributor ID using userID
	distributorID, err := products.GetDistributorID(cfg, userID)
	if err != nil {
		log.Fatalf("Error retrieving distributor ID: %v", err)
		return
	}

	reportproducts.ViewMostUnsoldProducts(cfg, distributorID)

}
