-- User Authentication and Authorization (ACC)
INSERT INTO Users (email, password_hash, address, role) VALUES
('user1@example.com', 'hash1', 'Address 1', 'Admin'),
('user2@example.com', 'hash2', 'Address 2', 'Customer'),
('user3@example.com', 'hash3', 'Address 3', 'Distributor'),
('user4@example.com', 'hash4', 'Address 4', 'Admin'),
('user5@example.com', 'hash5', 'Address 5', 'Customer'),
('user6@example.com', 'hash6', 'Address 6', 'Distributor'),
('user7@example.com', 'hash7', 'Address 7', 'Admin'),
('user8@example.com', 'hash8', 'Address 8', 'Customer'),
('user9@example.com', 'hash9', 'Address 9', 'Distributor'),
('user10@example.com', 'hash10', 'Address 10', 'Admin');

-- Delivery (ACC)
INSERT INTO Delivery (size, cost, estimated_date) VALUES
(1, 100.00, 1),
(2, 150.00, 2),
(3, 200.00, 3),
(4, 250.00, 1),
(5, 300.00, 2),
(6, 350.00, 3),
(7, 400.00, 1),
(8, 450.00, 2),
(9, 500.00, 3),
(10, 550.00, 1);

-- Payment Integration (ACC)
-- Assuming Orders table already has 10 records with order_id from 1 to 10
INSERT INTO Payments (order_id, payment_method, amount) VALUES
(1, 'Credit Card', 100.00),
(2, 'Cash', 150.00),
(3, 'Virtual Account', 200.00),
(4, 'Credit Card', 250.00),
(5, 'Cash', 300.00),
(6, 'Virtual Account', 350.00),
(7, 'Credit Card', 400.00),
(8, 'Cash', 450.00),
(9, 'Virtual Account', 500.00),
(10, 'Credit Card', 550.00);

-- Discount/Coupon Management (ACC)
INSERT INTO Coupons (coupon_code, discount_amount) VALUES
('COUPON1', 10.00),
('COUPON2', 15.00),
('COUPON3', 20.00),
('COUPON4', 25.00),
('COUPON5', 30.00),
('COUPON6', 35.00),
('COUPON7', 40.00),
('COUPON8', 45.00),
('COUPON9', 50.00),
('COUPON10', 55.00);

-- Product Management (ACC)
-- Assuming Categories and Distributors tables already have records with category_id from 1 to 3 and distributor_id from 1 to 3
INSERT INTO Products (name, description, price, category_id, distributor_id, stock, size) VALUES
('Product1', 'Description 1', 100.00, 1, 1, 10, 'Small'),
('Product2', 'Description 2', 150.00, 2, 2, 20, 'Medium'),
('Product3', 'Description 3', 200.00, 3, 3, 30, 'Large'),
('Product4', 'Description 4', 250.00, 1, 2, 40, 'Small'),
('Product5', 'Description 5', 300.00, 2, 3, 50, 'Medium'),
('Product6', 'Description 6', 350.00, 3, 1, 60, 'Large'),
('Product7', 'Description 7', 400.00, 1, 3, 70, 'Small'),
('Product8', 'Description 8', 450.00, 2, 1, 80, 'Medium'),
('Product9', 'Description 9', 500.00, 3, 2, 90, 'Large'),
('Product10', 'Description 10', 550.00, 1, 2, 100, 'Small');

-- Distributors (ACC)
INSERT INTO Distributors (name, address, phone_number) VALUES
('Distributor1', 'Address 1', '1234567890'),
('Distributor2', 'Address 2', '1234567891'),
('Distributor3', 'Address 3', '1234567892'),
('Distributor4', 'Address 4', '1234567893'),
('Distributor5', 'Address 5', '1234567894'),
('Distributor6', 'Address 6', '1234567895'),
('Distributor7', 'Address 7', '1234567896'),
('Distributor8', 'Address 8', '1234567897'),
('Distributor9', 'Address 9', '1234567898'),
('Distributor10', 'Address 10', '1234567899');

-- Categories (ACC)
INSERT INTO Categories (category_name, description) VALUES
('Category1', 'Description 1'),
('Category2', 'Description 2'),
('Category3', 'Description 3'),
('Category4', 'Description 4'),
('Category5', 'Description 5'),
('Category6', 'Description 6'),
('Category7', 'Description 7'),
('Category8', 'Description 8'),
('Category9', 'Description 9'),
('Category10', 'Description 10');

-- Order Management (Transaksi) (ACC)
-- Assuming Users, Delivery, Coupons tables already have records
INSERT INTO Orders (user_id, delivery_id, coupon_id, status, payment_status, total_amount) VALUES
(1, 1, 1, 'Processing', 'Pending', 100.00),
(2, 2, 2, 'Shipped', 'Completed', 150.00),
(3, 3, 3, 'Delivered', 'Failed', 200.00),
(4, 4, 4, 'Cancelled', 'Pending', 250.00),
(5, 5, 5, 'Processing', 'Completed', 300.00),
(6, 6, 6, 'Shipped', 'Failed', 350.00),
(7, 7, 7, 'Delivered', 'Pending', 400.00),
(8, 8, 8, 'Cancelled', 'Completed', 450.00),
(9, 9, 9, 'Processing', 'Failed', 500.00),
(10, 10, 10, 'Shipped', 'Pending', 550.00);

-- Detail Order (Transaksi) (ACC)
-- Assuming Products and Orders tables already have records
INSERT INTO OrderDetails (order_id, product_id, quantity, total_price) VALUES
(1, 1, 1, 100.00),
(2, 2, 2, 150.00),
(3, 3, 3, 200.00),
(4, 4, 4, 250.00),
(5, 5, 5, 300.00),
(6, 6, 6, 350.00),
(7, 7, 7, 400.00),
(8, 8, 8, 450.00),
(9, 9, 9, 500.00),
(10, 10, 10, 550.00);
