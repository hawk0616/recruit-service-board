package repository

import (
	"fmt"
	"recruit-info-service/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ICompanyRepository interface {
	GetAllCompanies(companies *[]model.Company) error
	GetCompanyById(company *model.Company, companyId uint) error
	CreateCompany(company *model.Company) error
	UpdateCompany(company *model.Company, companyId uint) error
	DeleteCompany( companyId uint) error
}

type CompanyRepository struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) ICompanyRepository {
	return &CompanyRepository{db}
}

func (cr *CompanyRepository) GetAllCompanies(companies *[]model.Company) error {
	if err := cr.db.Order("created_at").Find(&companies).Error; err != nil {
		return err
	}
	return nil
}

func (cr *CompanyRepository) GetCompanyById(company *model.Company, companyId uint) error {
	if err := cr.db.First(company, companyId).Error; err != nil {
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

func (cr *CompanyRepository) UpdateCompany(company *model.Company, companyId uint) error {
	company.ID = companyId
	result := cr.db.Model(company).Clauses(clause.Returning{}).Where("id=?", companyId).Updates(
		map[string]interface{}{
			"Name":        company.Name,
			"Description": company.Description,
			"Company": company.OpenSalary,
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

func (cr *CompanyRepository) DeleteCompany(companyId uint) error {
	result := cr.db.Where("id=?", companyId).Delete(&model.Company{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("record not found")
	}
	
	return nil
}