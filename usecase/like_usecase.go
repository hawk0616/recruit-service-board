package usecase

import (
	"recruit-info-service/model"
	"recruit-info-service/repository"
)

type ILikeUsecase interface {
	CheckLikeByCompanyId(userId uint, companyId uint) (bool, error)
	CreateLike(like model.Like) (model.LikeResponse, error)
	DeleteLike(userId uint, companyId uint) error
	CountLike(companyId uint) (int, error)
}

type LikeUsecase struct {
	lr repository.ILikeRepository
}

func NewLikeUsecase(lr repository.ILikeRepository) ILikeUsecase {
	return &LikeUsecase{lr}
}

func (lu *LikeUsecase) CheckLikeByCompanyId(userId uint, companyId uint) (bool, error) {
	isLike, err := lu.lr.CheckLikeByCompanyId(userId, companyId)
	if err != nil {
		return false, err
	}
	return isLike, nil
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

func (lu *LikeUsecase) CountLike(companyId uint) (int, error) {
	count, err := lu.lr.CountLike(companyId)
	if err != nil {
		return 0, err
	}
	return count, nil
}