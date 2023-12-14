package handler

import (
	"github.com/Hamid-Rezaei/goBasket/internal/infra/router/middleware"
	"github.com/Hamid-Rezaei/goBasket/internal/utils"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(v1 *echo.Group) {
	guestUsers := v1.Group("/users")
	guestUsers.POST("", h.SignUp)
	guestUsers.POST("/login", h.Login)

	user := v1.Group("/user", middleware.JWT(utils.GetSigningKey()))
	user.GET("", h.CurrentUser)

	baskets := v1.Group("/basket", middleware.JWTWithConfig(
		middleware.JWTConfig{
			SigningKey: utils.GetSigningKey(),
		},
	))
	baskets.GET("/", h.GetBaskets)
	baskets.POST("/", h.CreateBasket)
	baskets.PATCH("/:id", h.UpdateBasket)
	baskets.GET("/:id", h.GetBasketByID)
	baskets.DELETE("/:id", h.DeleteBasket)

}
