package entity

type Staff struct {
	ID       int
	Name     string
	Email    string
	Position string
}

// CREATE TABLE Products (
// 	product_id INT AUTO_INCREMENT PRIMARY KEY,
// 	name VARCHAR(255) NOT NULL,
// 	description TEXT,
// 	price DECIMAL(10, 2) NOT NULL,
// 	category_id INT,
// 	distributor_id INT,
// 	stock INT , 
// 	size VARCHAR(10) , 
// 	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
// 	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
// 	FOREIGN KEY (category_id) REFERENCES Categories(category_id),
// 	FOREIGN KEY (distributor_id) REFERENCES Distributors(distributor_id)
// );

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
