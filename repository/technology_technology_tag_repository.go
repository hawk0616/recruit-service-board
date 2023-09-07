package repository

import (
	"fmt"
	"recruit-info-service/model"

	"gorm.io/gorm"
)

type ITechnologyTechnologyTagRepository interface {
	CreateTechnologyTechnologyTag(technology_technology_tag *model.TechnologyTechnologyTag) error
	DeleteTechnologyTechnologyTag(technologyId uint, TechnologyTagId uint) error
}

type TechnologyTechnologyTagRepository struct {
	db *gorm.DB
}

func NewTechnologyTechnologyTagRepository(db *gorm.DB) ITechnologyTechnologyTagRepository {
	return &TechnologyTechnologyTagRepository{db}
}

func (tttr *TechnologyTechnologyTagRepository) CreateTechnologyTechnologyTag(technology_technology_tag *model.TechnologyTechnologyTag) error {
	if err := tttr.db.Create(technology_technology_tag).Error; err != nil {
		return err
	}
	return nil
}

func (tttr *TechnologyTechnologyTagRepository) DeleteTechnologyTechnologyTag(technologyId uint, TechnologyTagId uint) error {
	result := tttr.db.Where("technology_id = ? AND technology_tag_id = ?", technologyId, TechnologyTagId).Delete(&model.TechnologyTechnologyTag{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 0 {
		return fmt.Errorf("record not found")
	}
	return nil
}