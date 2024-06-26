package models

import (
	uuid "github.com/satori/go.uuid"
)

type Product struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name        string    `gorm:"type:varchar(55);not null"`
	Amount      int       `gorm:"not null"`
	Price       float64   `gorm:"not null"`
	Type        string    `gorm:"type:varchar(55);not null"`
	Category    string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:varchar(255);not null"`
	AttachFile  string
	Status      int `gorm:"not null"`
}

type CreateProduct struct {
	Name       string  `json:"product_name"`
	Amount     int     `json:"product_amount"`
	Price      float64 `json:"product_Price"`
	Type       string  `json:"product_type"`
	Category   string  `json:"product_category"`
	AttachFile string  `json:"attach_file"`
}

type UpdateProduct struct {
	ID          uuid.UUID `json:"product_id"`
	Name        string    `json:"product_name"`
	Amount      int       `json:"product_amount"`
	Price       float64   `json:"product_Price"`
	Type        string    `json:"product_type"`
	Category    string    `json:"product_category"`
	Description string    `json:"product_description"`
	AttachFile  string    `json:"attach_file"`
}

type GetProduct struct {
	ID          uuid.UUID `json:"product_id"`
	Name        string    `json:"product_name"`
	Amount      int       `json:"product_amount"`
	Price       float64   `json:"product_Price"`
	Type        string    `json:"product_type"`
	Category    string    `json:"product_category"`
	Description string    `json:"product_description"`
	AttachFile  string    `json:"attach_file"`
}

type StatusProduct struct {
	Status int `json:"status_product"`
}

type DeleteProduct struct {
	ID uuid.UUID `json:"product_id"`
}
