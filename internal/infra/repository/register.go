package repository

import (
	"context"
	"github.com/Hamid-Rezaei/goBasket/internal/domain/model"
)

type BasketRepo interface {
	Create(ctx context.Context, model model.Basket) (uint, error)
	Update(ctx context.Context, model model.Basket, id int) error
	Delete(ctx context.Context, id int) error
	GetUserBaskets(_ context.Context, userID uint) ([]model.Basket, error)
	GetUserBasketByID(_ context.Context, userID uint, basketID uint) (*model.Basket, error)
}

type UserRepo interface {
	Create(ctx context.Context, model model.User) (uint, error)
	GetByEmail(_ context.Context, password string) (*UserDTO, error)
	GetUserByID(_ context.Context, id uint) (*UserDTO, error)
}
