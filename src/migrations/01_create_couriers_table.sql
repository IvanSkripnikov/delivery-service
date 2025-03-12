CREATE TABLE IF NOT EXISTS couriers (
    id INT auto_increment PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created BIGINT UNSIGNED,
    updated BIGINT UNSIGNED,
    busy TINYINT DEFAULT 0
);