-- name: CreateAccount :exec
INSERT INTO account (id, first_name, last_name, username, email, password, balance, image_url)
VALUES (?, ?, ?, ?, ?, ?, ?, ?);

-- name: Deposit :exec
UPDATE account
SET balance = balance + ?
WHERE id = ? OR username = ?;

-- name: Withdraw :exec
UPDATE account
SET balance = balance - ?
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

-- name: CheckUserExists :one
SELECT id FROM account WHERE username = ?;

-- name: DeleteAccount :exec
DELETE FROM account
WHERE id = ?;

-- name: DeleteHistory :exec
DELETE FROM history
WHERE (sender = ? OR receiver = ?)
  AND NOT EXISTS (SELECT 1 FROM account WHERE username = history.sender)
  AND NOT EXISTS (SELECT 1 FROM account WHERE username = history.receiver);
