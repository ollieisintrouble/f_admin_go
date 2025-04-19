package db

import (
	"time"
)

type Asset struct {
	Id          int64     `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Cost        int64     `json:"cost" db:"cost"`
	Description string    `json:"description" db:"description"`
	CreatedBy   string    `json:"createdBy" db:"created_by"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
}

type Product struct {
	Id          int64     `json:"id" db:"id"`
	ProductName string    `json:"productName" db:"product_name"`
	Description string    `json:"description" db:"description"`
	ProductUrl  string    `json:"productUrl" db:"product_url"`
	CreatedBy   string    `json:"createdBy" db:"created_by"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
}

type Transactions struct {
	Id          int64     `json:"id" db:"id"`
	Amount      int64     `json:"amount" db:"amount"`
	Description string    `json:"description" db:"description"`
	Method      string    `json:"method" db:"method"`
	CreatedBy   string    `json:"createdBy" db:"created_by"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
}

type User struct {
	Id           string    `json:"id" db:"id"`
	Username     string    `json:"username" db:"username"`
	PasswordHash string    `json:"passwordHash" db:"password_hash"`
	FullName     string    `json:"fullName" db:"full_name"`
	Email        string    `json:"email" db:"email"`
	Phone        string    `json:"phone" db:"phone"`
	Image        string    `json:"image" db:"image"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`
}
