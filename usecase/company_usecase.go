package usecase

import (
	"recruit-info-service/model"
	"recruit-info-service/repository"
	"recruit-info-service/validator"
)

type ICompanyUsecase interface {
	GetAllCompanies() ([]model.CompanyResponse, error)
	GetCompanyById(companyId uint) (model.CompanyResponse, error)
	CreateCompany(company model.Company) (model.CompanyResponse, error)
	UpdateCompany(company model.Company, companyId uint) (model.CompanyResponse, error)
	DeleteCompany(companyId uint) error
}

type CompanyUsecase struct {
	cr repository.ICompanyRepository
	cv validator.ICompanyValidator
}

func NewCompanyUsecase(cr repository.ICompanyRepository, cv validator.ICompanyValidator) ICompanyUsecase {
	return &CompanyUsecase{cr, cv}
}

func (cu *CompanyUsecase) GetAllCompanies() ([]model.CompanyResponse, error) {
	companies := []model.Company{}
	if err := cu.cr.GetAllCompanies(&companies); err != nil {
		return nil, err
	}
	resCompanies := []model.CompanyResponse{}
	for _, v := range companies {
		c := model.CompanyResponse{
			ID: v.ID,
			Name: v.Name,
			Description: v.Description,
			OpenSalary: v.OpenSalary,
			Address: v.Address,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resCompanies = append(resCompanies, c)
	}
	return resCompanies, nil
}

func (cu *CompanyUsecase) GetCompanyById(companyId uint) (model.CompanyResponse, error) {
	company := model.Company{}
	if err := cu.cr.GetCompanyById(&company, companyId); err != nil {
		return model.CompanyResponse{}, err
	}
	resCompany := model.CompanyResponse{
		ID: company.ID,
		Name: company.Name,
		Description: company.Description,
		OpenSalary: company.OpenSalary,
		Address: company.Address,
		CreatedAt: company.CreatedAt,
		UpdatedAt: company.UpdatedAt,
	}
	return resCompany, nil
}

func (cu *CompanyUsecase) CreateCompany(company model.Company) (model.CompanyResponse, error) {
	if err := cu.cv.CompanyValidate(company); err != nil {
		return model.CompanyResponse{}, err
	}
	if err := cu.cr.CreateCompany(&company); err != nil {
		return model.CompanyResponse{}, err
	}
	resCompany := model.CompanyResponse{
		ID: company.ID,
		Name: company.Name,
		Description: company.Description,
		CreatedAt: company.CreatedAt,
		UpdatedAt: company.UpdatedAt,
	}
	return resCompany, nil
}

func (cu *CompanyUsecase) UpdateCompany(company model.Company, companyId uint) (model.CompanyResponse, error) {
	if err := cu.cv.CompanyValidate(company); err != nil {
		return model.CompanyResponse{}, err
	}
	// 成功した場合は第一引数で渡したcompanyのアドレスが示す先のメモリ領域のcompanyの値が書き変わっている
	if err := cu.cr.UpdateCompany(&company, companyId); err != nil {
		return model.CompanyResponse{}, err
	}
	resCompany := model.CompanyResponse{
		ID: company.ID,
		Name: company.Name,
		Description: company.Description,
		CreatedAt: company.CreatedAt,
		UpdatedAt: company.UpdatedAt,
	}
	return resCompany, nil
}

func (company *CompanyUsecase) DeleteCompany(companyId uint) error {
	if err := company.cr.DeleteCompany(companyId); err != nil {
		return err
	}
	return nil
}