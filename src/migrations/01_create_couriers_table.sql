CREATE TABLE IF NOT EXISTS couriers (
    id INT auto_increment PRIMARY KEY,
    name TEXT NOT NULL,
    created BIGINT UNSIGNED,
    updated BIGINT UNSIGNED,
    busy TINYINT default 0
);