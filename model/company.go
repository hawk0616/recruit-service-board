package model

import "time"

type Company struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name     string    `json:"name" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId    uint      `json:"user_id" gorm:"not null"`
	Technologies []Technology `gorm:"many2many:company_technologies;"`
}

type CompanyResponse struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name     string    `json:"name" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}