package controller

import (
	"net/http"
	"recruit-info-service/model"
	"recruit-info-service/usecase"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type ILikeController interface {
	CreateLike(c echo.Context) error
	DeleteLike(c echo.Context) error
}

type LikeController struct {
	lu usecase.ILikeUsecase
}

func NewLikeController(lu usecase.ILikeUsecase) ILikeController {
	return &LikeController{lu}
}

func (lc *LikeController) CreateLike(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	var reqBody struct {
		CompanyID uint `json:"companyId"`
	}

	if err := c.Bind(&reqBody); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	like := model.Like{
		UserID: uint(userId.(float64)),
		CompanyID:    reqBody.CompanyID,
	}

	likeRes, err := lc.lu.CreateLike(like)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, likeRes)
}

func (lc *LikeController) DeleteLike(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	var reqBody struct {
		CompanyID uint `json:"companyId"`
	}

	if err := lc.lu.DeleteLike(uint(userId.(float64)), reqBody.CompanyID); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}