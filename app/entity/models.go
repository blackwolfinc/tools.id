package entity

type Staff struct {
	ID       int
	Name     string
	Email    string
	Position string
}

type Product struct {
	ID            int
	Name          string
	Description   string
	Price         float64
	CategoryID    int
	DistributorID int
}

type Category struct {
	ID   int
	Name string
}

type Distributor struct {
	ID      int
	Name    string
	Address string
	Phone   string
}
