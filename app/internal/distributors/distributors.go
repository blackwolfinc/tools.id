package distributors

import (
	"database/sql"
	"fmt"
	"log"

	"app/config"
	"app/entity"
)

func ShowDistributor(cfg *config.Config) {
	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}

	catQuery := "SELECT distributor_id, name, address, phone_number  FROM Distributors"
	catRows, err := db.Query(catQuery)
	if err != nil {
		log.Fatal(err)
	}

	distributor := make([]*entity.Distributor, 0)
	for catRows.Next() {
		cat := new(entity.Distributor)
		err := catRows.Scan(&cat.ID, &cat.Name, &cat.Address, &cat.Phone)
		if err != nil {
			log.Fatal(err)
		}
		distributor = append(distributor, cat)
	}
	if err = catRows.Err(); err != nil {
		log.Fatal(err)
	}

	defer catRows.Close()
	defer db.Close()

	for _, c := range distributor {
		fmt.Printf(" %d. Nama Distributor :%s, Alamat :%s, Handpone:%s \n", c.ID, c.Name, c.Address, c.Phone)
	}

}

func AddDistributor(cfg *config.Config, name, address, phone string) {
	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO distributors (name, address, phone) VALUES (?, ?, ?)", name, address, phone)
	if err != nil {
		fmt.Println("Error adding distributor:", err)
	} else {
		fmt.Println("Distributor added successfully")
	}
}

func EditDistributor(cfg *config.Config, id int, name, address, phone string) {
	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	_, err = db.Exec("UPDATE distributors SET name=?, address=?, phone=? WHERE id=?", name, address, phone, id)
	if err != nil {
		fmt.Println("Error editing distributor:", err)
	} else {
		fmt.Println("Distributor edited successfully")
	}
}

func DeleteDistributor(cfg *config.Config, id int) {
	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM distributors WHERE id=?", id)
	if err != nil {
		fmt.Println("Error deleting distributor:", err)
	} else {
		fmt.Println("Distributor deleted successfully")
	}
}
