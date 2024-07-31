package cli

import (
	"app/entity"
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func Order(db *sql.DB, user_id int) {
	fmt.Println(user_id)
	reader := bufio.NewReader(os.Stdin)
	now := time.Now()

    // Format the time for SQL TIMESTAMP insertion
    formattedTime := now.Format("2006-01-02 15:04:05")

    // Example query to insert into a table
    query := "INSERT INTO Orders (user_id, order_date) VALUES (?, ?)"
    _, err := db.Exec(query, user_id, formattedTime)
    if err != nil {
        log.Fatal(err)
    }
	// list category
	catQuery := "SELECT category_name, description FROM Categories"
	catRows, err := db.Query(catQuery)
	if err != nil {
		log.Fatal(err)
	}

	categories := make([]*entity.Category, 0)
	for catRows.Next() {
		cat := new(entity.Category)
		err := catRows.Scan(&cat.Name, &cat.Description)
		if err != nil {
			log.Fatal(err)
		}
		categories = append(categories, cat)
	}
	if err = catRows.Err(); err != nil {
		log.Fatal(err)
	}

	defer catRows.Close()
	// cli
	for {
		fmt.Println("=======================================================================================")
		fmt.Println("Please select a category:")
		for i, c := range categories {
			fmt.Printf("%d. %s, Description: %s\n", i+1, c.Name, c.Description)
		}
	
		length := len(categories)
		var choice int
		for {
			fmt.Printf("Please enter a number between 1 to %d: ", length)
			
			choiceInp, _ := reader.ReadString('\n')
			choiceInp = strings.TrimSpace(choiceInp)
			choice, err = strconv.Atoi(choiceInp)
			if err != nil || choice < 1 || choice > length {
				fmt.Println("Invalid input.")
				continue
			}
			break
		}

		// list produk
		productQuery := "SELECT p.name, p.description, p.price, d.distributor_id, p.size FROM Products p, Distributors d, Categories c WHERE p.distributor_id = d.distributor_id AND c.category_id = p.category_id AND c.category_id = ?"
		productRows, err := db.Query(productQuery, choice)
		if err != nil {
			log.Fatal(err)
		}

		products := make([]*entity.Product, 0)
		for productRows.Next() {
			item := new(entity.Product)
			err := productRows.Scan(&item.Name, &item.Description, &item.Price, &item.DistributorID, &item.Size)
			if err != nil {
				log.Fatal(err)
			}
			products = append(products, item)
		}
		if err = productRows.Err(); err != nil {
			log.Fatal(err)
		}

		defer productRows.Close()
		fmt.Println()
		fmt.Println("=======================================================================================")
		fmt.Printf("You selected the %s category. Below are the list of available products\n", categories[choice-1].Name)
		for i, p := range products {
			fmt.Printf("%d. %s\n", i+1, p.Name)
			fmt.Printf("Description: %s\n", p.Description)
			fmt.Printf("Price: $%.2f\n", p.Price)
			fmt.Printf("Size: %s\n", p.Size)
		}

		length = len(products)
		for {
			fmt.Printf("Please enter a number between 1 to %d: ", length)
			choiceInp, _ := reader.ReadString('\n')
			choiceInp = strings.TrimSpace(choiceInp)
			choice, err = strconv.Atoi(choiceInp)
			if err != nil || choice < 1 || choice > length {
				fmt.Println("Invalid input.")
				continue
			}
			break
		}
		
		var quantity int
		for {
			fmt.Printf("Please enter a quantity: ")
			choiceInp, _ := reader.ReadString('\n')
			choiceInp = strings.TrimSpace(choiceInp)
			quantity, err = strconv.Atoi(choiceInp)
			if err != nil || quantity < 1 || quantity > 10 {
				fmt.Printf("Invalid input. Please enter a number between 1 to 10\n")
				continue
			}
			break
		}
		fmt.Println()
		fmt.Println("=======================================================================================")
		fmt.Printf("You selected %s with a quantity of %d\n", products[choice-1].Name, quantity)

		fmt.Print("Would you like to make another order? (y/n) ")
		again, _ := reader.ReadString('\n')
		again = strings.TrimSpace(again)
		if again == "n" {
			break
		}
	}
}
