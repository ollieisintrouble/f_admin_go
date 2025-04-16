package db

import (
	"time"
)

type Asset struct {
	Id          int64     `json:"id"`
	Title       string    `json:"title"`
	Cost        int64     `json:"cost"`
	Description string    `json:"description"`
	CreatedBy   string    `json:"createdBy"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type Product struct {
	Id          int64     `json:"id"`
	ProductName string    `json:"productName"`
	Description string    `json:"description"`
	ProductUrl  string    `json:"productUrl"`
	CreatedBy   string    `json:"createdBy"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type Transactions struct {
	Id          int64     `json:"id"`
	Amount      int64     `json:"amount"`
	Description string    `json:"description"`
	Method      string    `json:"method"`
	CreatedBy   string    `json:"createdBy"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type User struct {
	Id           string    `json:"id"`
	Username     string    `json:"username"`
	PasswordHash string    `json:"passwordHash"`
	FullName     string    `json:"fullName"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	Image        string    `json:"image"`
	CreatedAt    time.Time `json:"createdAt"`
}
