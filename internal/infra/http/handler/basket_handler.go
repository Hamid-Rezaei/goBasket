package handler

import (
	"github.com/Hamid-Rezaei/goBasket/internal/domain/model"
	"github.com/Hamid-Rezaei/goBasket/internal/infra/http/request"
	"github.com/Hamid-Rezaei/goBasket/internal/infra/repository"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Basket struct {
	repo repository.BasketRepo
}

func NewBasket(repo repository.BasketRepo) *Basket {
	return &Basket{
		repo: repo,
	}
}

func (b *Basket) Create(c echo.Context) error {
	var req request.BasketCreate

	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}
	// Validate Request
	if err := req.Validate(); err != nil {
		return echo.ErrBadRequest
	}

	id, err := b.repo.Create(c.Request().Context(), model.Basket{
		Data:  req.Data,
		State: req.State,
	})

	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusCreated, id)
}

func (b *Basket) Register(g *echo.Group) {
	//	g.GET("", b)
	g.POST("/", b.Create)
	//g.GET(":id", s.GetByID)
}
