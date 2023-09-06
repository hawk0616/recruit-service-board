package repository

import (
	"fmt"
	"recruit-info-service/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ITechnologyTagRepository interface {
	GetAllTechnologyTags(technologyTags *[]model.TechnologyTag, userId uint) error
	GetTechnologyTagById(technologyTag *model.TechnologyTag, userId uint, technologyTagId uint) error
	CreateTechnologyTag(technologyTag *model.TechnologyTag) error
	UpdateTechnologyTag(technologyTag *model.TechnologyTag, userId uint, technologyTagId uint) error
	DeleteTechnologyTag(userId uint, technologyTagId uint) error
}

type TechnologyTagRepository struct {
	db *gorm.DB
}

func NewTechnologyTagRepository(db *gorm.DB) ITechnologyTagRepository {
	return &TechnologyTagRepository{db}
}

func (ttr *TechnologyTagRepository) GetAllTechnologyTags(technologyTags *[]model.TechnologyTag, userId uint) error {
	if err := ttr.db.Joins("User").Where("user_id = ?", userId).Order("created_at").Find(&technologyTags).Error; err != nil {
		return err
	}
	return nil
}

func (ttr *TechnologyTagRepository) GetTechnologyTagById(technologyTag *model.TechnologyTag, userId uint, technologyTagId uint) error {
	if err := ttr.db.Joins("User").Where("user_id = ?", userId).First(technologyTag, technologyTagId).Error; err != nil {
		return err
	}
	return nil
}

func (ttr *TechnologyTagRepository) CreateTechnologyTag(technologyTag *model.TechnologyTag) error {
	if err := ttr.db.Create(technologyTag).Error; err != nil {
		return err
	}
	return nil
}

func (ttr *TechnologyTagRepository) UpdateTechnologyTag(technologyTag *model.TechnologyTag, userId uint, technologyTagId uint) error {
	technologyTag.ID = technologyTagId
	result := ttr.db.Model(technologyTag).Clauses(clause.Returning{}).Where("id=? AND user_id=?", technologyTagId, userId).Updates(
		map[string]interface{}{
			"Name": technologyTag.Name,
		},
	)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("record not found")
	}
	
	return nil
}

func (ttr *TechnologyTagRepository) DeleteTechnologyTag(userId uint, technologyTagId uint) error {
	result := ttr.db.Where("id=? AND user_id =?", technologyTagId, userId).Delete(&model.Technology{}, technologyTagId)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 0 {
		return fmt.Errorf("record not found")
	}
	return nil
}