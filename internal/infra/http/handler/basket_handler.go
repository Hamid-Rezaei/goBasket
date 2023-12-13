package handler

import (
	"github.com/Hamid-Rezaei/goBasket/internal/domain/model"
	"github.com/Hamid-Rezaei/goBasket/internal/infra/http/request"
	"github.com/Hamid-Rezaei/goBasket/internal/infra/repository"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
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

	// Bind Request
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

func (b *Basket) GetBaskets(c echo.Context) error {
	var req request.BasketCreate

	// Bind Request
	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}
	// Validate Request
	if err := req.Validate(); err != nil {
		return echo.ErrBadRequest
	}

	baskets, err := b.repo.GetBaskets(c.Request().Context())

	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusCreated, baskets)

}

func (b *Basket) GetByID(c echo.Context) error {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.ErrBadRequest
	}

	var req request.BasketCreate

	// Bind Request
	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}
	// Validate Request
	if err := req.Validate(); err != nil {
		return echo.ErrBadRequest
	}

	basket, err := b.repo.GetBasketByID(c.Request().Context(), int(id))

	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusCreated, basket)

}

func (b *Basket) Register(g *echo.Group) {
	g.GET("/", b.GetBaskets)
	g.POST("/", b.Create)
	g.GET("/:id", b.GetByID)
}
