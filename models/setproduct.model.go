package models

import (
	uuid "github.com/satori/go.uuid"
)

type SetProduct struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name      string    `gorm:"type:varchar(55);not null"`
	Amount    int       `gorm:"not null"`
	Price     float64   `gorm:"not null"`
	ProductID uuid.UUID `gorm:"type:uuid;not null"`
	Type      string    `gorm:"type:varchar(55);not null"`
	Status    int       `gorm:"not null"`
}

type CreateSetProduct struct {
	Name      string    `json:"setproduct_name" binding:"required"`
	Amount    int       `json:"setproduct_amount" binding:"required"`
	Price     float64   `json:"setproduct_price" binding:"required"`
	ProductID uuid.UUID `json:"product_id" binding:"required"`
	Status    int       `json:"status" binding:"required"`
	Type      string    `json:"setproduct_type"`
}

type UpdateSetProduct struct {
	ID        uuid.UUID `json:"id" binding:"required"`
	Name      string    `json:"setproduct_name"`
	Amount    int       `json:"setproduct_amount"`
	Price     float64   `json:"setproduct_price"`
	ProductID uuid.UUID `json:"product_id"`
	Status    int       `json:"status"`
	Type      string    `json:"setproduct_type"`
}

type GetSetProduct struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"setproduct_name"`
	Amount    int       `json:"setproduct_amount"`
	Price     float64   `json:"setproduct_price"`
	ProductID uuid.UUID `json:"product_id"`
	Status    int       `json:"status"`
	Type      string    `json:"setproduct_type"`
}

type DeleteSetProduct struct {
	ID uuid.UUID `json:"id" binding:"required"`
}
