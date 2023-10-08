package model

import "time"

type Company struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name" gorm:"not null"`
	Description  string    `json:"description" gorm:"not null"`
	OpenSalary   string    `json:"open_salary" gorm:"type:VARCHAR(512); unique; index"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Technologies []Technology `gorm:"many2many:company_technologies;"`
}

type CompanyResponse struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name     string    `json:"name" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	OpenSalary string `json:"open_salary" gorm:"type:VARCHAR(512); unique; index"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}