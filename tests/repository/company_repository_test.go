package repository_test

import (
	"recruit-info-service/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type CompanyDBHandler interface {
	GetAllCompanies() ([]model.Company, error)
	GetCompanyById(companyID uint) (model.Company, error)
	CreateCompany(company *model.Company) error
	UpdateCompany(company *model.Company, companyID uint) error
	DeleteCompany(companyID uint) error
}

type RealCompanyDBHandler struct {
	db *gorm.DB
}

func NewRealCompanyDBHandler(db *gorm.DB) *RealCompanyDBHandler {
	return &RealCompanyDBHandler{db}
}

type MockCompanyDBHandler struct {
	mock.Mock
}

func NewMockCompanyDBHandler() *MockCompanyDBHandler {
	return &MockCompanyDBHandler{}
}

func (m *MockCompanyDBHandler) GetAllCompanies() ([]model.Company, error) {
	args := m.Called()
	return args.Get(0).([]model.Company), args.Error(1)
}

func (m *MockCompanyDBHandler) GetCompanyById(companyID uint) (model.Company, error) {
	args := m.Called(companyID)
	return args.Get(0).(model.Company), args.Error(1)
}

func (m *MockCompanyDBHandler) CreateCompany(company *model.Company) error {
	args := m.Called(company)
	return args.Error(0)
}

func (m *MockCompanyDBHandler) UpdateCompany(company *model.Company, companyID uint) error {
	args := m.Called(company, companyID)
	return args.Error(0)
}

func (m *MockCompanyDBHandler) DeleteCompany(companyID uint) error {
	args := m.Called(companyID)
	return args.Error(0)
}

func TestGetAllCompanies(t *testing.T) {
	mockDB := NewMockCompanyDBHandler()
	expectCompanies := []model.Company{
			{ID: 1, Name: "Company A"},
			{ID: 2, Name: "Company B"},
	}

	mockDB.On("GetAllCompanies").Return(expectCompanies, nil)

	companies, err := mockDB.GetAllCompanies()
	mockDB.AssertExpectations(t)

	assert.NoError(t, err)
	assert.Equal(t, expectCompanies, companies)
}

func TestGetCompanyById(t *testing.T) {
	mockDB := NewMockCompanyDBHandler()
	expectCompany := model.Company{ID: 1, Name: "Company A"}
	mockDB.On("GetCompanyById", uint(1)).Return(expectCompany, nil)

	company, err := mockDB.GetCompanyById(uint(1))
	mockDB.AssertExpectations(t)

	assert.NoError(t, err)
	assert.Equal(t, expectCompany, company)
}

func TestCreateCompany(t *testing.T) {
	mockDB := NewMockCompanyDBHandler()
	company := model.Company{ID: 1, Name: "Company A"}

	mockDB.On("CreateCompany", &company).Return(nil)
	err := mockDB.CreateCompany(&company)
	mockDB.AssertExpectations(t)

	assert.NoError(t, err)
}

func TestUpdateCompany(t *testing.T) {
	mockDB := NewMockCompanyDBHandler()
	company := model.Company{ID: 1, Name: "Company A"}
	mockDB.On("UpdateCompany", &company, uint(1)).Return(nil)

	err := mockDB.UpdateCompany(&company, uint(1))
	mockDB.AssertExpectations(t)

	assert.NoError(t, err)
}

func TestDeleteCompany(t *testing.T) {
	mockDB := NewMockCompanyDBHandler()
	mockDB.On("DeleteCompany", uint(1)).Return(nil)
	err := mockDB.DeleteCompany(uint(1))

	mockDB.AssertExpectations(t)
	assert.NoError(t, err)
}