CREATE DATABASE alta_online_shop;

USE alta_online_shop;

CREATE TABLE user (
    user_id INT PRIMARY KEY AUTO_INCREMENT,
    nama VARCHAR(255)
);

CREATE TABLE product (
    product_id INT PRIMARY KEY AUTO_INCREMENT,
    product_name VARCHAR(255)
);

CREATE TABLE product_type (
    type_id INT PRIMARY KEY AUTO_INCREMENT,
    product_id INT,
    type_name VARCHAR(255),
    FOREIGN KEY (product_id) REFERENCES product(product_id)
);

CREATE TABLE product_description (
    description_id INT PRIMARY KEY AUTO_INCREMENT,
    product_id INT,
    description_text TEXT,
    FOREIGN KEY (product_id) REFERENCES product(product_id)
);

CREATE TABLE payment_method (
    method_id INT PRIMARY KEY AUTO_INCREMENT,
    method_name VARCHAR(255)
);

CREATE TABLE transaction (
    transaction_id INT PRIMARY KEY AUTO_INCREMENT
);

CREATE TABLE transaction_detail (
    detail_id INT PRIMARY KEY AUTO_INCREMENT,
    transaction_id INT,
    product_id INT,
    quantity INT,
    FOREIGN KEY (transaction_id) REFERENCES transaction(transaction_id),
    FOREIGN KEY (product_id) REFERENCES product(product_id)
);

CREATE TABLE kurir (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

ALTER TABLE kurir ADD COLUMN ongkos_dasar DECIMAL(10, 2);

RENAME TABLE kurir TO shipping;

ALTER TABLE payment_method
ADD COLUMN description_id INT,
ADD CONSTRAINT fk_payment_description
FOREIGN KEY (description_id) REFERENCES product_description(description_id);

CREATE TABLE address (
    address_id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT,
    address_text VARCHAR(255),
    FOREIGN KEY (user_id) REFERENCES user(user_id)
);

CREATE TABLE user_payment_method_detail (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT,
    method_id INT,
    FOREIGN KEY (user_id) REFERENCES user(user_id),
    FOREIGN KEY (method_id) REFERENCES payment_method(method_id)
);

