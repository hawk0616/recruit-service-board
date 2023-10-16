package controller_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"recruit-info-service/controller"
	"recruit-info-service/model"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCompanyUsecase struct {
    mock.Mock
}

func (m *MockCompanyUsecase) GetAllCompanies() ([]model.CompanyResponse, error) {
	args := m.Called()
	return args.Get(0).([]model.CompanyResponse), args.Error(1)
}

func (m *MockCompanyUsecase) GetCompanyById(companyId uint) (model.CompanyResponse, error) {
	args := m.Called(companyId)
	return args.Get(0).(model.CompanyResponse), args.Error(1)
}

func (m *MockCompanyUsecase) CreateCompany(company model.Company) (model.CompanyResponse, error) {
	args := m.Called(company)
	return args.Get(0).(model.CompanyResponse), args.Error(1)
}

func (m *MockCompanyUsecase) UpdateCompany(company model.Company, companyId uint) (model.CompanyResponse, error) {
	args := m.Called(company, companyId)
	return args.Get(0).(model.CompanyResponse), args.Error(1)
}

func (m *MockCompanyUsecase) DeleteCompany(companyId uint) error {
	args := m.Called(companyId)
	return args.Error(0)
}

func TestGetAllCompanies(t *testing.T) {
	mockUsecase := &MockCompanyUsecase{}
	expectCompaniesRes := []model.CompanyResponse{
			{
					ID:          1,
					Name:        "Company A",
					Description: "Description A",
					OpenSalary:  "OpenSalary A",
					Address:     "Address A",
			},
			{
					ID:          2,
					Name:        "Company B",
					Description: "Description B",
					OpenSalary:  "OpenSalary B",
					Address:     "Address B",
			},
	}

	mockUsecase.On("GetAllCompanies").Return(expectCompaniesRes, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/companies", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	controller := controller.NewCompanyController(mockUsecase)
	err := controller.GetAllCompanies(c)
	if assert.NoError(t, err) {
			assert.Equal(t, http.StatusOK, rec.Code)

			var response []model.CompanyResponse
			err := json.Unmarshal(rec.Body.Bytes(), &response)
			if assert.NoError(t, err) {
					assert.Equal(t, expectCompaniesRes, response)
			}
	}
	mockUsecase.AssertExpectations(t)
}

func TestGetCompanyById(t *testing.T) {
	mockUsecase := &MockCompanyUsecase{}
	expectCompanyRes := model.CompanyResponse{
			ID:          1,
			Name:        "Company A",
			Description: "Description A",
			OpenSalary:  "OpenSalary A",
			Address:     "Address A",
	}

	mockUsecase.On("GetCompanyById", uint(1)).Return(expectCompanyRes, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/companies/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("companyId")
	c.SetParamValues("1")

	controller := controller.NewCompanyController(mockUsecase)
	err := controller.GetCompanyById(c)
	if assert.NoError(t, err) {
			assert.Equal(t, http.StatusOK, rec.Code)

			var response model.CompanyResponse
			err := json.Unmarshal(rec.Body.Bytes(), &response)
			if assert.NoError(t, err) {
					assert.Equal(t, expectCompanyRes, response)
			}
	}
	mockUsecase.AssertExpectations(t)
}

func TestCreateCompany(t *testing.T) {
	mockUsecase := &MockCompanyUsecase{}
	expectCompanyRes := model.CompanyResponse{
			ID:          1,
			Name:        "Company A",
			Description: "Description A",
			OpenSalary:  "OpenSalary A",
			Address:     "Address A",
	}

	mockUsecase.On("CreateCompany", mock.Anything).Return(expectCompanyRes, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/companies", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	company := model.Company{
			Name:        "Company A",
			Description: "Description A",
			OpenSalary:  "OpenSalary A",
			Address:     "Address A",
	}
	companyJson, _ := json.Marshal(company)
	companyBody := ioutil.NopCloser(strings.NewReader(string(companyJson)))
	c.Request().Body = companyBody

	controller := controller.NewCompanyController(mockUsecase)
	err := controller.CreateCompany(c)
	if assert.NoError(t, err) {
			assert.Equal(t, http.StatusCreated, rec.Code)

			var response model.CompanyResponse
			err := json.Unmarshal(rec.Body.Bytes(), &response)
			if assert.NoError(t, err) {
					assert.Equal(t, expectCompanyRes, response)
			}
	}
	mockUsecase.AssertExpectations(t)
}

func TestUpdateCompany(t *testing.T) {
	mockUsecase := &MockCompanyUsecase{}
	expectCompanyRes := model.CompanyResponse{
			ID:          1,
			Name:        "Company A",
			Description: "Description A",
			OpenSalary:  "OpenSalary A",
			Address:     "Address A",
	}

	mockUsecase.On("UpdateCompany", mock.Anything, uint(1)).Return(expectCompanyRes, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/companies/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("companyId")
	c.SetParamValues("1")

	company := model.Company{
			Name:        "Company A",
			Description: "Description A",
			OpenSalary:  "OpenSalary A",
			Address:     "Address A",
	}
	companyJson, _ := json.Marshal(company)
	companyBody := ioutil.NopCloser(strings.NewReader(string(companyJson)))
	c.Request().Body = companyBody

	controller := controller.NewCompanyController(mockUsecase)
	err := controller.UpdateCompany(c)
	if assert.NoError(t, err) {
			assert.Equal(t, http.StatusOK, rec.Code)

			var response model.CompanyResponse
			err := json.Unmarshal(rec.Body.Bytes(), &response)
			if assert.NoError(t, err) {
					assert.Equal(t, expectCompanyRes, response)
			}
	}
	mockUsecase.AssertExpectations(t)
}

func TestDeleteCompany(t *testing.T) {
	mockUsecase := &MockCompanyUsecase{}

	mockUsecase.On("DeleteCompany", uint(1)).Return(nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/companies/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("companyId")
	c.SetParamValues("1")

	controller := controller.NewCompanyController(mockUsecase)
	err := controller.DeleteCompany(c)
	if assert.NoError(t, err) {
			assert.Equal(t, http.StatusOK, rec.Code)
	}
	mockUsecase.AssertExpectations(t)
}