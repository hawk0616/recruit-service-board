package usecase

import (
	"recruit-info-service/model"
	"recruit-info-service/repository"
)

type ICommentUsecase interface {
	CreateComment(comment model.Comment) (model.CommentResponse, error)
	DeleteComment(userId uint, companyId uint) error
}

type CommentUsecase struct {
	cr repository.ICommentRepository
}

func NewCommentUsecase(cr repository.ICommentRepository) ICommentUsecase {
	return &CommentUsecase{cr}
}

func (cu *CommentUsecase) CreateComment(comment model.Comment) (model.CommentResponse, error) {
	if err := cu.cr.CreateComment(&comment); err != nil {
		return model.CommentResponse{}, err
	}
	resComment := model.CommentResponse{
		Content: comment.Content,
		UserID: comment.UserID,
		CompanyID: comment.CompanyID,
	}
	return resComment, nil
}

func (cu *CommentUsecase) DeleteComment(userId uint, companyId uint) error {
	if err := cu.cr.DeleteComment(userId, companyId); err != nil {
		return err
	}
	return nil
}