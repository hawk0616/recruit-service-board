package repository

import (
	"fmt"
	"recruit-info-service/model"

	"gorm.io/gorm"
)

type ICompanyTechnologyRepository interface {
	GetCompanyTechnologyByCompanyId(companyId uint) ([]model.Technology, error)
	CreateCompanyTechnology(company_technology *model.CompanyTechnology) error
	DeleteCompanyTechnology(companyId uint, technologyId uint) error
}

type CompanyTechnologyRepository struct {
	db *gorm.DB
}

func NewCompanyTechnologyRepository(db *gorm.DB) ICompanyTechnologyRepository {
	return &CompanyTechnologyRepository{db}
}

func (ctr *CompanyTechnologyRepository) GetCompanyTechnologyByCompanyId(companyId uint) ([]model.Technology, error) {
	var technologies []model.Technology

	err := ctr.db.Table("technologies as t").
			Select("t.name, t.description").
			Joins("LEFT OUTER JOIN company_technologies AS ct ON t.id = ct.technology_id").
			Where("ct.company_id = ?", companyId).
			Scan(&technologies).Error

	if err != nil {
			return nil, err
	}

	return technologies, nil
}

func (ctr *CompanyTechnologyRepository) CreateCompanyTechnology(company_technology *model.CompanyTechnology) error {
	if err := ctr.db.Create(company_technology).Error; err != nil {
		return err
	}
	return nil
}

func (ctr *CompanyTechnologyRepository) DeleteCompanyTechnology(companyId uint, technologyId uint) error {
	result := ctr.db.Where("company_id = ? AND technology_id = ?", companyId, technologyId).Delete(&model.CompanyTechnology{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 0 {
		return fmt.Errorf("record not found")
	}
	return nil
}