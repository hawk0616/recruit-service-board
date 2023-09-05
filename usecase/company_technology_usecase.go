package usecase

import (
	"recruit-info-service/model"
	"recruit-info-service/repository"
)

type ICompanyTechnologyUsecase interface {
	CreateCompanyTechnology(company_technology model.CompanyTechnology) (model.CompanyTechnologyResponse, error)
	DeleteCompanyTechnology(companyId uint, technologyId uint) error
}

type CompanyTechnologyUsecase struct {
	ctr repository.ICompanyTechnologyRepository
}

func NewCompanyTechnologyUsecase(ctr repository.ICompanyTechnologyRepository) ICompanyTechnologyUsecase {
	return &CompanyTechnologyUsecase{ctr}
}

func (ctu *CompanyTechnologyUsecase) CreateCompanyTechnology(companyTechnology model.CompanyTechnology) (model.CompanyTechnologyResponse, error) {
	if err := ctu.ctr.CreateCompanyTechnology(&companyTechnology); err != nil {
		return model.CompanyTechnologyResponse{}, err
	}
	resCompanyTechnology := model.CompanyTechnologyResponse{
		ID: companyTechnology.ID,
	}
	return resCompanyTechnology, nil
}

func (ctu *CompanyTechnologyUsecase) DeleteCompanyTechnology(companyId uint, technologyId uint) error {
	if err := ctu.ctr.DeleteCompanyTechnology(companyId, technologyId); err != nil {
		return err
	}
	return nil
}