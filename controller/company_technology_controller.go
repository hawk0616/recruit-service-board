package controller

import (
	"net/http"
	"recruit-info-service/model"
	"recruit-info-service/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ICompanyTechnologyController interface {
	GetCompanyTechnologyByCompanyId(c echo.Context) error
	CreateCompanyTechnology(c echo.Context) error
	DeleteCompanyTechnology(c echo.Context) error
}

type CompanyTechnologyController struct {
	ctu usecase.ICompanyTechnologyUsecase
}

func NewCompanyTechnologyController(ctu usecase.ICompanyTechnologyUsecase) ICompanyTechnologyController {
	return &CompanyTechnologyController{ctu}
}

func (ctc *CompanyTechnologyController) GetCompanyTechnologyByCompanyId(c echo.Context) error {
	tempCompanyId := c.Param("companyId")
	companyId, err := strconv.Atoi(tempCompanyId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid company ID")
	}

	companyTechnologies, err := ctc.ctu.GetCompanyTechnologyByCompanyId(uint(companyId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, companyTechnologies)
}

func (ctc *CompanyTechnologyController) CreateCompanyTechnology(c echo.Context) error {
	tempCompanyId := c.Param("companyId")
	companyId, err := strconv.Atoi(tempCompanyId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid company ID")
	}

	var reqBody struct {
		TechnologyID uint `json:"technologyId"`
	}

	if err := c.Bind(&reqBody); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	companyTechnology := model.CompanyTechnology{
		CompanyID:    uint(companyId),
		TechnologyID: reqBody.TechnologyID,
	}

	companyTechnologyRes, err := ctc.ctu.CreateCompanyTechnology(companyTechnology)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, companyTechnologyRes)
}

func (ctc *CompanyTechnologyController) DeleteCompanyTechnology(c echo.Context) error {
	temCompanyId := c.Param("companyId")
	companyId, _ := strconv.Atoi(temCompanyId)
	temTecnologyId := c.Param("tecnologyId")
	tecnologyId, _ := strconv.Atoi(temTecnologyId)

	if err := ctc.ctu.DeleteCompanyTechnology(uint(companyId), uint(tecnologyId)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}