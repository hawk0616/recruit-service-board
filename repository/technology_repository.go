package repository

import (
	"fmt"
	"recruit-info-service/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ITechnologyRepository interface {
	GetAllTechnologies(technologies *[]model.Technology) error
	GetTechnologyById(technology *model.Technology, technologyId uint) error
	CreateTechnology(technology *model.Technology) error
	UpdateTechnology(technology *model.Technology, technologyId uint) error
	DeleteTechnology(technologyId uint) error
}

type TechnologyRepository struct {
	db *gorm.DB
}

func NewTechnologyRepository(db *gorm.DB) ITechnologyRepository {
	return &TechnologyRepository{db}
}

func (tr *TechnologyRepository) GetAllTechnologies(technologies *[]model.Technology) error {
	if err := tr.db.Order("created_at").Find(&technologies).Error; err != nil {
		return err
	}
	return nil
}

func (tr *TechnologyRepository) GetTechnologyById(technology *model.Technology, technologyId uint) error {
	if err := tr.db.First(technology, technologyId).Error; err != nil {
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

func (tr *TechnologyRepository) UpdateTechnology(technology *model.Technology, technologyId uint) error {
	technology.ID = technologyId
	result := tr.db.Model(technology).Clauses(clause.Returning{}).Where("id=? AND user_id=?", technologyId).Updates(
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

func (tr *TechnologyRepository) DeleteTechnology(technologyId uint) error {
	result := tr.db.Where("id=? AND user_id=?", technologyId).Delete(&model.Technology{}, technologyId)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 0 {
		return fmt.Errorf("record not found")
	}
	return nil
}