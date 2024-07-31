package cli

import (
	"app/entity"
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"

)

// -- Customer
// -- Belanja
// -- Login


func Order(db *sql.DB) {
	// -- pilih kategory
	fmt.Println("Please select a category:")
	query := "SELECT category_name, description FROM Categories"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	categories := make([]*entity.Category, 0)
	for rows.Next() {
		cat := new(entity.Category)
		err := rows.Scan(&cat.Name, &cat.Description)
		if err != nil {
			log.Fatal(err)
		}
		categories = append(categories, cat)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	for i, c := range categories {
		fmt.Printf("%d. %s, Description: %s\n", i + 1, c.Name, c.Description)
	}

	length := len(categories)
	fmt.Printf("Please enter a number between 1 to %d: ", length)
	reader := bufio.NewReader(os.Stdin)
	choiceInp, _ := reader.ReadString('\n')
	choiceInp = strings.TrimSpace(choiceInp)
	choice, err := strconv.Atoi(choiceInp)
	if err != nil || choice < 1 || choice > length {
		fmt.Printf("Invalid input. Please enter a number between 1 to %d\n", length)
		return
	}

	rows.Close()

	
// -- pilih produk
	query = "SELECT p.name, p.description, p.price, d.name, p.size FROM Products p, Distributors d, Categories c WHERE p.distributor_id = d.distributor_id AND c.category_id = p.category_id AND c.category_id = ?"
	rows, err = db.Query(query, choice)
	if err != nil {
		log.Fatal(err)
	}

	products := make([]*entity.Product, 0)
	for rows.Next() {
		item := new(entity.Product)
		err := rows.Scan(&item.Name, &item.Description, &item.Price, &item.Distributor, &item.Size)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, item)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	fmt.Println("=======================================================================================")
	fmt.Printf("You selected the %s category. Below are the list of available products\n", categories[choice - 1].Name)
	for i, p := range products {
		fmt.Printf("%d. %s\n", i + 1, p.Name)
		fmt.Printf("Description: %s\n", p.Description)
		fmt.Printf("Price: $%.2f\n", p.Price)
		fmt.Printf("Size: %s\n", p.Size)
	}

	length = len(products)
	fmt.Printf("Please enter a number between 1 to %d: ", length)
	choiceInp, _ = reader.ReadString('\n')
	choiceInp = strings.TrimSpace(choiceInp)
	choice, err = strconv.Atoi(choiceInp)
	if err != nil || choice < 1 || choice > length {
		fmt.Printf("Invalid input. Please enter a number between 1 to %d\n", length)
		return
	}
	fmt.Println()
	fmt.Println("=======================================================================================")
	fmt.Printf("You selected %s. Please choose a delivery method\n", products[choice - 1].Name)

	rows.Close()

// -- pilih delivery 
// -- pilih pembayaran 
// -- input kupon (optional)
// -- bayar
// -- recipt & shiping estimate
	
	// fmt.Println("Select product")
	// fmt.Println("Select product")
	// fmt.Println("Select product")
	// fmt.Println("Select product")
	// fmt.Println("Select product")
	// fmt.Println("2. Most Popular Game Report")
	// fmt.Println("3. Total Revenue Per Game Report")
	// fmt.Println("4. Player Count Per Game Report")
	// fmt.Print("Enter the number of the report you want to generate: ")
	// reader := bufio.NewReader(os.Stdin)
	// choiceInp, _ := reader.ReadString('\n')
	// choiceInp = strings.TrimSpace(choiceInp)
	// choice, err := strconv.Atoi(choiceInp)
	// if err != nil || choice < 1 || choice > 4 {
	// 	fmt.Println("Invalid input. Please enter a number between 1 to 4")
	// 	return
	// }
	// switch choice {
	// case 1:
	// 	handler.ShowTotalGameSalesReport(db)
	// case 2:
	// 	handler.ShowMostPopularGameReport(db)
	// case 3:
	// 	handler.ShowTotalRevenuePerGameReport(db)
	// case 4:
	// 	handler.ShowPlayerCountPerGameReport(db)
	// }

}