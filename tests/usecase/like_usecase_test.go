package usecase_test

import (
    "testing"
    "recruit-info-service/model"
    "recruit-info-service/usecase"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

type MockLikeRepository struct {
	mock.Mock
}

func (m *MockLikeRepository) CheckLikeByCompanyId(userId uint, companyId uint) (bool, error) {
	args := m.Called(companyId)
	return args.Bool(0), args.Error(1)
}

func (m *MockLikeRepository) CreateLike(like *model.Like) error {
	args := m.Called(like)
	return args.Error(0)
}

func (m *MockLikeRepository) DeleteLike(userId uint, companyId uint) error {
	args := m.Called(userId, companyId)
	return args.Error(0)
}

func (m *MockLikeRepository) CountLike(companyId uint) (int, error) {
	args := m.Called(companyId)
	return args.Int(0), args.Error(1)
}

func TestCheckLikeByCompanyId(t *testing.T) {
	mockRepo := &MockLikeRepository{}
	mockUsecase := usecase.NewLikeUsecase(mockRepo)
	expectLikeRes := true
	mockRepo.On("CheckLikeByCompanyId", uint(1)).Return(expectLikeRes, nil)

	isLike, err := mockUsecase.CheckLikeByCompanyId(uint(1), uint(1))
	assert.NoError(t, err)
	assert.Equal(t, expectLikeRes, isLike)
}

func TestCreateLike(t *testing.T) {
	mockRepo := &MockLikeRepository{}
	mockUsecase := usecase.NewLikeUsecase(mockRepo)
	expectLike := model.Like{
		UserID:    1,
		CompanyID: 1,
	}
	expectLikeRes := model.LikeResponse{
		UserID:    1,
		CompanyID: 1,
	}
	mockRepo.On("CreateLike", &expectLike).Return(nil)

	resLike, err := mockUsecase.CreateLike(expectLike)
	assert.NoError(t, err)
	assert.Equal(t, expectLikeRes, resLike)
}

func TestDeleteLike(t *testing.T) {
	mockRepo := &MockLikeRepository{}
	mockUsecase := usecase.NewLikeUsecase(mockRepo)
	mockRepo.On("DeleteLike", uint(1), uint(1)).Return(nil)

	err := mockUsecase.DeleteLike(uint(1), uint(1))
	assert.NoError(t, err)
}

func TestCountLike(t *testing.T) {
	mockRepo := &MockLikeRepository{}
	mockUsecase := usecase.NewLikeUsecase(mockRepo)
	expectCount := 1
	mockRepo.On("CountLike", uint(1)).Return(expectCount, nil)

	count, err := mockUsecase.CountLike(uint(1))
	assert.NoError(t, err)
	assert.Equal(t, expectCount, count)
}