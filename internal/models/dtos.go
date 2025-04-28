package models

import (
	"time"
)

type AssetDTO struct {
	ID           int64      `json:"id"`
	Title        string     `json:"title"`
	Cost         int64      `json:"cost"`
	Description  *string    `json:"description"`
	CreatedBy    *string    `json:"createdBy"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	Status       *string    `json:"status"`
	Type         *string    `json:"type"`
	PurchaseDate *time.Time `json:"purchaseDate"`
}

type TransactionDTO struct {
	ID           int64      `json:"id"`
	Amount       int64      `json:"amount"`
	Description  *string    `json:"description"`
	Method       *string    `json:"method"`
	CreatedBy    *string    `json:"createdBy"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	Status       *string    `json:"status"`
	Type         *string    `json:"type"`
	RecordedDate *time.Time `json:"recordedDate"`
}

type ProductDTO struct {
	ID          int64     `json:"id"`
	ProductName string    `json:"productName"`
	Description *string   `json:"description"`
	ProductUrl  *string   `json:"productUrl"`
	CreatedBy   *string   `json:"createdBy"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type UserDTO struct {
	ID        string     `json:"id" db:"id"`
	Username  string     `json:"username" db:"username"`
	Password  string     `json:"password" db:"password"`
	FullName  *string    `json:"fullName" db:"full_name"`
	Email     *string    `json:"email" db:"email"`
	Phone     *string    `json:"phone" db:"phone"`
	Image     *string    `json:"image" db:"image"`
	CreatedAt *time.Time `json:"createdAt" db:"created_at"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterForm struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	FullName *string `json:"fullName"`
	Email    *string `json:"email"`
	Phone    *string `json:"phone"`
}

type OrganizationDTO struct {
	ID              int64     `json:"id"`
	Name            string    `json:"name"`
	Image           *string   `json:"image"`
	CreatedAt       time.Time `json:"createdAt"`
	PurchasePackage *string   `json:"purchasePackage"`
}

type CreateOrginzationForm struct {
	UserID       string          `json:"userId"`
	Organization OrganizationDTO `json:"organization"`
}
