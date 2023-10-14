package repository

import (
	"errors"
	"fmt"
	"recruit-info-service/model"

	"gorm.io/gorm"
)

type ILikeRepository interface {
	CheckLikeByCompanyId(userId uint, companyId uint) (bool, error)
	CreateLike(like *model.Like) error
	DeleteLike(userId uint, companyId uint) error
	CountLike(companyId uint) (int, error)
}

type LikeRepository struct {
	db *gorm.DB
}

func NewLikeRepository(db *gorm.DB) ILikeRepository {
	return &LikeRepository{db}
}

func (lr *LikeRepository) CheckLikeByCompanyId(userId uint, companyId uint) (bool, error) {
	if err := lr.db.Where("user_id = ? AND company_id = ?", userId, companyId).First(&model.Like{}).Error; err != nil {
		// レコードが見つからない場合は、falseを返す
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
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
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("record not found")
	}
	return nil
}

func (lr *LikeRepository) CountLike(companyId uint) (int, error) {
	var count int64
	if err := lr.db.Model(&model.Like{}).Where("company_id=?", companyId).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}