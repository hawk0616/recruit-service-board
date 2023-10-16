package repository_test

import (
	"recruit-info-service/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type DBHandler interface {
	CheckLikeByCompanyId(companyId string) (bool, error)
	CreateLike(like *model.Like) error
	DeleteLike(userId uint, companyId uint) error
	CountLike(companyId uint) (int, error)
}

type RealLikeDBHandler struct {
	db *gorm.DB
}

func NewRealLikeDBHandler() *RealLikeDBHandler {
	return &RealLikeDBHandler{}
}

type MockLikeDBHandler struct {
	mock.Mock
}

func NewMockLikeDBHandler() *MockLikeDBHandler {
	return &MockLikeDBHandler{}
}

func (m *MockLikeDBHandler) CheckLikeByCompanyId(userId uint, companyId uint) (bool, error) {
	args := m.Called(companyId)
	return args.Bool(0), args.Error(1)
}

func (m *MockLikeDBHandler) CreateLike(like *model.Like) error {
	args := m.Called(like)
	return args.Error(0)
}

func (m *MockLikeDBHandler) DeleteLike(userId uint, companyId uint) error {
	args := m.Called(userId, companyId)
	return args.Error(0)
}

func (m *MockLikeDBHandler) CountLike(companyId uint) (int, error) {
	args := m.Called(companyId)
	return args.Int(0), args.Error(1)
}

func TestCheckLikeByCompanyId(t *testing.T) {
	mockDB := NewMockLikeDBHandler()

	expectLikeRes := true
	mockDB.On("CheckLikeByCompanyId", uint(1)).Return(expectLikeRes, nil)

	isLike, err := mockDB.CheckLikeByCompanyId(uint(1), uint(1))
	assert.NoError(t, err)
	assert.Equal(t, expectLikeRes, isLike)
}

func TestCreateLike(t *testing.T) {
	mockDB := NewMockLikeDBHandler()

	expectLike := model.Like{
		UserID:    1,
		CompanyID: 1,
	}
	mockDB.On("CreateLike", &expectLike).Return(nil)

	err := mockDB.CreateLike(&expectLike)
	assert.NoError(t, err)
}

func TestDeleteLike(t *testing.T) {
	mockDB := NewMockLikeDBHandler()

	mockDB.On("DeleteLike", uint(1), uint(1)).Return(nil)

	err := mockDB.DeleteLike(uint(1), uint(1))
	assert.NoError(t, err)
}

func TestCountLike(t *testing.T) {
	mockDB := NewMockLikeDBHandler()

	expectCount := 1
	mockDB.On("CountLike", uint(1)).Return(expectCount, nil)

	count, err := mockDB.CountLike(uint(1))
	assert.NoError(t, err)
	assert.Equal(t, expectCount, count)
}