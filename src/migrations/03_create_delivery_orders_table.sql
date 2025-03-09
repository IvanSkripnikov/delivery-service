CREATE TABLE IF NOT EXISTS delivery_orders (
    id INT auto_increment PRIMARY KEY,
    order_id INT NOT NULL,
    courier_id INT NOT NULL,
    created BIGINT UNSIGNED,
    completed BIGINT UNSIGNED,
    status TINYINT DEFAULT 0
);