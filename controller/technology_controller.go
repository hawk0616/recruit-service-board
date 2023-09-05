package controller

import (
	"net/http"
	"recruit-info-service/model"
	"recruit-info-service/usecase"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type ITechnologyController interface {
	GetAllTechnologies(c echo.Context) error
	GetTechnologyById(c echo.Context) error
	CreateTechnology(c echo.Context) error
	UpdateTechnology(c echo.Context) error
	DeleteTechnology(c echo.Context) error
}

type technologyController struct {
	tu usecase.ITechnologyUsecase
}

func NewTechnologyController(tu usecase.ITechnologyUsecase) ITechnologyController {
	return &technologyController{tu}
}

func (tc *technologyController) GetAllTechnologies(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	technologiesRes, err := tc.tu.GetAllTechnologies(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, technologiesRes)
}

func (tc *technologyController) GetTechnologyById(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("technologyId")
	technologyId, _ := strconv.Atoi(id)
	technologyRes, err := tc.tu.GetTechnologyById(uint(userId.(float64)), uint(technologyId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, technologyRes)
}

func (tc *technologyController) CreateTechnology(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	technology := model.Technology{}
	if err := c.Bind(&technology); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	technology.UserId = uint(userId.(float64))
	technologyRes, err := tc.tu.CreateTechnology(technology)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, technologyRes)
}

func (tc *technologyController) UpdateTechnology(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("technologyId")
	technologyId, _ := strconv.Atoi(id)

	technology := model.Technology{}
	if err := c.Bind(&technology); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	technologyRes, err := tc.tu.UpdateTechnology(technology, uint(userId.(float64)), uint(technologyId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, technologyRes)
}

func (tc *technologyController) DeleteTechnology(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("technologyId")
	technologyId, _ := strconv.Atoi(id)

	if err := tc.tu.DeleteTechnology(uint(userId.(float64)), uint(technologyId)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}