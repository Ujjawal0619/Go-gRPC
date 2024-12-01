CRATE TABLE IF NOT EXISTS orders (
    id CHAR(30) PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    account_id CHAR(30) NOT NULL,
    total_price MONEY NOT NULL
);

CREATE TABLE IF NOT EXISTS order_products (
    order_id CHAR(30) REFERENCES orders (id) ON DELETE CASCADE,
    product_id CHAR(30),
    quantity INT NOT NULL,
    PRIMARY KEY (product_id, order_id)
)