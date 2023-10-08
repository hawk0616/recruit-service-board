package repository

import (
	"fmt"
	"recruit-info-service/model"

	"gorm.io/gorm"
)

type ILikeRepository interface {
	CreateLike(like *model.Like) error
	DeleteLike(userId uint, companyId uint) error
	CountLike() (int, error)
}

type LikeRepository struct {
	db *gorm.DB
}

func NewLikeRepository(db *gorm.DB) ILikeRepository {
	return &LikeRepository{db}
}

func (lr *LikeRepository) CreateLike(like *model.Like) error {
	if err := lr.db.Create(like).Error; err != nil {
		return err
	}
	return nil
}

func (lr *LikeRepository) DeleteLike(userId uint, companyId uint) error {
	result := lr.db.Where("user_id = ? AND company_id = ?", userId, companyId).Delete(&model.Like{})
	if result.Error != nil {
		fmt.Println("aaa")
		return result.Error
	}
	if result.RowsAffected < 0 {
		fmt.Println("bbb")
		return fmt.Errorf("record not found")
	}
	fmt.Println("ccc")
	return nil
}

func (lr *LikeRepository) CountLike() (int, error) {
	var count int64
	if err := lr.db.Model(&model.Like{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}