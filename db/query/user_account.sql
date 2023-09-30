-- name: CreateUserAccount :one
INSERT INTO user_account (
  username, email
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetUserAccount :one
SELECT * FROM user_account
WHERE user_account_id = $1 LIMIT 1;

-- name: ListUserAccounts :many
SELECT * FROM user_account
ORDER BY user_account_id
LIMIT $1
OFFSET $2;

-- name: UpdateUserAccount :one
UPDATE user_account
  set username = $2,
  email = $3
WHERE user_account_id = $1
RETURNING *;

-- name: DeleteUserAccount :exec
DELETE FROM user_account
WHERE user_account_id = $1;
