package main

import (
	"github.com/Hamid-Rezaei/goBasket/internal/infra/db"
	"github.com/Hamid-Rezaei/goBasket/internal/infra/http/handler"
	"github.com/Hamid-Rezaei/goBasket/internal/infra/repository"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Cannot find env file")
	}

	app := echo.New()

	// Middleware
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	dbConnection, err := db.New()
	if err != nil {
		log.Fatal("Cannot connect to database.")
	}

	db.AutoMigrate(dbConnection)

	v1 := app.Group("/api")

	br := repository.NewBasketRepo(dbConnection)
	ur := repository.NewUserRepo(dbConnection)

	h := handler.NewHandler(ur, br)
	h.Register(v1)
	if err := app.Start("0.0.0.0:1373"); err != nil {
		log.Fatalf("server failed to start %v", err)
	}
}
