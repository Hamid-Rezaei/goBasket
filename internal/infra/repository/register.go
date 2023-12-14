package repository

import (
	"context"
	"github.com/Hamid-Rezaei/goBasket/internal/domain/model"
)

type BasketRepo interface {
	Create(ctx context.Context, model model.Basket) (uint, error)
	Update(ctx context.Context, model model.Basket, id int) error
	Delete(ctx context.Context, id int) error
	GetBaskets(_ context.Context) ([]model.Basket, error)
	GetBasketByID(_ context.Context, id int) (*model.Basket, error)
}

type UserRepo interface {
	Create(ctx context.Context, model model.User) (uint, error)
	GetByPassword(_ context.Context, password string) (*model.User, error)
}
