package handler

import (
	"github.com/Hamid-Rezaei/goBasket/internal/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (h *Handler) Register(v1 *echo.Group) {
	jwtMiddleware := middleware.JWT(utils.JWTSecret)
	guestUsers := v1.Group("/users")
	guestUsers.POST("", h.SignUp)
	guestUsers.POST("/login", h.Login)

	v1.GET("/basket/", h.GetBaskets)
	v1.POST("/basket/", h.CreateBasket)
	v1.PATCH("/basket/:id", h.UpdateBasket)
	v1.GET("/basket/:id", h.GetBasketByID)
	v1.DELETE("/basket/:id", h.DeleteBasket)

}
