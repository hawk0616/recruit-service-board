package usecase

import (
	"recruit-info-service/model"
	"recruit-info-service/repository"
)

type ILikeUsecase interface {
	CreateLike(like model.Like) (model.LikeResponse, error)
	DeleteLike(userId uint, companyId uint) error
}

type LikeUsecase struct {
	lr repository.ILikeRepository
}

func NewLikeUsecase(lr repository.ILikeRepository) ILikeUsecase {
	return &LikeUsecase{lr}
}

func (lu *LikeUsecase) CreateLike(like model.Like) (model.LikeResponse, error) {
	if err := lu.lr.CreateLike(&like); err != nil {
		return model.LikeResponse{}, err
	}
	resLike := model.LikeResponse{
		UserID: like.UserID,
		CompanyID: like.CompanyID,
	}
	return resLike, nil
}

func (lu *LikeUsecase) DeleteLike(userId uint, companyId uint) error {
	if err := lu.lr.DeleteLike(userId, companyId); err != nil {
		return err
	}
	return nil
}