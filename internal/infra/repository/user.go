package repository

import (
	"context"
	"errors"
	"github.com/Hamid-Rezaei/goBasket/internal/domain/model"
	"gorm.io/gorm"
)

type UserDTO struct {
	gorm.Model
	model.User
	Baskets []BasketDTO `gorm:"foreignkey:UserID" json:"baskets,omitempty"`
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Create(ctx context.Context, model model.User) (uint, error) {
	tx := ur.db.WithContext(ctx).Begin()

	userDTO := UserDTO{User: model}
	if err := tx.Create(&userDTO).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	return userDTO.ID, tx.Commit().Error
}

func (ur *UserRepository) GetByEmail(_ context.Context, email string) (*UserDTO, error) {
	var userDTO UserDTO
	if err := ur.db.Where(&model.User{Email: email}).First(&userDTO).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &userDTO, nil
}

func (ur *UserRepository) GetUserByID(_ context.Context, id uint) (*UserDTO, error) {
	var userDTO UserDTO

	if err := ur.db.First(&userDTO, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &userDTO, nil
}
