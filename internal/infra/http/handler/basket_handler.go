package handler

import (
	"errors"
	"github.com/Hamid-Rezaei/goBasket/internal/domain/model"
	"github.com/Hamid-Rezaei/goBasket/internal/infra/http/request"
	"github.com/Hamid-Rezaei/goBasket/internal/infra/repository"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"log"
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
		log.Printf("%v\n", err)
		return echo.ErrBadRequest
	}
	// Validate Request
	if err := req.Validate(); err != nil {
		log.Printf("%v\n", err)
		return echo.ErrBadRequest
	}

	id, err := b.repo.Create(c.Request().Context(), model.Basket{
		Data:  req.Data,
		State: req.State,
	})

	if err != nil {
		log.Printf("%v\n", err)
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusCreated, id)
}

func (b *Basket) Update(c echo.Context) error {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Printf("%v\n", err)
		return echo.ErrBadRequest
	}

	var req request.BasketUpdate

	// Bind Request
	if err := c.Bind(&req); err != nil {
		log.Printf("%v\n", err)
		return echo.ErrBadRequest
	}
	// Validate Request
	if err := req.Validate(); err != nil {
		log.Printf("%v\n", err)
		return echo.ErrBadRequest
	}

	basket, err := b.repo.GetBasketByID(c.Request().Context(), int(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNoContent, "Basket Not Found!")
		}

		return echo.ErrInternalServerError
	}

	if basket.State == "COMPLETED" {
		return c.JSON(http.StatusNotAcceptable, "Cannot Update Completed Basket!")
	}

	basketModel := model.Basket{
		Data:  req.Data,
		State: req.State,
	}

	if err := b.repo.Update(c.Request().Context(), basketModel, int(id)); err != nil {
		log.Printf("%v\n", err)
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, "Basket Was Updated Successfully.")

}

func (b *Basket) GetBaskets(c echo.Context) error {

	baskets, err := b.repo.GetBaskets(c.Request().Context())

	if err != nil {
		log.Printf("%v\n", err)
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, baskets)

}

func (b *Basket) GetByID(c echo.Context) error {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Printf("%v\n", err)
		return echo.ErrBadRequest
	}

	basket, err := b.repo.GetBasketByID(c.Request().Context(), int(id))

	if err != nil {
		log.Printf("%v\n", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNoContent, "Basket Not Found!")
		}

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, basket)

}

func (b *Basket) Delete(c echo.Context) error {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Printf("%v\n", err)
		return echo.ErrBadRequest
	}

	if err := b.repo.Delete(c.Request().Context(), int(id)); err != nil {
		log.Printf("%v\n", err)
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, "Basket Was Deleted Successfully.")

}

func (b *Basket) Register(g *echo.Group) {
	g.GET("/", b.GetBaskets)
	g.POST("/", b.Create)
	g.PATCH("/:id", b.Update)
	g.GET("/:id", b.GetByID)
	g.DELETE("/:id", b.Delete)
}
