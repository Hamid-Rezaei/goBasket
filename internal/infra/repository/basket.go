package repository

import (
	"context"
	"github.com/Hamid-Rezaei/goBasket/internal/domain/model"
	"gorm.io/gorm"
)

type BasketDTO struct {
	gorm.Model
	model.Basket
}

type BasketRepository struct {
	db *gorm.DB
}

func NewBasketRepo(db *gorm.DB) *BasketRepository {
	return &BasketRepository{
		db: db,
	}
}

func (r *BasketRepository) Create(ctx context.Context, model model.Basket) (uint, error) {
	tx := r.db.WithContext(ctx).Begin()

	basketDTO := BasketDTO{Basket: model}
	if err := tx.Create(&basketDTO).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	return basketDTO.ID, tx.Commit().Error
}

func (r *BasketRepository) Update(ctx context.Context, model model.Basket, id int) error {
	tx := r.db.WithContext(ctx).Begin()

	if err := tx.Model(&BasketDTO{}).Where("id = ?", id).Updates(&BasketDTO{Basket: model}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *BasketRepository) Delete(ctx context.Context, id int) error {
	tx := r.db.WithContext(ctx).Begin()

	if err := tx.Delete(&BasketDTO{}, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *BasketRepository) GetBaskets(_ context.Context) ([]model.Basket, error) {
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

func (r *BasketRepository) GetBasketByID(_ context.Context, id int) (*model.Basket, error) {
	var basketDTO BasketDTO

	if err := r.db.First(&basketDTO, id).Error; err != nil {
		return nil, err
	}

	return &basketDTO.Basket, nil
}
