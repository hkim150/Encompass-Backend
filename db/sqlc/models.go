// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package db

import (
	"database/sql"
	"time"
)

type Deal struct {
	DealID       int64          `json:"deal_id"`
	Author       int64          `json:"author"`
	StoreName    string         `json:"store_name"`
	Description  string         `json:"description"`
	RegularPrice sql.NullString `json:"regular_price"`
	SalePrice    sql.NullString `json:"sale_price"`
	Upvote       int64          `json:"upvote"`
	CreateTime   time.Time      `json:"create_time"`
}

type UserAccount struct {
	UserAccountID int64     `json:"user_account_id"`
	Username      string    `json:"username"`
	Email         string    `json:"email"`
	CreateTime    time.Time `json:"create_time"`
}

type UserProfile struct {
	UserAccountID int64          `json:"user_account_id"`
	Zipcode       sql.NullString `json:"zipcode"`
	Timezone      sql.NullString `json:"timezone"`
}
