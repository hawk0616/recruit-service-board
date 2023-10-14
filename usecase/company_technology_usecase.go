package usecase

import (
	"recruit-info-service/model"
	"recruit-info-service/repository"
)

type TechnologyResponseForCompany struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ICompanyTechnologyUsecase interface {
	GetCompanyTechnologyByCompanyId(companyId uint) ([]TechnologyResponseForCompany, error)
	CreateCompanyTechnology(companyTechnology model.CompanyTechnology) (model.CompanyTechnologyResponse, error)
	DeleteCompanyTechnology(companyId uint, technologyId uint) error
}

type CompanyTechnologyUsecase struct {
	ctr repository.ICompanyTechnologyRepository
}

func NewCompanyTechnologyUsecase(ctr repository.ICompanyTechnologyRepository) ICompanyTechnologyUsecase {
	return &CompanyTechnologyUsecase{ctr}
}

func (ctu *CompanyTechnologyUsecase) GetCompanyTechnologyByCompanyId(companyId uint) ([]TechnologyResponseForCompany, error) {
	technologies, err := ctu.ctr.GetCompanyTechnologyByCompanyId(companyId)
	if err != nil {
			return nil, err
	}

	var techResponses []TechnologyResponseForCompany
	for _, tech := range technologies {
		techResponses = append(techResponses, TechnologyResponseForCompany{
			Name:        tech.Name,
			Description: tech.Description,
		})
	}

	return techResponses, nil
}

func (ctu *CompanyTechnologyUsecase) CreateCompanyTechnology(companyTechnology model.CompanyTechnology) (model.CompanyTechnologyResponse, error) {
	if err := ctu.ctr.CreateCompanyTechnology(&companyTechnology); err != nil {
		return model.CompanyTechnologyResponse{}, err
	}
	resCompanyTechnology := model.CompanyTechnologyResponse{
		ID: companyTechnology.ID,
		CompanyID: companyTechnology.CompanyID,
		TechnologyID: companyTechnology.TechnologyID,
	}
	return resCompanyTechnology, nil
}

func (ctu *CompanyTechnologyUsecase) DeleteCompanyTechnology(companyId uint, technologyId uint) error {
	if err := ctu.ctr.DeleteCompanyTechnology(companyId, technologyId); err != nil {
		return err
	}
	return nil
}