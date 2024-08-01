package delivery

import (
	"database/sql"
	"fmt"
	"log"

	"app/config"
	"app/entity"
)

func ShowDelivery(cfg *config.Config) {
	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}

	catQuery := "SELECT delivery_id, name, size, cost  , delivery_days FROM Delivery"
	catRows, err := db.Query(catQuery)
	if err != nil {
		log.Fatal(err)
	}

	delivery := make([]*entity.Delivery, 0)
	for catRows.Next() {
		cat := new(entity.Delivery)
		err := catRows.Scan(&cat.ID, &cat.Name, &cat.Size, &cat.Cost, &cat.DeliveryDays)
		if err != nil {
			log.Fatal(err)
		}
		delivery = append(delivery, cat)
	}
	if err = catRows.Err(); err != nil {
		log.Fatal(err)
	}

	defer catRows.Close()
	defer db.Close()

	for _, c := range delivery {
		fmt.Printf(" %d. Kurir :%s, Ukuran:%s, Harga:%v, Estimasi Hari:%d \n", c.ID, c.Name, c.Size, c.Cost, c.DeliveryDays)
	}

}

func AddDelivery(cfg *config.Config, name, size, cost, estimated_date string) {
	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO Delivery (name,size, cost, estimated_date ) VALUES (?,?, ?, ?)", name, size, cost, estimated_date)
	if err != nil {
		fmt.Println("Error adding distributor:", err)
	} else {
		fmt.Println("Distributor added successfully")
	}
}

func EditDelivery(cfg *config.Config, id int, size, name, cost, estimated_date string) {
	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	_, err = db.Exec("UPDATE Delivery SET name=?,  size=?, cost=?, estimated_date=? WHERE id=?", name, size, cost, estimated_date, id)
	if err != nil {
		fmt.Println("Error editing distributor:", err)
	} else {
		fmt.Println("Distributor edited successfully")
	}
}

func DeleteDelivery(cfg *config.Config, id int) {
	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM Delivery WHERE id=?", id)
	if err != nil {
		fmt.Println("Error deleting distributor:", err)
	} else {
		fmt.Println("Distributor deleted successfully")
	}
}
