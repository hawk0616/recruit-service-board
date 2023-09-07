package model

type TechnologyTechnologyTag struct {
	ID               uint          `json:"id" gorm:"primaryKey"`
	TechnologyID     uint          `json:"technology_id" gorm:"not null"`
	Technology       Technology    `json:"technology" gorm:"foreignKey:TechnologyID; constraint:OnDelete:CASCADE"`
	TechnologyTagID  uint          `json:"technology_tag_id" gorm:"not null"`
	TechnologyTag    TechnologyTag `json:"technology_tag" gorm:"foreignKey:TechnologyTagID; constraint:OnDelete:CASCADE"`
}

type TechnologyTechnologyTagResponse struct {
	ID               uint `json:"id" gorm:"primaryKey"`
	TechnologyID     uint `json:"technology_id" gorm:"not null"`
	TechnologyTagID  uint `json:"technology_tag_id" gorm:"not null"`
}