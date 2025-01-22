// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package schema

import (
	"context"
	"database/sql"
)

const checkUserExists = `-- name: CheckUserExists :one
SELECT id FROM account WHERE username = ?
`

func (q *Queries) CheckUserExists(ctx context.Context, username string) (string, error) {
	row := q.db.QueryRowContext(ctx, checkUserExists, username)
	var id string
	err := row.Scan(&id)
	return id, err
}

const createAccount = `-- name: CreateAccount :exec
INSERT INTO account (id, first_name, last_name, username, email, password, balance)
VALUES (?, ?, ?, ?, ?, ?, ?)
`

type CreateAccountParams struct {
	ID        string         `json:"id"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Username  string         `json:"username"`
	Email     sql.NullString `json:"email"`
	Password  string         `json:"password"`
	Balance   float64        `json:"balance"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) error {
	_, err := q.db.ExecContext(ctx, createAccount,
		arg.ID,
		arg.FirstName,
		arg.LastName,
		arg.Username,
		arg.Email,
		arg.Password,
		arg.Balance,
	)
	return err
}

const deleteAccount = `-- name: DeleteAccount :exec
DELETE FROM account
WHERE id = ?
`

func (q *Queries) DeleteAccount(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteAccount, id)
	return err
}

const deleteHistory = `-- name: DeleteHistory :exec
DELETE FROM history
WHERE (sender = ? OR receiver = ?)
  AND NOT EXISTS (SELECT 1 FROM account WHERE username = history.sender)
  AND NOT EXISTS (SELECT 1 FROM account WHERE username = history.receiver)
`

type DeleteHistoryParams struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
}

func (q *Queries) DeleteHistory(ctx context.Context, arg DeleteHistoryParams) error {
	_, err := q.db.ExecContext(ctx, deleteHistory, arg.Sender, arg.Receiver)
	return err
}

const deposit = `-- name: Deposit :exec
UPDATE account
SET balance = balance + ?
WHERE id = ? OR username = ?
`

type DepositParams struct {
	Balance  float64 `json:"balance"`
	ID       string  `json:"id"`
	Username string  `json:"username"`
}

func (q *Queries) Deposit(ctx context.Context, arg DepositParams) error {
	_, err := q.db.ExecContext(ctx, deposit, arg.Balance, arg.ID, arg.Username)
	return err
}

const getAccountByUsername = `-- name: GetAccountByUsername :one
SELECT id, first_name, last_name, username, email, balance
FROM account
WHERE username = ? AND password = ?
`

type GetAccountByUsernameParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type GetAccountByUsernameRow struct {
	ID        string         `json:"id"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Username  string         `json:"username"`
	Email     sql.NullString `json:"email"`
	Balance   float64        `json:"balance"`
}

func (q *Queries) GetAccountByUsername(ctx context.Context, arg GetAccountByUsernameParams) (GetAccountByUsernameRow, error) {
	row := q.db.QueryRowContext(ctx, getAccountByUsername, arg.Username, arg.Password)
	var i GetAccountByUsernameRow
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Username,
		&i.Email,
		&i.Balance,
	)
	return i, err
}

const getTransactions = `-- name: GetTransactions :many
SELECT id, sender, receiver, amount, timestamp
FROM history
WHERE sender = ? OR receiver = ?
ORDER BY timestamp DESC
`

type GetTransactionsParams struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
}

func (q *Queries) GetTransactions(ctx context.Context, arg GetTransactionsParams) ([]History, error) {
	rows, err := q.db.QueryContext(ctx, getTransactions, arg.Sender, arg.Receiver)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []History
	for rows.Next() {
		var i History
		if err := rows.Scan(
			&i.ID,
			&i.Sender,
			&i.Receiver,
			&i.Amount,
			&i.Timestamp,
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

const insertTransaction = `-- name: InsertTransaction :exec
INSERT INTO history (id, sender, receiver, amount, timestamp)
VALUES (?, ?, ?, ?, ?)
`

type InsertTransactionParams struct {
	ID        string  `json:"id"`
	Sender    string  `json:"sender"`
	Receiver  string  `json:"receiver"`
	Amount    float64 `json:"amount"`
	Timestamp string  `json:"timestamp"`
}

func (q *Queries) InsertTransaction(ctx context.Context, arg InsertTransactionParams) error {
	_, err := q.db.ExecContext(ctx, insertTransaction,
		arg.ID,
		arg.Sender,
		arg.Receiver,
		arg.Amount,
		arg.Timestamp,
	)
	return err
}

const withdraw = `-- name: Withdraw :exec
UPDATE account
SET balance = balance - ?
WHERE id = ?
`

type WithdrawParams struct {
	Balance float64 `json:"balance"`
	ID      string  `json:"id"`
}

func (q *Queries) Withdraw(ctx context.Context, arg WithdrawParams) error {
	_, err := q.db.ExecContext(ctx, withdraw, arg.Balance, arg.ID)
	return err
}
