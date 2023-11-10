-- name: CreateUser :exec
INSERT INTO users (id,name,balance) VALUES ($1,$2,$3);

-- name: ListUsers :many
SELECT * FROM users;

-- name: GetUserByName :one
SELECT * FROM users WHERE name = $1 LIMIT 1;

-- name: UpdateUserBalance :one
UPDATE users SET balance = $2  WHERE name = $1 RETURNING *;

-- name: CreateTransaction :exec
INSERT INTO transactions (id,type,value,date,product_name,seller_name) VALUES ($1,$2,$3,$4,$5,$6) RETURNING *;

-- name: ListTransactions :many
SELECT * FROM transactions WHERE id = $1 LIMIT 1;

-- name: GetTransactionById :one
SELECT * FROM transactions WHERE id = $1 LIMIT 1;

-- name: CreateProduct :exec
INSERT INTO products (id,title,producer_name,value) VALUES ($1,$2,$3,$4);

-- name: GetProductByName :one
SELECT * FROM products WHERE title = $1 LIMIT 1;

-- name: ListProducts :many
SELECT * FROM products;

-- name: DeleteProduct :exec
DELETE FROM products WHERE id = $1;