package db

import (
	"database/sql"
	"time"
)

type Asset struct {
	ID           int64          `json:"id" db:"id"`
	Title        string         `json:"title" db:"title"`
	Cost         int64          `json:"cost" db:"cost"`
	Description  sql.NullString `json:"description" db:"description"`
	CreatedBy    sql.NullString `json:"createdBy" db:"created_by"`
	CreatedAt    time.Time      `json:"createdAt" db:"created_at"`
	UpdatedAt    time.Time      `json:"updatedAt" db:"updated_at"`
	Status       sql.NullString `json:"status" db:"status"`
	Type         sql.NullString `json:"type" db:"type"`
	PurchaseDate sql.NullTime   `json:"purchaseDate" db:"purchase_date"`
	Organization int64          `json:"organization" db:"organization"`
}

type Transaction struct {
	ID           int64          `json:"id" db:"id"`
	Amount       int64          `json:"amount" db:"amount"`
	Description  sql.NullString `json:"description" db:"description"`
	Method       sql.NullString `json:"method" db:"method"`
	CreatedBy    sql.NullString `json:"createdBy" db:"created_by"`
	CreatedAt    time.Time      `json:"createdAt" db:"created_at"`
	UpdatedAt    time.Time      `json:"updatedAt" db:"updated_at"`
	Status       sql.NullString `json:"status" db:"status"`
	Type         sql.NullString `json:"type" db:"type"`
	RecordedDate sql.NullTime   `json:"purchaseDate" db:"purchase_date"`
	Organization int64          `json:"organization" db:"organization"`
}

type Product struct {
	ID           int64          `json:"id" db:"id"`
	ProductName  string         `json:"productName" db:"product_name"`
	Description  sql.NullString `json:"description" db:"description"`
	ProductUrl   sql.NullString `json:"productUrl" db:"product_url"`
	CreatedBy    sql.NullString `json:"createdBy" db:"created_by"`
	CreatedAt    time.Time      `json:"createdAt" db:"created_at"`
	UpdatedAt    time.Time      `json:"updatedAt" db:"updated_at"`
	Organization int64          `json:"organization" db:"organization"`
}

type User struct {
	ID           string         `json:"id" db:"id"`
	Username     string         `json:"username" db:"username"`
	PasswordHash string         `json:"passwordHash" db:"password_hash"`
	FullName     sql.NullString `json:"fullName" db:"full_name"`
	Email        sql.NullString `json:"email" db:"email"`
	Phone        sql.NullString `json:"phone" db:"phone"`
	Image        sql.NullString `json:"image" db:"image"`
	CreatedAt    time.Time      `json:"createdAt" db:"created_at"`
}

type Organization struct {
	ID              int64          `json:"id" db:"id"`
	Name            string         `json:"name" db:"name"`
	Image           sql.NullString `json:"image" db:"image"`
	CreatedAt       time.Time      `json:"createdAt" db:"created_at"`
	PurchasePackage sql.NullString `json:"purchasePackage" db:"purchase_package"`
}

type Membership struct {
	UserID         string `json:"userId" db:"user_id"`
	OrganizationID int64  `json:"organizationId" db:"organization_id"`
	Role           string `json:"role" db:"role"`
}
