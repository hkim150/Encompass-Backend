-- name: GetDeal :one
SELECT * FROM deal
WHERE deal_id = $1 LIMIT 1;
