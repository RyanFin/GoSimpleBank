-- File is created so that when I run `$ sqlc generate`, the code will generate CRUD code in the sqlc/ dir 
-- name: CreateAccount :one
INSERT INTO accounts (
  owner, balance, currency
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: GetAccountForUpdate :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1
-- Tell postgres that we do not update the key or id column to prevent deadlock
FOR NO KEY UPDATE;

-- name: ListAccounts :many
SELECT * FROM accounts
WHERE owner = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateAccount :one
UPDATE accounts 
SET balance = $2
WHERE id = $1
RETURNING *;

-- name: AddAccountBalance :one
UPDATE accounts 
-- set 'amount' to be the name of this generated parameter
SET balance = balance + sqlc.arg(amount)
-- set 'id' to be the name of this generated parameter
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts WHERE id = $1;