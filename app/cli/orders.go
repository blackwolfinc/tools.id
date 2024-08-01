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
	query := "INSERT INTO Orders (user_id, order_date) VALUES (?, ?);"
	_, err := db.Exec(query, user_id, formattedTime)
	if err != nil {
		log.Fatal(err)
	}

	query = "SELECT order_id FROM Orders WHERE user_id = ? AND order_date = ?;"
	rows, err := db.Query(query, user_id, formattedTime)
	if err != nil {
		log.Fatal(err)
	}
	var orderID int
	for rows.Next() {
		err := rows.Scan(&orderID)
		if err != nil {
			log.Fatal(err)
		}
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

	orders := make([]*entity.OrderDetail, 0)

	sizes := []string{"S", "M", "L"}
	maxSize := 0

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
		
		
		productQuery := "SELECT p.product_id, p.name, p.description, p.price, d.distributor_id, p.size FROM Products p, Distributors d, Categories c WHERE p.distributor_id = d.distributor_id AND c.category_id = p.category_id AND c.category_id = ?"
		productRows, err := db.Query(productQuery, choice)
		if err != nil {
			log.Fatal(err)
		}

		products := make([]*entity.Product, 0)
		for productRows.Next() {
			item := new(entity.Product)
			err := productRows.Scan(&item.ID, &item.Name, &item.Description, &item.Price, &item.DistributorID, &item.Size)
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
		order := new(entity.OrderDetail)
		order.OrderID = orderID
		order.ProductID = products[choice-1].ID
		order.ProductName = products[choice-1].Name
		order.ProductSize = products[choice-1].Size
		order.Quantity = quantity
		order.TotalPrice = products[choice-1].Price * float64(quantity)

		query := "INSERT INTO OrderDetails (order_id, product_id, quantity, total_price) VALUES (?, ?, ?, ?)"
		_, err = db.Exec(query, orderID, order.ProductID, quantity, order.TotalPrice)
		if err != nil {
			log.Fatal(err)
		}
		orders = append(orders, order)
		if again == "n" {
			break
		}
	}
	fmt.Println()
	fmt.Println("=======================================================================================")
	fmt.Println("Here are a list of your orders")
	sumTotalPrice := 0.0
	for i, order := range orders {
		sumTotalPrice += order.TotalPrice 
		fmt.Printf("%d. %s, quantity: %d, total price = %.2f\n", i+1, order.ProductName, order.Quantity, order.TotalPrice)
		for i, s := range(sizes) {
			if order.ProductSize == s {
				if i > maxSize {
					maxSize = i
				}
			}
		} 
	}

	// checkout
	fmt.Println()
	fmt.Println("=======================================================================================")
	fmt.Println("Choose a delivery method")
	query = "SELECT delivery_id, name, size, cost, delivery_days FROM Delivery"
	rows, err = db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	deliveries:= make([]*entity.Delivery, 0)
	for rows.Next() {
		item := new(entity.Delivery)
		err := rows.Scan(&item.ID, &item.Name, &item.Size, &item.Cost, &item.DeliveryDays)
		if err != nil {
			log.Fatal(err)
		}
		deliveries = append(deliveries, item)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	fmt.Println()
	fmt.Println("=======================================================================================")
	fmt.Printf("Your maximum size item is %s. Below are the list of available choices of delivery:\n", sizes[maxSize])
	length := 0
	ind := 1
	available_deliveries := []*entity.Delivery{}
	chosenDelivery := new(entity.Delivery)
	for _, d := range deliveries {
		for i, s := range(sizes) {
			if d.Size == s {
				if i >= maxSize {
					fmt.Printf("%d. %s\n", ind, d.Name)
					fmt.Printf("Size: %s\n", d.Size)
					fmt.Printf("Cost: $%.2f\n", d.Cost)
					fmt.Printf("Delivery Days: %d\n", d.DeliveryDays)
					length++
					ind++
					available_deliveries = append(available_deliveries, d)
				}
			}
		} 
	}

	for {
		fmt.Printf("Please enter a number between 1 to %d: ", length)
		choiceInp, _ := reader.ReadString('\n')
		choiceInp = strings.TrimSpace(choiceInp)
		choice, err := strconv.Atoi(choiceInp)
		if err != nil || choice < 1 || choice > length {
			fmt.Println("Invalid input.")
			continue
		}
		chosenDelivery = available_deliveries[choice - 1]
		break
	}

	query = "UPDATE Orders SET delivery_id = ? WHERE user_id = ? AND order_date = ?;"
	_, err = db.Exec(query, chosenDelivery.ID, user_id, formattedTime)
	if err != nil {
		log.Fatal(err)
	}
	
	sumTotalPrice += chosenDelivery.Cost

	fmt.Println()
	fmt.Println("=======================================================================================")
	fmt.Println("Would you like to input a coupon? (y/n)")
	inp, _ := reader.ReadString('\n')
	inp = strings.TrimSpace(inp)
	var validCode string
	var couponID int
	appliedDiscount := 0.0
	valid := false
	if inp == "y" {
		fmt.Print("Enter a coupon code: ")
		code, _ := reader.ReadString('\n')
		code = strings.TrimSpace(code)
		query = "SELECT coupon_id, coupon_code, discount_amount FROM Coupons"
		couponRows, err := db.Query(query)
		if err != nil {
			log.Fatal(err)
		}
		
		for couponRows.Next() {
			err := couponRows.Scan(&couponID, &validCode, &appliedDiscount)
			if err != nil {
				log.Fatal(err)
			}
			if validCode == code {
				valid = true
				break
			}
		}
		if err = couponRows.Err(); err != nil {
			log.Fatal(err)
		}
		defer couponRows.Close()
	}

	if valid {
		fmt.Printf("Code valid. Discount applied (%d%%)", int(appliedDiscount))
		query := "UPDATE Orders SET coupon_id = ? WHERE user_id = ? AND order_date = ?;"
		_, err := db.Exec(query, couponID, user_id, formattedTime)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("Your code is not valid. No discount applied")
	}

	

	query = "UPDATE Orders SET total_amount = ? WHERE user_id = ? AND order_date = ?;"
	_, err = db.Exec(query, sumTotalPrice*(100.0-appliedDiscount)/100.0, user_id, formattedTime)
	if err != nil {
		log.Fatal(err)
	}


	payment_methods := []string{"Credit Card", "Virtual Account", "Cash on Delivery"}

	fmt.Println("=======================================================================================")
	fmt.Printf("Your total fee is %.2f\n", sumTotalPrice)
	fmt.Println("Choose a payment method")
	fmt.Println("1. Credit card")
	fmt.Println("2. Virtual Account")
	fmt.Println("3. Cash on Delivery")
	var payChoice int
	for {
		fmt.Printf("Please enter a number between 1 to 3: ")
		choiceInp, _ := reader.ReadString('\n')
		choiceInp = strings.TrimSpace(choiceInp)
		payChoice, err = strconv.Atoi(choiceInp)
		if err != nil || payChoice < 1 || payChoice> 3 {
			fmt.Printf("Invalid input.\n")
			continue
		}
		break
	}
	fmt.Println("=======================================================================================")
	fmt.Println("Here is a summary of your order: ")
	for i, order := range orders {
		fmt.Printf("%d. %s, quantity: %d, total price = %.2f\n", i+1, order.ProductName, order.Quantity, order.TotalPrice)
	}
	fmt.Printf("Delivery name: %s\n", chosenDelivery.Name)
	fmt.Printf("Delivery cost: %.2f\n", chosenDelivery.Cost)
	newDate := now.Add(time.Duration(chosenDelivery.DeliveryDays) * 24 * time.Hour)
	formattedDate := newDate.Format("2006-01-02")
	fmt.Printf("Estimated arrival: %s\n", formattedDate)
	fmt.Printf("Subtotal: %.2f\n", sumTotalPrice)
	if valid {
		fmt.Printf("Coupon applied. Discount: %d%%\n", int(appliedDiscount))
	}
	fmt.Printf("Final payment amount: %.2f\n", sumTotalPrice*(100.0-appliedDiscount)/100.0)
	fmt.Printf("Payment method chosen: %s\n", payment_methods[payChoice-1])
	fmt.Println("=======================================================================================")
	if payChoice == 1 {
		query := "UPDATE Orders SET payment_method = ? WHERE user_id = ? AND order_date = ?;"
		_, err := db.Exec(query, "Credit Card", user_id, formattedTime)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Proceed with payment (y/n) ")
		pay, _ := reader.ReadString('\n')
		pay = strings.TrimSpace(pay)
		if pay == "y" {
			fmt.Println("Thank you for your purchase!")
			query := "UPDATE Orders SET payment_status = ? WHERE user_id = ? AND order_date = ?;"
			_, err := db.Exec(query, "Completed", user_id, formattedTime)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Println("Payment failed")
			query := "UPDATE Orders SET payment_status = ? WHERE user_id = ? AND order_date = ?;"
			_, err := db.Exec(query, "Failed", user_id, formattedTime)
			if err != nil {
				log.Fatal(err)
			}
		}
	} else if payChoice == 2 {
		query := "UPDATE Orders SET payment_method = ? WHERE user_id = ? AND order_date = ?;"
		_, err := db.Exec(query, "Virtual Account", user_id, formattedTime)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Proceed with payment (y/n) ")
		pay, _ := reader.ReadString('\n')
		pay = strings.TrimSpace(pay)
		if pay == "y" {
			fmt.Println("Thank you for your purchase!")
			query := "UPDATE Orders SET payment_status = ? WHERE user_id = ? AND order_date = ?;"
			_, err := db.Exec(query, "Completed", user_id, formattedTime)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Println("Payment failed")
			query := "UPDATE Orders SET payment_status = ? WHERE user_id = ? AND order_date = ?;"
			_, err := db.Exec(query, "Failed", user_id, formattedTime)
			if err != nil {
				log.Fatal(err)
			}
		}
	} else {
		fmt.Println("Thank you for your order")
		query := "UPDATE Orders SET payment_method = ? WHERE user_id = ? AND order_date = ?;"
		_, err := db.Exec(query, "Cash on Delivery", user_id, formattedTime)
		if err != nil {
			log.Fatal(err)
		}
	}
}
