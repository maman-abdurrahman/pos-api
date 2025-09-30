-- =========================
-- SCHEMA
-- =========================

-- Tabel Category
CREATE TABLE categories (
    id BIGSERIAL PRIMARY KEY,
    category_code VARCHAR(50) UNIQUE NOT NULL,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabel PaymentMethod
CREATE TABLE payment_methods (
    id BIGSERIAL PRIMARY KEY,
    payment_method_code VARCHAR(50) UNIQUE NOT NULL,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabel Product
CREATE TABLE products (
    id BIGSERIAL PRIMARY KEY,
    product_code VARCHAR(50) UNIQUE NOT NULL,
    name VARCHAR(100) NOT NULL,
    category_code VARCHAR(50),
    price NUMERIC(12,2) NOT NULL,
    stock INT NOT NULL,
    date_in TIMESTAMP,
    created_by VARCHAR(100),
    updated_by VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    CONSTRAINT fk_product_category FOREIGN KEY (category_code) REFERENCES categories(category_code)
);

-- Tabel Role
CREATE TABLE roles (
    id BIGSERIAL PRIMARY KEY,
    role_code VARCHAR(50) UNIQUE NOT NULL,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabel Users
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    user_code VARCHAR(50) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role_code VARCHAR(50) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user_role FOREIGN KEY (role_code) REFERENCES roles(role_code)
);

-- Tabel Sales
CREATE TABLE sales (
    id BIGSERIAL PRIMARY KEY,
    sales_code VARCHAR(50) UNIQUE NOT NULL,
    invoice_number VARCHAR(100) UNIQUE NOT NULL,
    user_code VARCHAR(50) NOT NULL,
    payment_method_code VARCHAR(50) NOT NULL,
    total NUMERIC(12,2) NOT NULL,
    discount NUMERIC(12,2) DEFAULT 0,
    final_total NUMERIC(12,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_sales_user FOREIGN KEY (user_code) REFERENCES users(user_code),
    CONSTRAINT fk_sales_payment FOREIGN KEY (payment_method_code) REFERENCES payment_methods(payment_method_code)
);

-- Tabel SaleItem
CREATE TABLE sale_items (
    id BIGSERIAL PRIMARY KEY,
    sale_code VARCHAR(50) NOT NULL,
    product_code VARCHAR(50) NOT NULL,
    quantity INT NOT NULL,
    unit_price NUMERIC(12,2) NOT NULL,
    subtotal NUMERIC(12,2) NOT NULL,
    CONSTRAINT fk_saleitem_sales FOREIGN KEY (sale_code) REFERENCES sales(sales_code),
    CONSTRAINT fk_saleitem_product FOREIGN KEY (product_code) REFERENCES products(product_code)
);

-- =========================
-- SEEDER DATA
-- =========================

-- Roles
INSERT INTO roles (role_code, name) VALUES
('ADMIN', 'Administrator'),
('CASHIER', 'Kasir');

-- Users
INSERT INTO users (user_code, name, email, password, role_code, is_active)
VALUES
('USR001', 'Admin Utama', 'admin@example.com', 'hashed_password_admin', 'ADMIN', TRUE),
('USR002', 'Kasir Toko', 'kasir@example.com', 'hashed_password_kasir', 'CASHIER', TRUE);

-- Categories
INSERT INTO categories (category_code, name) VALUES
('CAT001', 'Makanan'),
('CAT002', 'Minuman'),
('CAT003', 'ATK');

-- Payment Methods
INSERT INTO payment_methods (payment_method_code, name, description, is_active)
VALUES
('CASH', 'Cash', 'Pembayaran tunai', TRUE),
('TRANSFER', 'Transfer Bank', 'Pembayaran via transfer bank', TRUE),
('QRIS', 'QRIS', 'Pembayaran via QRIS', TRUE);

-- Products
INSERT INTO products (product_code, name, category_code, price, stock, created_by)
VALUES
('PRD001', 'Indomie Goreng', 'CAT001', 3000, 100, 'USR001'),
('PRD002', 'Aqua Botol 600ml', 'CAT002', 4000, 200, 'USR001'),
('PRD003', 'Buku Tulis Sidu', 'CAT003', 5000, 50, 'USR001');

-- Sales
INSERT INTO sales (sales_code, invoice_number, user_code, payment_method_code, total, discount, final_total)
VALUES
('SLS001', 'INV-20250930-001', 'USR002', 'CASH', 10000, 0, 10000),
('SLS002', 'INV-20250930-002', 'USR002', 'QRIS', 8000, 500, 7500);

-- Sale Items
INSERT INTO sale_items (sale_code, product_code, quantity, unit_price, subtotal)
VALUES
('SLS001', 'PRD001', 2, 3000, 6000),
('SLS001', 'PRD002', 1, 4000, 4000),
('SLS002', 'PRD003', 1, 5000, 5000),
('SLS002', 'PRD002', 1, 4000, 3000);

