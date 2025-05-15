package models

import (
	"f_admin_go/internal/db"
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

type TransactionForm struct {
	Amount       int64      `json:"amount"`
	Description  *string    `json:"description"`
	Method       *string    `json:"method"`
	CreatedBy    *string    `json:"createdBy"`
	Status       *string    `json:"status"`
	Type         *string    `json:"type"`
	RecordedDate *time.Time `json:"recordedDate"`
}

type ProductDTO struct {
	ID           int64      `json:"id"`
	ProductName  string     `json:"productName"`
	Description  *string    `json:"description"`
	ProductURL   *string    `json:"productUrl"`
	CreatedBy    *string    `json:"createdBy"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	Organization int64      `json:"organization"`
	Status       *string    `json:"status"`
	Type         *string    `json:"type"`
	LaunchDate   *time.Time `json:"launchDate"`
	MetricsURL   *string    `json:"metricsUrl"`
	Logo         *string    `json:"logo"`
}

type UserDTO struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	FullName  *string   `json:"fullName"`
	Email     *string   `json:"email"`
	Phone     *string   `json:"phone"`
	Image     *string   `json:"image"`
	CreatedAt time.Time `json:"createdAt"`
	Role      string    `json:"role"`
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
	Image    *string `json:"image"`
	Role     string  `json:"role"`
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

type TokenValidationResponse struct {
	User          UserDTO           `json:"user"`
	Organizations []OrganizationDTO `json:"orgs"`
	Memberships   []db.Membership   `json:"memberships"`
}
