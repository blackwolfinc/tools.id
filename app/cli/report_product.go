package cli

import (
	"app/config"
	"app/handler"
	"fmt"
	"os"
)

func ReportProducts(cfg *config.Config, userID int) {
	var choice int
	fmt.Println("=======================================================================================")
	fmt.Println("Pilih Report Distributor")
	fmt.Println("=======================================================================================")
	fmt.Println("1. Products with the Most and least Stock")
	fmt.Println("2. Products with the Most Expensive and Cheapest Prices")
	fmt.Println("3. Largest and Smallest Product Sizes")
	fmt.Println("4. Best Selling Products")
	fmt.Println("5. Most Unsold Products")
	fmt.Println("0. Exit")
	fmt.Println("=======================================================================================")

	fmt.Scan(&choice)

	switch choice {
	case 1:
		handler.MostAndLeastStock(cfg, userID)
	case 2:
		handler.MostExpensiveAndCheapestPrices(cfg, userID)
	case 3:
		handler.LargestAndSmallest(cfg, userID)
	case 4:
		handler.BestSellingProducts(cfg, userID)
	case 5:
		handler.MostUnsoldProducts(cfg, userID)
	case 0:
		os.Exit(0)
	default:
		fmt.Println("Invalid choice. Please try again.")
	}
}
