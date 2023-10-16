package usecase_test

import (
    "testing"
    "recruit-info-service/model"
    "recruit-info-service/usecase"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

type MockCompanyRepository struct {
    mock.Mock
}

func (m *MockCompanyRepository) GetAllCompanies(companies *[]model.Company) error {
    args := m.Called(companies)
    return args.Error(0)
}

func (m *MockCompanyRepository) GetCompanyById(company *model.Company, companyId uint) error {
    args := m.Called(company, companyId)
    return args.Error(0)
}

func (m *MockCompanyRepository) CreateCompany(company *model.Company) error {
    args := m.Called(company)
    return args.Error(0)
}

func (m *MockCompanyRepository) UpdateCompany(company *model.Company, companyId uint) error {
    args := m.Called(company, companyId)
    return args.Error(0)
}

func (m *MockCompanyRepository) DeleteCompany(companyId uint) error {
    args := m.Called(companyId)
    return args.Error(0)
}

type MockCompanyValidator struct {
    mock.Mock
}

func (m *MockCompanyValidator) CompanyValidate(company model.Company) error {
    args := m.Called(company)
    return args.Error(0)
}

func TestGetAllCompanies(t *testing.T) {
	mockRepo := &MockCompanyRepository{}
	mockValidator := &MockCompanyValidator{}
	mockUsecase := usecase.NewCompanyUsecase(mockRepo, mockValidator)
	expectCompanies := []model.Company{
			{ID: 1, Name: "Company A", Description: "Description A", OpenSalary: "OpenSalary A", Address: "Address A"},
			{ID: 2, Name: "Company B", Description: "Description B", OpenSalary: "OpenSalary B", Address: "Address B"},
	}

	mockRepo.On("GetAllCompanies", mock.AnythingOfType("*[]model.Company")).Return(nil).Run(func(args mock.Arguments) {
			companies := args.Get(0).(*[]model.Company)
			*companies = expectCompanies
	})

	companiesRes, err := mockUsecase.GetAllCompanies()
	
	assert.NoError(t, err)
	assert.Len(t, companiesRes, 2)
	assert.Equal(t, "Company A", companiesRes[0].Name)
	assert.Equal(t, "Company B", companiesRes[1].Name)

	mockRepo.AssertExpectations(t)
}

func TestGetCompanyById(t *testing.T) {
	mockRepo := &MockCompanyRepository{}
	mockValidator := &MockCompanyValidator{}
	mockUsecase := usecase.NewCompanyUsecase(mockRepo, mockValidator)
	expectCompany := model.Company{ID: 1, Name: "Company A", Description: "Description A", OpenSalary: "OpenSalary A", Address: "Address A"}

	mockRepo.On("GetCompanyById", mock.AnythingOfType("*model.Company"), mock.AnythingOfType("uint")).Return(nil).Run(func(args mock.Arguments) {
			company := args.Get(0).(*model.Company)
			*company = expectCompany
	})

	companyRes, err := mockUsecase.GetCompanyById(1)
	
	assert.NoError(t, err)
	assert.Equal(t, "Company A", companyRes.Name)

	mockRepo.AssertExpectations(t)
}

func TestCreateCompany(t *testing.T) {
	mockRepo := &MockCompanyRepository{}
	mockValidator := &MockCompanyValidator{}
	mockUsecase := usecase.NewCompanyUsecase(mockRepo, mockValidator)
	expectCompany := model.Company{ID: 1, Name: "Company A", Description: "Description A", OpenSalary: "OpenSalary A", Address: "Address A"}
	
	mockValidator.On("CompanyValidate", mock.AnythingOfType("model.Company")).Return(nil)

	mockRepo.On("CreateCompany", mock.AnythingOfType("*model.Company")).Return(nil).Run(func(args mock.Arguments) {
			company := args.Get(0).(*model.Company)
			*company = expectCompany
	})

	companyRes, err := mockUsecase.CreateCompany(expectCompany)
	
	assert.NoError(t, err)
	assert.Equal(t, "Company A", companyRes.Name)

	mockRepo.AssertExpectations(t)
}

func TestUpdateCompany(t *testing.T) {
	mockRepo := &MockCompanyRepository{}
	mockValidator := &MockCompanyValidator{}
	mockUsecase := usecase.NewCompanyUsecase(mockRepo, mockValidator)
	expectCompany := model.Company{ID: 1, Name: "Company A", Description: "Description A", OpenSalary: "OpenSalary A", Address: "Address A"}

	mockValidator.On("CompanyValidate", mock.AnythingOfType("model.Company")).Return(nil)

	mockRepo.On("UpdateCompany", mock.AnythingOfType("*model.Company"), mock.AnythingOfType("uint")).Return(nil).Run(func(args mock.Arguments) {
			company := args.Get(0).(*model.Company)
			*company = expectCompany
	})

	companyRes, err := mockUsecase.UpdateCompany(expectCompany, 1)
	
	assert.NoError(t, err)
	assert.Equal(t, "Company A", companyRes.Name)

	mockRepo.AssertExpectations(t)
}

func TestDeleteCompany(t *testing.T) {
	mockRepo := &MockCompanyRepository{}
	mockValidator := &MockCompanyValidator{}
	mockUsecase := usecase.NewCompanyUsecase(mockRepo, mockValidator)

	mockRepo.On("DeleteCompany", mock.AnythingOfType("uint")).Return(nil)

	err := mockUsecase.DeleteCompany(1)
	
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}