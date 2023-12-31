// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: query.sql

package db

import (
	"context"
)

const createProduct = `-- name: CreateProduct :exec
INSERT INTO products (id,title,producer_name,value) VALUES ($1,$2,$3,$4)
`

type CreateProductParams struct {
	ID           string
	Title        string
	ProducerName string
	Value        int32
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) error {
	_, err := q.db.ExecContext(ctx, createProduct,
		arg.ID,
		arg.Title,
		arg.ProducerName,
		arg.Value,
	)
	return err
}

const createTransaction = `-- name: CreateTransaction :exec
INSERT INTO transactions (id,type,value,date,product_name,seller_name) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id, type, date, product_name, seller_name, value, created_at
`

type CreateTransactionParams struct {
	ID          string
	Type        int32
	Value       int32
	Date        string
	ProductName string
	SellerName  string
}

func (q *Queries) CreateTransaction(ctx context.Context, arg CreateTransactionParams) error {
	_, err := q.db.ExecContext(ctx, createTransaction,
		arg.ID,
		arg.Type,
		arg.Value,
		arg.Date,
		arg.ProductName,
		arg.SellerName,
	)
	return err
}

const createUser = `-- name: CreateUser :exec
INSERT INTO users (id,name,balance) VALUES ($1,$2,$3)
`

type CreateUserParams struct {
	ID      string
	Name    string
	Balance int32
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.ExecContext(ctx, createUser, arg.ID, arg.Name, arg.Balance)
	return err
}

const deleteProduct = `-- name: DeleteProduct :exec
DELETE FROM products WHERE id = $1
`

func (q *Queries) DeleteProduct(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteProduct, id)
	return err
}

const deleteTransaction = `-- name: DeleteTransaction :exec
DELETE FROM transactions WHERE id = $1
`

func (q *Queries) DeleteTransaction(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteTransaction, id)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getProductByName = `-- name: GetProductByName :one
SELECT id, title, producer_name, value FROM products WHERE title = $1 LIMIT 1
`

func (q *Queries) GetProductByName(ctx context.Context, title string) (Product, error) {
	row := q.db.QueryRowContext(ctx, getProductByName, title)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.ProducerName,
		&i.Value,
	)
	return i, err
}

const getTransactionById = `-- name: GetTransactionById :one
SELECT id, type, date, product_name, seller_name, value, created_at FROM transactions WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTransactionById(ctx context.Context, id string) (Transaction, error) {
	row := q.db.QueryRowContext(ctx, getTransactionById, id)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.Date,
		&i.ProductName,
		&i.SellerName,
		&i.Value,
		&i.CreatedAt,
	)
	return i, err
}

const getUserByName = `-- name: GetUserByName :one
SELECT id, name, balance, created_at FROM users WHERE name = $1 LIMIT 1
`

func (q *Queries) GetUserByName(ctx context.Context, name string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByName, name)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Balance,
		&i.CreatedAt,
	)
	return i, err
}

const listProducts = `-- name: ListProducts :many
SELECT id, title, producer_name, value FROM products
`

func (q *Queries) ListProducts(ctx context.Context) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, listProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.ProducerName,
			&i.Value,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTransactions = `-- name: ListTransactions :many
SELECT id, type, date, product_name, seller_name, value, created_at FROM transactions
`

func (q *Queries) ListTransactions(ctx context.Context) ([]Transaction, error) {
	rows, err := q.db.QueryContext(ctx, listTransactions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transaction
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.ID,
			&i.Type,
			&i.Date,
			&i.ProductName,
			&i.SellerName,
			&i.Value,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUsers = `-- name: ListUsers :many
SELECT id, name, balance, created_at FROM users
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Balance,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUserBalance = `-- name: UpdateUserBalance :one
UPDATE users SET balance = $2  WHERE name = $1 RETURNING id, name, balance, created_at
`

type UpdateUserBalanceParams struct {
	Name    string
	Balance int32
}

func (q *Queries) UpdateUserBalance(ctx context.Context, arg UpdateUserBalanceParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUserBalance, arg.Name, arg.Balance)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Balance,
		&i.CreatedAt,
	)
	return i, err
}
