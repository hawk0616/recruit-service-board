package repository

import (
	"fmt"
	"recruit-info-service/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ICompanyRepository interface {
	GetAllCompanies(companies *[]model.Company, userId uint) error
	GetCompanyById(company *model.Company, userId uint, companyId uint) error
	CreateCompany(company *model.Company) error
	UpdateCompany(company *model.Company, userId uint, companyId uint) error
	DeleteCompany(userId uint, companyId uint) error
}

type CompanyRepository struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) ICompanyRepository {
	return &CompanyRepository{db}
}

func (cr *CompanyRepository) GetAllCompanies(companies *[]model.Company, userId uint) error {
	if err := cr.db.Joins("User").Where("user_id = ?", userId).Order("created_at").Find(&companies).Error; err != nil {
		return err
	}
	return nil
}

func (cr *CompanyRepository) GetCompanyById(company *model.Company, userId uint, companyId uint) error {
	if err := cr.db.Joins("User").Where("user_id = ?", userId).First(company, companyId).Error; err != nil {
		return err
	}
	return nil
}

func (cr *CompanyRepository) CreateCompany(company *model.Company) error {
	if err := cr.db.Create(company).Error; err != nil {
		return err
	}
	return nil
}

func (cr *CompanyRepository) UpdateCompany(company *model.Company, userId uint, companyId uint) error {
	company.ID = companyId
	result := cr.db.Model(company).Clauses(clause.Returning{}).Where("id=? AND user_id=?", companyId, userId).Updates(
		map[string]interface{}{
			"Name":        company.Name,
			"Description": company.Description,
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

func (cr *CompanyRepository) DeleteCompany(userId uint, companyId uint) error {
	result := cr.db.Where("id=? AND user_id=?", companyId, userId).Delete(&model.Company{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("record not found")
	}
	
	return nil
}