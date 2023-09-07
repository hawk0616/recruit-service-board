package controller

import (
	"net/http"
	"recruit-info-service/model"
	"recruit-info-service/usecase"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type ITechnologyTagController interface {
	GetAllTechnologyTags(c echo.Context) error
	GetTechnologyTagById(c echo.Context) error
	CreateTechnologyTag(c echo.Context) error
	UpdateTechnologyTag(c echo.Context) error
	DeleteTechnologyTag(c echo.Context) error
}

type technologyTagController struct {
	ttu usecase.ITechnologyTagUsecase
}

func NewTechnologyTagController(ttu usecase.ITechnologyTagUsecase) ITechnologyTagController {
	return &technologyTagController{ttu}
}

func (ttc *technologyTagController) GetAllTechnologyTags(c echo.Context) error {
	technologyTagsRes, err := ttc.ttu.GetAllTechnologyTags()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, technologyTagsRes)
}

func (ttc *technologyTagController) GetTechnologyTagById(c echo.Context) error {
	id := c.Param("technologyTagId")
	technologyTagId, _ := strconv.Atoi(id)
	technologyTagRes, err := ttc.ttu.GetTechnologyTagById(uint(technologyTagId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, technologyTagRes)
}

func (ttc *technologyTagController) CreateTechnologyTag(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	technologyTag := model.TechnologyTag{}
	if err := c.Bind(&technologyTag); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	technologyTag.UserId = uint(userId.(float64))
	technologyRes, err := ttc.ttu.CreateTechnologyTag(technologyTag)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, technologyRes)
}

func (ttc *technologyTagController) UpdateTechnologyTag(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("technologyId")
	technologyTagId, _ := strconv.Atoi(id)

	technologyTag := model.TechnologyTag{}
	if err := c.Bind(&technologyTag); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	technologyTagRes, err := ttc.ttu.UpdateTechnologyTag(technologyTag, uint(userId.(float64)), uint(technologyTagId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, technologyTagRes)
}

func (ttc *technologyTagController) DeleteTechnologyTag(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("technologyTagId")
	technologyTagId, _ := strconv.Atoi(id)

	if err := ttc.ttu.DeleteTechnologyTag(uint(userId.(float64)), uint(technologyTagId)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}