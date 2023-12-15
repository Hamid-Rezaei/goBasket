package handler

import (
	"errors"
	"github.com/Hamid-Rezaei/goBasket/internal/domain/model"
	"github.com/Hamid-Rezaei/goBasket/internal/infra/http/request"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) CreateBasket(c echo.Context) error {
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

	var basket model.Basket
	basket.Data = req.Data
	basket.State = req.State
	basket.UserID = userIDFromToken(c)

	id, err := h.basketRepo.Create(c.Request().Context(), basket)

	if err != nil {
		log.Printf("%v\n", err)
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusCreated, id)
}

func (h *Handler) UpdateBasket(c echo.Context) error {

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

	userID := userIDFromToken(c)
	basket, err := h.basketRepo.GetUserBasketByID(c.Request().Context(), userID, uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNoContent, "Basket Not Found!")
		}

		return echo.ErrInternalServerError
	}

	if basket.State == "COMPLETED" {
		return c.JSON(http.StatusNotAcceptable, "Cannot UpdateBasket Completed Basket!")
	}

	basketModel := model.Basket{
		Data:  req.Data,
		State: req.State,
	}

	if err := h.basketRepo.Update(c.Request().Context(), basketModel, int(id)); err != nil {
		log.Printf("%v\n", err)
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, "Basket Was Updated Successfully.")

}

func (h *Handler) GetBaskets(c echo.Context) error {
	userID := userIDFromToken(c)
	baskets, err := h.basketRepo.GetUserBaskets(c.Request().Context(), userID)

	if err != nil {
		log.Printf("%v\n", err)
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, baskets)

}

func (h *Handler) GetBasketByID(c echo.Context) error {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Printf("%v\n", err)
		return echo.ErrBadRequest
	}

	userID := userIDFromToken(c)
	basket, err := h.basketRepo.GetUserBasketByID(c.Request().Context(), userID, uint(id))

	if err != nil {
		log.Printf("%v\n", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNoContent, "Basket Not Found!")
		}

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, basket)

}

func (h *Handler) DeleteBasket(c echo.Context) error {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Printf("%v\n", err)
		return echo.ErrBadRequest
	}

	userID := userIDFromToken(c)
	basket, err := h.basketRepo.GetUserBasketByID(c.Request().Context(), userID, uint(id))

	if basket == nil {
		log.Printf("%s", "User does not have permission!")
		return c.JSON(http.StatusForbidden, "User does not have permission or not found!")
	}
	if err != nil {
		log.Printf("%v\n", err)
		return echo.ErrInternalServerError
	}

	if err := h.basketRepo.Delete(c.Request().Context(), int(id)); err != nil {
		log.Printf("%v\n", err)
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, "Basket Was Deleted Successfully.")

}
