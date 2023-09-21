package controller

import (
	"net/http"
	"recruit-info-service/model"
	"recruit-info-service/usecase"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type ICommentController interface {
	CreateComment(c echo.Context) error
	DeleteComment(c echo.Context) error
}

type CommentController struct {
	cu usecase.ICommentUsecase
}

func NewCommentController(cu usecase.ICommentUsecase) ICommentController {
	return &CommentController{cu}
}

func (cmc *CommentController) CreateComment(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	var reqBody struct {
		Content string `json:"content"`
		CompanyID uint `json:"companyId"`
	}

	if err := c.Bind(&reqBody); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	comment := model.Comment{
		Content: reqBody.Content,
		UserID: uint(userId.(float64)),
		CompanyID:    reqBody.CompanyID,
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

	var reqBody struct {
		CompanyID uint `json:"companyId"`
	}

	if err := cmc.cu.DeleteComment(uint(userId.(float64)), reqBody.CompanyID); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}