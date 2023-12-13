package repository

import (
	"context"
	"github.com/Hamid-Rezaei/goBasket/internal/domain/model"
	"gorm.io/gorm"
	"time"
)

type BasketDTO struct {
	model.Basket
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(ctx context.Context, model model.Basket) error {
	tx := r.db.WithContext(ctx).Begin()

	if err := tx.Create(&BasketDTO{Basket: model}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *Repository) Update(ctx context.Context, model model.Basket) error {
	tx := r.db.WithContext(ctx).Begin()

	if err := tx.Model(&BasketDTO{}).Where("id = ?", model.ID).Updates(&BasketDTO{Basket: model}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *Repository) Delete(ctx context.Context, model model.Basket, id int) error {
	tx := r.db.WithContext(ctx).Begin()

	if err := tx.Delete(&BasketDTO{}, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *Repository) GetBaskets(_ context.Context) ([]model.Basket, error) {
	var basketDTOs []BasketDTO

	if err := r.db.Find(&basketDTOs).Error; err != nil {
		return nil, err
	}

	baskets := make([]model.Basket, len(basketDTOs))
	for index, dto := range basketDTOs {
		baskets[index] = dto.Basket
	}

	return baskets, nil
}

func (r *Repository) GetBasketByID(_ context.Context, id int) (*model.Basket, error {
	var basketDTO BasketDTO

	if err := r.db.First(&basketDTO, id).Error; err != nil {
		return nil, err
	}

	return &basketDTO.Basket, nil
}
