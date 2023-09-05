package model

type CompanyTechnology struct {
	ID           uint `json:"id" gorm:"primaryKey"`
	CompanyID    uint `json:"company_id" gorm:"not null"`
	Company      Company `json:"company" gorm:"foreignKey:CompanyID; constraint:OnDelete:CASCADE"`
	TechnologyID uint `json:"technology_id" gorm:"not null"`
	Technology   Technology `json:"technology" gorm:"foreignKey:TechnologyID; constraint:OnDelete:CASCADE"`
}

type CompanyTechnologyResponse struct {
	ID           uint `json:"id" gorm:"primaryKey"`
	CompanyID    uint `json:"company_id" gorm:"not null"`
	TechnologyID uint `json:"technology_id" gorm:"not null"`
}