package entity

import "time"

type User struct {
	ID        int
	Email     string
	Password  string
	Address   string
	Role      string
	CreatedAt time.Time
}

type Product struct {
	Name string
	Description string
	Price float64
	Distributor string
	Stock int
	Size string
}

type Category struct {
	Name string
	Description string
}
