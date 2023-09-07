package model

type Like struct {
	ID           uint    `json:"id" gorm:"primaryKey"`
	UserID       uint    `json:"user_id" gorm:"not null"`
	User         User    `json:"user" gorm:"foreignKey:UserID; constraint:OnDelete:CASCADE"`
	CompanyID    uint    `json:"company_id" gorm:"not null"`
	Company      Company `json:"company" gorm:"foreignKey:CompanyID; constraint:OnDelete:CASCADE"`
}

type LikeResponse struct {
	UserID       uint `json:"user_id" gorm:"not null"`
	CompanyID    uint `json:"company_id" gorm:"not null"`
}