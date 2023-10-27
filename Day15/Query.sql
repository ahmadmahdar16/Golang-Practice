CREATE DATABASE alta_online_shop;

USE alta_online_shop;

CREATE TABLE user (
    user_id INT PRIMARY KEY AUTO_INCREMENT,
    nama VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
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

/*------------------------INSERT-------------------------------------------------*/ 

INSERT INTO product_type (type_name) VALUES ('HP'), ('Baju'), ('Buku');


INSERT INTO product (product_name) VALUES ('Xiaomi', 'Iphone');
INSERT INTO product_description (product_id, description_text) VALUES (1, 'Gaming Phone'), (2, 'HP mahal');


INSERT INTO product (product_name) VALUES ('Bilabong'), ('Crocodile'), ('Nevada');
INSERT INTO product_description (product_id, description_text) VALUES (3, 'Baju jadul'), (4, 'Baju Bapak'), (5, 'Baju anak');


INSERT INTO product (product_name) VALUES ('Buku 1'), ('Buku 2'), ('Buku 3');
INSERT INTO product_description (product_id, description_text) VALUES (6, 'Buku bestseller'), (7, 'Buku Novel'), (8, 'Buku Pelajaran');


INSERT INTO payment_method (method_name) VALUES ('Qris'), ('Transfer Bank'), ('Gopay/Ovo');


INSERT INTO user (nama) VALUES ('Asep'), ('Udin'), ('Jamal'), ('Soleh'), ('Mahmud');

ALTER TABLE transaction ADD COLUMN user_id INT;

ALTER TABLE transaction ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES user(user_id);

INSERT INTO transaction (user_id) VALUES 
(1), (1), (1), 
(2), (2), (2), 
(3), (3), (3), 
(4), (4), (4), 
(5), (5), (5);

INSERT INTO transaction_detail (transaction_id, product_id, quantity) VALUES

(1, 1, 2), (1, 3, 1), (1, 5, 3),

(4, 2, 1), (4, 4, 2), (4, 6, 1),

(7, 3, 1), (7, 5, 2), (7, 7, 1),

(10, 1, 3), (10, 3, 1), (10, 5, 2),

(13, 2, 2), (13, 4, 1), (13, 6, 3);

ALTER TABLE user ADD COLUMN gender CHAR(1);

/*------------------------SELECT-------------------------------------------------*/ 
SELECT nama
FROM user
WHERE gender = 'M';

SELECT *
FROM product
WHERE product_id = 3;

SELECT *
FROM user
WHERE created_at >= CURDATE() - INTERVAL 7 DAY
AND nama LIKE '%a%';

SELECT COUNT(*)
FROM user
WHERE gender = 'F';

SELECT *
FROM user
ORDER BY nama;

SELECT *
FROM transaction_detail
WHERE product_id = 3
LIMIT 5;

/*------------------------UPDATE-------------------------------------------------*/ 

UPDATE product
SET product_name = 'product dummy'
WHERE product_id = 1;

UPDATE transaction_detail
SET quantity = 3
WHERE product_id = 1;

/*------------------------DELETE-------------------------------------------------*/ 

DELETE FROM transaction_detail
WHERE product_id = 1;

DELETE FROM product_description
WHERE product_id = 1;

DELETE FROM product
WHERE product_id = 1;