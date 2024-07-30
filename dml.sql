-- Distributors (ACC)
INSERT INTO Distributors (name, address, phone_number) VALUES
('PT United Tractors Tbk', 'Jl. Raya Bekasi Km. 22, Cakung, Jakarta Timur, DKI Jakarta', '021-24579845'),
('PT Trakindo Utama', 'Jl. Raya Cakung Cilincing Kav. 3-4, Jakarta Utara, DKI Jakarta', '021-24579844'),
('PT Krisbow Indonesia', 'Jl. Sultan Iskandar Muda No. 29, Jakarta Selatan, DKI Jakarta', '021-7221234'),
('PT Wirya Krenindo Perkasa', 'Jl. Mangga Dua Raya No. 15, Jakarta Pusat, DKI Jakarta', '021-6123456'),
('PT Semen Indonesia Tbk', 'Jl. Veteran, Gresik, Jawa Timur', '031-3981731'),
('PT Indocement Tunggal Prakarsa Tbk', 'Jl. Jend. Sudirman Kav. 70-71, Jakarta Selatan, DKI Jakarta', '021-2512121'),
('PT Chandra Asri Petrochemical Tbk', 'Jl. Jend. Sudirman Kav. 10-11, Jakarta Selatan, DKI Jakarta', '021-5302950'),
('PT Pupuk Indonesia (Persero)', 'Jl. Taman Anggrek Kemanggisan Jaya, Jakarta Barat, DKI Jakarta', '021-53666745'),
('PT Schneider Electric Manufacturing Batam', 'Jl. Hang Kesturi I No. 2, Batam, Kepulauan Riau', '0778-482200'),
('PT Siemens Indonesia', 'Jl. Jend. Gatot Subroto Kav. 27, Jakarta Selatan, DKI Jakarta', '021-2754444'),
('PT Astra Otoparts Tbk', 'Jl. Raya Pegangsaan Dua Km. 2,2, Jakarta Utara, DKI Jakarta', '021-4605555'),
('PT Kawan Lama Sejahtera', 'Jl. Puri Kencana No. 1, Kembangan, Jakarta Barat, DKI Jakarta', '021-5828282'),
('PT Mayora Indah Tbk', 'Jl. Tomang Raya No. 21-23, Jakarta Barat, DKI Jakarta', '021-80638888'),
('PT Indofood CBP Sukses Makmur Tbk', 'Jl. Jend. Sudirman Kav. 76-78, Jakarta Selatan, DKI Jakarta', '021-57959810'),
('PT ABB Sakti Industri', 'Jl. Cilandak KKO No. 1, Jakarta Selatan, DKI Jakarta', '021-7801688'),
('PT Omron Manufacturing of Indonesia', 'Jl. Raya Jakarta-Bogor Km. 47, Cibinong, Jawa Barat', '021-87922323');


-- Categories (ACC)
INSERT INTO Categories (category_name, description) VALUES
('Machinery', 'Machines and equipment used in manufacturing processes.'),
('Tools', 'Hand tools and power tools used in various applications.'),
('Materials', 'Raw materials and processed materials for manufacturing.'),
('Chemicals', 'Industrial chemicals and substances used in manufacturing.'),
('Electrical Components', 'Components and parts used in electrical systems.'),
('Industrial Supplies', 'Supplies and equipment used in industrial settings.'),
('Packaging and Storage', 'Materials and solutions for packaging and storing products.'),
('Automation and Robotics', 'Automated systems and robotic equipment for manufacturing.');


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
('S', 100.00, 1),
('M', 150.00, 2),
('L', 200.00, 3),
('S', 250.00, 1),
('M', 300.00, 2),
('L', 350.00, 3),
('S', 400.00, 1),
('M', 450.00, 2),
('L', 500.00, 3),
('S', 550.00, 1);

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

-- Products (ACC)
-- Assuming Categories and Distributors tables already have records with category_id from 1 to 3 and distributor_id from 1 to 3
INSERT INTO Products (name, description, price, category_id, distributor_id, stock, size) VALUES
('Excavator', 'Heavy-duty machinery used for construction and excavation.', 1000000.00, 1, 1, 10, 'L'),
('Bulldozer', 'Powerful machine used for pushing large quantities of soil, sand, or other material.', 850000.00, 1, 2, 5, 'L'),
('Drill Machine', 'Portable drilling machine used for various applications.', 1500.00, 2, 3, 100, 'M'),
('Wrench Set', 'Set of wrenches used for tightening or loosening nuts and bolts.', 500.00, 2, 4, 200, 'S'),
('Cement', 'Building material used for construction.', 100.00, 3, 5, 1000, 'L'),
('Concrete Mix', 'Mixture of cement, aggregates, and water used in construction.', 200.00, 3, 6, 500, 'L'),
('Polyethylene', 'Type of plastic used for various industrial applications.', 300.00, 4, 7, 800, 'M'),
('Fertilizer', 'Substance used to enhance the growth of plants and crops.', 50.00, 4, 8, 1500, 'M'),
('Circuit Breaker', 'Electrical device used to protect an electrical circuit from damage caused by overload or short circuit.', 100.00, 5, 9, 300, 'S'),
('Transformer', 'Electrical device used to transfer electrical energy between two or more circuits.', 1000.00, 5, 10, 50, 'L'),
('Bearing', 'Machine element used to constrain relative motion to only the desired motion.', 200.00, 6, 11, 500, 'S'),
('Hydraulic Pump', 'Mechanical source of power that converts mechanical power into hydraulic energy.', 1500.00, 6, 12, 100, 'M'),
('Packaging Tape', 'Tape used for sealing boxes and packages.', 10.00, 7, 13, 1000, 'S'),
('Storage Box', 'Container used for storing various items.', 50.00, 7, 14, 500, 'M'),
('Industrial Robot', 'Robot system used for manufacturing.', 50000.00, 8, 15, 20, 'L'),
('PLC Controller', 'Programmable Logic Controller used for industrial automation.', 2000.00, 8, 16, 100, 'M');

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
