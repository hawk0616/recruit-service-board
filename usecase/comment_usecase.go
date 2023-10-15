package usecase

import (
	"recruit-info-service/model"
	"recruit-info-service/repository"
)

type ICommentUsecase interface {
	GetCommentsByCompanyId(companyId uint) ([]model.CommentResponse, error)
	CreateComment(comment model.Comment) (model.CommentResponse, error)
	DeleteComment(userId uint, companyId uint) error
	CountComment(companyId uint) (int, error)
}

type CommentUsecase struct {
	cr repository.ICommentRepository
}

func NewCommentUsecase(cr repository.ICommentRepository) ICommentUsecase {
	return &CommentUsecase{cr}
}

func (cu *CommentUsecase) GetCommentsByCompanyId(companyId uint) ([]model.CommentResponse, error) {
	comments, err := cu.cr.GetCommentsByCompanyId(companyId)
	if err != nil {
		return nil, err
	}
	resComments := []model.CommentResponse{}
	for _, v := range comments {
		c := model.CommentResponse{
			Content: v.Content,
			UserID: v.UserID,
			CompanyID: v.CompanyID,
			CreatedAt: v.CreatedAt,
		}
		resComments = append(resComments, c)
	}
	return resComments, nil
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

func (cu *CommentUsecase) CountComment(companyId uint) (int, error) {
	count, err := cu.cr.CountComment(companyId)
	if err != nil {
		return 0, err
	}
	return count, nil
}