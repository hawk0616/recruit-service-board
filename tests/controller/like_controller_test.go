package controller_test

import (
	"net/http"
	"net/http/httptest"
	"recruit-info-service/controller"
	"recruit-info-service/model"
	"strings"
	"testing"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockLikeUsecase struct {
	mock.Mock
}

func (m *MockLikeUsecase) CheckLikeByCompanyId(userId uint, companyId uint) (bool, error) {
	args := m.Called(userId, companyId)
	return args.Bool(0), args.Error(1)
}

func (m *MockLikeUsecase) CreateLike(like model.Like) (model.LikeResponse, error) {
	args := m.Called(like)
	return args.Get(0).(model.LikeResponse), args.Error(1)
}

func (m *MockLikeUsecase) DeleteLike(userId uint, companyId uint) error {
	args := m.Called(userId, companyId)
	return args.Error(0)
}

func (m *MockLikeUsecase) CountLike(companyId uint) (int, error) {
	args := m.Called(companyId)
	return args.Int(0), args.Error(1)
}

func TestCheckLikeByCompanyId(t *testing.T) {
	mockUsecase := &MockLikeUsecase{}
	expectLikeRes := true
	mockUsecase.On("CheckLikeByCompanyId", uint(1), uint(1)).Return(expectLikeRes, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/companies/likes/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/companies/likes/1")
	c.SetParamNames("companyId")
	c.SetParamValues("1")

	// JWTのユーザー情報をセットアップ
	userToken := &jwt.Token{
		Claims: jwt.MapClaims{
			"user_id": float64(1),
		},
	}
	c.Set("user", userToken)

	controller := controller.NewLikeController(mockUsecase)
	err := controller.CheckLikeByCompanyId(c)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, `{"liked":true}`, strings.TrimSpace(rec.Body.String()))
}

func TestCreateLike(t *testing.T) {
	mockUsecase := &MockLikeUsecase{}
	expectLikeRes := model.LikeResponse{
		UserID:    1,
		CompanyID: 1,
	}
	mockUsecase.On("CreateLike", model.Like{
		UserID:    1,
		CompanyID: 1,
	}).Return(expectLikeRes, nil)

	e := echo.New()
	reqBody := `{"companyId": 1}`
	req := httptest.NewRequest(http.MethodPost, "/companies/likes", strings.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// JWTのユーザー情報をセットアップ
	userToken := &jwt.Token{
		Claims: jwt.MapClaims{
			"user_id": float64(1),
		},
	}
	c.Set("user", userToken)

	controller := controller.NewLikeController(mockUsecase)
	err := controller.CreateLike(c)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, `{"user_id":1,"company_id":1}`, strings.TrimSpace(rec.Body.String()))
}

func TestDeleteLike(t *testing.T) {
	mockUsecase := &MockLikeUsecase{}
	mockUsecase.On("DeleteLike", uint(1), uint(1)).Return(nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/companies/likes/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/companies/likes/1")
	c.SetParamNames("companyId")
	c.SetParamValues("1")

	// JWTのユーザー情報をセットアップ
	userToken := &jwt.Token{
		Claims: jwt.MapClaims{
			"user_id": float64(1),
		},
	}
	c.Set("user", userToken)

	controller := controller.NewLikeController(mockUsecase)
	err := controller.DeleteLike(c)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "", strings.TrimSpace(rec.Body.String()))
}

func TestCountLike(t *testing.T) {
	mockUsecase := &MockLikeUsecase{}
	expectLikeRes := 1
	mockUsecase.On("CountLike", uint(1)).Return(expectLikeRes, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/companies/likes/1/count", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/companies/likes/1/count")
	c.SetParamNames("companyId")
	c.SetParamValues("1")

	controller := controller.NewLikeController(mockUsecase)
	err := controller.CountLike(c)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, `1`, strings.TrimSpace(rec.Body.String()))
}