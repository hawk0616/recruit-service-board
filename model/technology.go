package model

import "time"

type Technology struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name     string    `json:"name" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Companies   []Company `gorm:"many2many:company_technologies;"`
}

type TechnologyResponse struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name     string    `json:"name" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}