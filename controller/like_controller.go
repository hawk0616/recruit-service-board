package controller

import (
	"net/http"
	"recruit-info-service/model"
	"recruit-info-service/usecase"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type ILikeController interface {
	CheckLikeByCompanyId(c echo.Context) error
	CreateLike(c echo.Context) error
	DeleteLike(c echo.Context) error
	CountLike(c echo.Context) error
}

type LikeController struct {
	lu usecase.ILikeUsecase
}

func NewLikeController(lu usecase.ILikeUsecase) ILikeController {
	return &LikeController{lu}
}

func (lc *LikeController) CheckLikeByCompanyId(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	temCompanyId := c.Param("companyId")
	companyId, _ := strconv.Atoi(temCompanyId)

	isLike, err := lc.lu.CheckLikeByCompanyId(uint(userId.(float64)), uint(companyId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, map[string]bool{"liked": isLike})
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

	temCompanyId := c.Param("companyId")
	companyId, err := strconv.Atoi(temCompanyId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid companyId")
	}

	if err := lc.lu.DeleteLike(uint(userId.(float64)), uint(companyId)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func (lc *LikeController) CountLike(c echo.Context) error {
	temCompanyId := c.Param("companyId")
	companyId, _ := strconv.Atoi(temCompanyId)
	
	count, err := lc.lu.CountLike(uint(companyId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, count)
}