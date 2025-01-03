-- name: CreateAccount :exec
INSERT INTO account (id, first_name, last_name, username, email, password, balance)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: Deposit :exec
UPDATE account
SET balance = balance + ?
WHERE id = ?;

-- name: Withdraw :exec
UPDATE account
SET balance = balance - ?
WHERE id = ? AND balance >= ?;

-- name: GetBalance :one
SELECT balance
FROM account
WHERE id = ?;

-- name: GetTransactions :many
SELECT id, sender, receiver, amount, timestamp
FROM history
WHERE sender = ? OR receiver = ?
ORDER BY timestamp DESC;

-- name: InsertTransaction :exec
INSERT INTO history (id, sender, receiver, amount, timestamp)
VALUES (?, ?, ?, ?, ?);

-- name: GetAccountByID :one
SELECT id, first_name, last_name, username, email, balance
FROM account
WHERE id = ?;

-- name: GetAccountByUsername :one
SELECT id, first_name, last_name, username, email, balance
FROM account
WHERE username = ?;

-- name: DeleteAccount :exec
DELETE FROM account
WHERE id = ?;
