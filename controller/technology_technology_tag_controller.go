package controller

import (
	"net/http"
	"recruit-info-service/model"
	"recruit-info-service/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ITechnologyTechnologyTagController interface {
	CreateTechnologyTechnologyTag(c echo.Context) error
	DeleteTechnologyTechnologyTag(c echo.Context) error
}

type TechnologyTechnologyTagController struct {
	tttu usecase.ITechnologyTechnologyTagUsecase
}

func NewTechnologyTechnologyTagController(tttu usecase.ITechnologyTechnologyTagUsecase) ITechnologyTechnologyTagController {
	return &TechnologyTechnologyTagController{tttu}
}

func (tttc *TechnologyTechnologyTagController) CreateTechnologyTechnologyTag(c echo.Context) error {
	tempTechnologyId := c.Param("technologyId")
	technologyId, err := strconv.Atoi(tempTechnologyId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid company ID")
	}

	var reqBody struct {
		TechnologyTagID uint `json:"technologyTagId"`
	}

	if err := c.Bind(&reqBody); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	companyTechnology := model.TechnologyTechnologyTag{
		TechnologyID:    uint(technologyId),
		TechnologyTagID: reqBody.TechnologyTagID,
	}

	TechnologyTechnologyTagRes, err := tttc.tttu.CreateTechnologyTechnologyTag(companyTechnology)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, TechnologyTechnologyTagRes)
}


func (tttc *TechnologyTechnologyTagController) DeleteTechnologyTechnologyTag(c echo.Context) error {
	temTecnologyId := c.Param("tecnologyId")
	tecnologyId, _ := strconv.Atoi(temTecnologyId)
	temTechnologyTagId := c.Param("technologyTagId")
	technologyTagId, _ := strconv.Atoi(temTechnologyTagId)

	if err := tttc.tttu.DeleteTechnologyTechnologyTag(uint(tecnologyId), uint(technologyTagId)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}