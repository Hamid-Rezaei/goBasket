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

func (ur *UserRepository) GetByPassword(_ context.Context, password string) (*model.User, error) {
	var user model.User
	if err := ur.db.Where(&model.User{Password: password}).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
