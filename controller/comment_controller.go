package controller

import (
	"net/http"
	"recruit-info-service/model"
	"recruit-info-service/usecase"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type ICommentController interface {
	GetCommentsByCompanyId(c echo.Context) error
	CreateComment(c echo.Context) error
	DeleteComment(c echo.Context) error
	CountComment(c echo.Context) error
}

type CommentController struct {
	cu usecase.ICommentUsecase
}

func NewCommentController(cu usecase.ICommentUsecase) ICommentController {
	return &CommentController{cu}
}

func (cmc *CommentController) GetCommentsByCompanyId(c echo.Context) error {
	temCompanyId := c.Param("companyId")
	companyId, _ := strconv.Atoi(temCompanyId)

	comments, err := cmc.cu.GetCommentsByCompanyId(uint(companyId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, comments)
}

func (cmc *CommentController) CreateComment(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	var reqBody struct {
		Content   string `json:"content"`
		CompanyID uint   `json:"companyId"`
	}

	if err := c.Bind(&reqBody); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	comment := model.Comment{
		Content:   reqBody.Content,
		UserID:    uint(userId.(float64)),
		CompanyID: reqBody.CompanyID,
	}

	commentRes, err := cmc.cu.CreateComment(comment)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, commentRes)
}

func (cmc *CommentController) DeleteComment(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	temCompanyId := c.Param("companyId")
	companyId, err := strconv.Atoi(temCompanyId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid companyId")
	}

	if err := cmc.cu.DeleteComment(uint(userId.(float64)), uint(companyId)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func (cmc *CommentController) CountComment(c echo.Context) error {
	temCompanyId := c.Param("companyId")
	companyId, _ := strconv.Atoi(temCompanyId)

	count, err := cmc.cu.CountComment(uint(companyId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, count)
}
