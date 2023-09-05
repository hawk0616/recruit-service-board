package repository

import (
	"fmt"
	"recruit-info-service/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ITechnologyRepository interface {
	GetAllTechnologies(technologies *[]model.Technology, userId uint) error
	GetTechnologyById(technology *model.Technology, userId uint, technologyId uint) error
	CreateTechnology(technology *model.Technology) error
	UpdateTechnology(technology *model.Technology, userId uint, technologyId uint) error
	DeleteTechnology(userId uint, technologyId uint) error
}

type TechnologyRepository struct {
	db *gorm.DB
}

func NewTechnologyRepository(db *gorm.DB) ITechnologyRepository {
	return &TechnologyRepository{db}
}

func (tr *TechnologyRepository) GetAllTechnologies(technologies *[]model.Technology, userId uint) error {
	if err := tr.db.Joins("User").Where("user_id = ?", userId).Order("created_at").Find(&technologies).Error; err != nil {
		return err
	}
	return nil
}

func (tr *TechnologyRepository) GetTechnologyById(technology *model.Technology, userId uint, technologyId uint) error {
	if err := tr.db.Joins("User").Where("user_id = ?", userId).First(technology, technologyId).Error; err != nil {
		return err
	}
	return nil
}

func (tr *TechnologyRepository) CreateTechnology(technology *model.Technology) error {
	if err := tr.db.Create(technology).Error; err != nil {
		return err
	}
	return nil
}

func (tr *TechnologyRepository) UpdateTechnology(technology *model.Technology, userId uint, technologyId uint) error {
	technology.ID = technologyId
	result := tr.db.Model(technology).Clauses(clause.Returning{}).Where("id=? AND user_id=?", technologyId, userId).Updates(
		map[string]interface{}{
			"Name":        technology.Name,
			"Description": technology.Description,
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

func (tr *TechnologyRepository) DeleteTechnology(userId uint, technologyId uint) error {
	result := tr.db.Where("id=? AND user_id=?", technologyId, userId).Delete(&model.Technology{}, technologyId)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 0 {
		return fmt.Errorf("record not found")
	}
	return nil
}