CREATE DATABASE IF NOT EXISTS orders;

USE orders;

CREATE TABLE IF NOT EXISTS orders
(
    id          VARCHAR(255)   NOT NULL,
    name        VARCHAR(255)   NOT NULL,
    price       DECIMAL(10, 2) NOT NULL,
    tax         DECIMAL(10, 2) NOT NULL,
    final_price DECIMAL(10, 2) NOT NULL,
    PRIMARY KEY (id)
);
