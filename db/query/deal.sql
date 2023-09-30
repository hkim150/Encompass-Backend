-- name: CreateDeal :one
INSERT INTO deal (
  author, store_name, description, regular_price, sale_price
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetDeal :one
SELECT * FROM deal
WHERE deal_id = $1 LIMIT 1;

-- name: ListDeals :many
SELECT * FROM deal
ORDER BY deal_id
LIMIT $1
OFFSET $2;

-- name: UpdateDeal :one
UPDATE deal
  set store_name = $2,
  description = $3,
  regular_price = $4,
  sale_price = $5
WHERE deal_id = $1
RETURNING *;

-- name: DeleteDeal :exec
DELETE FROM deal
WHERE deal_id = $1;
