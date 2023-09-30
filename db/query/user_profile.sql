-- name: CreateUserProfile :one
INSERT INTO user_profile (
  zipcode, timezone
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetUserProfile :one
SELECT * FROM user_profile
WHERE user_account_id = $1 LIMIT 1;

-- name: ListUserProfiles :many
SELECT * FROM user_profile
ORDER BY user_account_id
LIMIT $1
OFFSET $2;

-- name: UpdateUserProfile :one
UPDATE user_profile
  set zipcode = $2,
  timezone = $3
WHERE user_account_id = $1
RETURNING *;

-- name: DeleteUserProfile :exec
DELETE FROM user_profile
WHERE user_account_id = $1;
