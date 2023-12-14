package handler

import "github.com/Hamid-Rezaei/goBasket/internal/infra/repository"

type Handler struct {
	userRepo   repository.UserRepo
	basketRepo repository.BasketRepo
}

func NewHandler(ur repository.UserRepo, br repository.BasketRepo) *Handler {
	return &Handler{
		userRepo:   ur,
		basketRepo: br,
	}
}
