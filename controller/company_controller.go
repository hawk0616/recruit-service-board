package controller

import (
	"net/http"
	"recruit-info-service/model"
	"recruit-info-service/usecase"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type ICompanyController interface {
	GetAllCompanies(c echo.Context) error
	GetCompanyById(c echo.Context) error
	CreateCompany(c echo.Context) error
	UpdateCompany(c echo.Context) error
	DeleteCompany(c echo.Context) error
}

type companyController struct {
	cu usecase.ICompanyUsecase
}

func NewCompanyController(cu usecase.ICompanyUsecase) ICompanyController {
	return &companyController{cu}
}

func (cc *companyController) GetAllCompanies(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	companiesRes, err := cc.cu.GetAllCompanies(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, companiesRes)
}

func (cc *companyController) GetCompanyById(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("companyId")
	companyId, _ := strconv.Atoi(id)
	companyRes, err := cc.cu.GetCompanyById(uint(userId.(float64)), uint(companyId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, companyRes)
}

func (cc *companyController) CreateCompany(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	company := model.Company{}
	if err := c.Bind(&company); err != nil {
		return err
	}
	company.UserId = uint(userId.(float64))
	companyRes, err := cc.cu.CreateCompany(company)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, companyRes)
}

func (cc *companyController) UpdateCompany(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("companyId")
	companyId, _ := strconv.Atoi(id)

	company := model.Company{}
	if err := c.Bind(&company); err != nil {
		return err
	}
	companyRes, err := cc.cu.UpdateCompany(company, uint(userId.(float64)), uint(companyId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, companyRes)
}

func (cc *companyController) DeleteCompany(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("companyId")
	companyId, _ := strconv.Atoi(id)

	if err := cc.cu.DeleteCompany(uint(userId.(float64)), uint(companyId)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}