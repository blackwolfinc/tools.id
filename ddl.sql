-- B2C from Distributor to Customer

-- ===================================================================================================
-- Bundle 1 bagas
-- ===================================================================================================
-- Login Page & sign up

-- admin - fitur CRUD 
-- master Distributors (full)
-- master Delivery
-- master Coupons
-- master Categories

-- Distributor
-- master Products 
-- master Distributors (update)
-- input produk yang akan dijual, berdasarkan kategori yang sudah dibuat oleh admin

-- Customer
-- Belanja
-- Login
-- pilih kategori
-- pilih produk
-- pilih delivery
-- pilih pembayaran
-- input kupon (optional)
-- bayar
-- receipt & shipping estimate


-- User Authentication and Authorization (ACC)
CREATE TABLE Users (
    user_id INT AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    address TEXT NOT NULL,
    role ENUM('Admin', 'Customer', 'Distributor') NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Delivery (ACC)
-- estimated_date 1 = 1 hari
CREATE TABLE Delivery (
    delivery_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    size ENUM('S', 'M', 'L') NOT NULL,
    cost DECIMAL(10, 2) NOT NULL,
    delivery_days INT NOT NULL
);

-- Discount/Coupon Management (ACC)
-- diskon persen 
CREATE TABLE Coupons (
    coupon_id INT AUTO_INCREMENT PRIMARY KEY,
    coupon_code VARCHAR(255) UNIQUE NOT NULL,
    discount_amount DECIMAL(10, 2) NOT NULL
);

-- ===================================================================================================
-- Bundle 2 kamal
-- ===================================================================================================

-- Product Management (ACC)
-- size small | medium | large
CREATE TABLE Products (
    product_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL,
    category_id INT,
    distributor_id INT,
    stock INT, 
    size ENUM('S', 'M', 'L') NOT NULL, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (category_id) REFERENCES Categories(category_id),
    FOREIGN KEY (distributor_id) REFERENCES Distributors(distributor_id)
);

CREATE TABLE Distributors (
    distributor_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    address TEXT,
    phone_number VARCHAR(20)
);

-- Categories (ACC)
CREATE TABLE Categories (
    category_id INT AUTO_INCREMENT PRIMARY KEY,
    category_name VARCHAR(255) NOT NULL,
    description TEXT
);

-- ===================================================================================================
-- Bundle 3 marcel 
-- ===================================================================================================

-- Order Management (Transaksi) (ACC)

-- s 2 = medium
-- s 3 = large
-- m 2 = large

-- bisa mempunyai banyak produk dalam sekali beli 
-- maksimal 3 item yang berbeda
-- total_amount = total dari harga order detail per produk
CREATE TABLE Orders (
    order_id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT, 
    delivery_id INT,
    coupon_id INT,
    order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status ENUM('Processing', 'Shipped', 'Delivered', 'Cancelled') NOT NULL DEFAULT 'Processing',
    payment_method ENUM('Credit Card', 'Cash', 'Virtual Account'),
    payment_status ENUM('Pending', 'Completed', 'Failed') NOT NULL DEFAULT 'Pending',
    total_amount DECIMAL(10, 2) NOT NULL DEFAULT 0,
    FOREIGN KEY (user_id) REFERENCES Users(user_id),
    FOREIGN KEY (delivery_id) REFERENCES Delivery(delivery_id),
    FOREIGN KEY (coupon_id) REFERENCES Coupons(coupon_id)
);

-- Detail Order (Transaksi) (ACC)
-- total_price = total dari harga produk dikali qty
CREATE TABLE OrderDetails (
    order_detail_id INT AUTO_INCREMENT PRIMARY KEY,
    order_id INT,
    product_id INT,
    quantity INT NOT NULL,
    total_price DECIMAL(10, 2) NOT NULL DEFAULT 0,
    FOREIGN KEY (order_id) REFERENCES Orders(order_id),
    FOREIGN KEY (product_id) REFERENCES Products(product_id)
);


-- -- Payment Integration (ACC)
-- CREATE TABLE Payments (
--     payment_id INT AUTO_INCREMENT PRIMARY KEY,
--     order_id INT,
--     
--     transaction_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     amount DECIMAL(10, 2) NOT NULL,
--     FOREIGN KEY (order_id) REFERENCES Orders(order_id)
-- );
