// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: deal.sql

package db

import (
	"context"
	"database/sql"
)

const createDeal = `-- name: CreateDeal :one
INSERT INTO deal (
  author, store_name, description, regular_price, sale_price
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING deal_id, author, store_name, description, regular_price, sale_price, upvote, create_time
`

type CreateDealParams struct {
	Author       int64
	StoreName    string
	Description  string
	RegularPrice sql.NullString
	SalePrice    sql.NullString
}

func (q *Queries) CreateDeal(ctx context.Context, arg CreateDealParams) (Deal, error) {
	row := q.db.QueryRowContext(ctx, createDeal,
		arg.Author,
		arg.StoreName,
		arg.Description,
		arg.RegularPrice,
		arg.SalePrice,
	)
	var i Deal
	err := row.Scan(
		&i.DealID,
		&i.Author,
		&i.StoreName,
		&i.Description,
		&i.RegularPrice,
		&i.SalePrice,
		&i.Upvote,
		&i.CreateTime,
	)
	return i, err
}

const deleteDeal = `-- name: DeleteDeal :exec
DELETE FROM deal
WHERE deal_id = $1
`

func (q *Queries) DeleteDeal(ctx context.Context, dealID int64) error {
	_, err := q.db.ExecContext(ctx, deleteDeal, dealID)
	return err
}

const getDeal = `-- name: GetDeal :one
SELECT deal_id, author, store_name, description, regular_price, sale_price, upvote, create_time FROM deal
WHERE deal_id = $1 LIMIT 1
`

func (q *Queries) GetDeal(ctx context.Context, dealID int64) (Deal, error) {
	row := q.db.QueryRowContext(ctx, getDeal, dealID)
	var i Deal
	err := row.Scan(
		&i.DealID,
		&i.Author,
		&i.StoreName,
		&i.Description,
		&i.RegularPrice,
		&i.SalePrice,
		&i.Upvote,
		&i.CreateTime,
	)
	return i, err
}

const listDeals = `-- name: ListDeals :many
SELECT deal_id, author, store_name, description, regular_price, sale_price, upvote, create_time FROM deal
ORDER BY deal_id
LIMIT $1
OFFSET $2
`

type ListDealsParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) ListDeals(ctx context.Context, arg ListDealsParams) ([]Deal, error) {
	rows, err := q.db.QueryContext(ctx, listDeals, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Deal
	for rows.Next() {
		var i Deal
		if err := rows.Scan(
			&i.DealID,
			&i.Author,
			&i.StoreName,
			&i.Description,
			&i.RegularPrice,
			&i.SalePrice,
			&i.Upvote,
			&i.CreateTime,
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

const updateDeal = `-- name: UpdateDeal :one
UPDATE deal
  set store_name = $2,
  description = $3,
  regular_price = $4,
  sale_price = $5
WHERE deal_id = $1
RETURNING deal_id, author, store_name, description, regular_price, sale_price, upvote, create_time
`

type UpdateDealParams struct {
	DealID       int64
	StoreName    string
	Description  string
	RegularPrice sql.NullString
	SalePrice    sql.NullString
}

func (q *Queries) UpdateDeal(ctx context.Context, arg UpdateDealParams) (Deal, error) {
	row := q.db.QueryRowContext(ctx, updateDeal,
		arg.DealID,
		arg.StoreName,
		arg.Description,
		arg.RegularPrice,
		arg.SalePrice,
	)
	var i Deal
	err := row.Scan(
		&i.DealID,
		&i.Author,
		&i.StoreName,
		&i.Description,
		&i.RegularPrice,
		&i.SalePrice,
		&i.Upvote,
		&i.CreateTime,
	)
	return i, err
}
