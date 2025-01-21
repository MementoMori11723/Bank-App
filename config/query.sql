-- name: CreateAccount :exec
INSERT INTO account (id, first_name, last_name, username, email, password, balance)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: Deposit :exec
UPDATE account
SET balance = balance + ?
WHERE id = ? OR username = ?;

-- name: Withdraw :exec
UPDATE account
SET balance = balance - ?
WHERE id = ?;

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

-- name: GetAccountByUsername :one
SELECT id, first_name, last_name, username, email, balance
FROM account
WHERE username = ? AND password = ?;

-- name: DeleteAccount :exec
DELETE FROM account
WHERE id = ?;

-- name: DeleteHistory :exec
DELETE FROM history
WHERE (sender = ? OR receiver = ?)
  AND NOT EXISTS (SELECT 1 FROM account WHERE username = sender)
  AND NOT EXISTS (SELECT 1 FROM account WHERE username = receiver);

