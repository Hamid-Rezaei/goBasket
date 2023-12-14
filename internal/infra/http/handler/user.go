package handler

import (
	"github.com/Hamid-Rezaei/goBasket/internal/domain/model"
	"github.com/Hamid-Rezaei/goBasket/internal/infra/http/request"
	"github.com/Hamid-Rezaei/goBasket/internal/infra/http/response"
	_ "github.com/Hamid-Rezaei/goBasket/internal/infra/http/response"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func (h *Handler) SignUp(c echo.Context) error {
	var req request.UserRegisterRequest

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

	var user model.User

	hash, err := user.HashPassword(req.Password)
	if err != nil {
		return err
	}
	user.Password = hash
	user.Username = req.Username
	user.Email = req.Email

	id, err := h.userRepo.Create(c.Request().Context(), user)

	if err != nil {
		log.Printf("%v\n", err)
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusCreated, id)
}

func (h *Handler) Login(c echo.Context) error {
	var req request.UserLoginRequest

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

	user, err := h.userRepo.GetByPassword(c.Request().Context(), req.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if user == nil {
		return c.JSON(http.StatusForbidden, "Access Forbidden!")
	}
	if !user.CheckPassword(req.Password) {
		return c.JSON(http.StatusForbidden, "Access Forbidden!")
	}

	return c.JSON(http.StatusOK, response.NewUserResponse(user))
}
