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

type Coupons struct {
	ID             int
	CouponCode     string
	DiscountAmount string
}

type Distributor struct {
	ID      int
	Name    string
	Address string
	Phone   string
}

type OrderDetail struct {
	ID          int
	OrderID     int
	ProductID   int
	ProductName string
	Quantity    int
	TotalPrice  float64
}

type Delivery struct {
	ID           int
	Name         string
	Size         string
	Cost         float64
	DeliveryDays int
}
