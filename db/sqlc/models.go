// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package db

import (
	"database/sql"
	"time"
)

type Deal struct {
	DealID       int64
	Author       int64
	StoreName    string
	Description  string
	RegularPrice sql.NullString
	SalePrice    sql.NullString
	Upvote       int64
	CreateTime   time.Time
}

type UserAccount struct {
	UserAccountID int64
	Username      string
	Email         string
	CreateTime    time.Time
}

type UserProfile struct {
	UserAccountID int64
	Zipcode       sql.NullString
	Timezone      sql.NullString
}