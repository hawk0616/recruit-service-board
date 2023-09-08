package model

type Comment struct {
	ID           uint    `json:"id" gorm:"primaryKey"`
	Content      string  `json:"content" gorm:"not null"`
	UserID       uint    `json:"user_id" gorm:"not null"`
	User         User    `json:"user" gorm:"foreignKey:UserID; constraint:OnDelete:CASCADE"`
	CompanyID    uint    `json:"company_id" gorm:"not null"`
	Company      Company `json:"company" gorm:"foreignKey:CompanyID; constraint:OnDelete:CASCADE"`
}

type CommentResponse struct {
	Content      string  `json:"content" gorm:"not null"`
	UserID       uint    `json:"user_id" gorm:"not null"`
	CompanyID    uint    `json:"company_id" gorm:"not null"`
}