package repository

import (
	"fmt"
	"recruit-info-service/model"

	"gorm.io/gorm"
)

type ICommentRepository interface {
	CreateComment(comment *model.Comment) error
	DeleteComment(userId uint, companyId uint) error
	CountComment(companyId uint) (int, error)
}

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) ICommentRepository {
	return &CommentRepository{db}
}

func (cr *CommentRepository) CreateComment(comment *model.Comment) error {
	if err := cr.db.Create(comment).Error; err != nil {
		return err
	}
	return nil
}

func (cr *CommentRepository) DeleteComment(userId uint, companyId uint) error {
	result := cr.db.Where("user_id = ? AND company_id = ?", userId, companyId).Delete(&model.Comment{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("record not found")
	}
	return nil
}

func (cr *CommentRepository) CountComment(companyId uint) (int, error) {
	var count int64
	if err := cr.db.Model(&model.Comment{}).Where("company_id = ?", companyId).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}