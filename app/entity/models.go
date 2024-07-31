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
	ID            int
	Name          string
	Description   string
	Price         float64
	DistributorID int
	CategoryID    int
	Stock         int
	Size          string
}

type Category struct {
	ID          int
	Name        string
	Description string
}

type Distributor struct {
	ID      int
	Name    string
	Address string
	Phone   string
}

type Order struct {
	ID      int
	Name    string
	Address string
	Phone   string
}
