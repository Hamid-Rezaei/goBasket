package main

import (
	"github.com/Hamid-Rezaei/goBasket/internal/infra/db"
	"github.com/Hamid-Rezaei/goBasket/internal/infra/http/handler"
	"github.com/Hamid-Rezaei/goBasket/internal/infra/repository"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Cannot find env file")
	}

	app := echo.New()

	dbConnection, err := db.New()
	if err != nil {
		log.Fatal("Cannot connect to database.")
	}

	db.AutoMigrate(dbConnection)

	basketRepo := repository.New(dbConnection)

	h := handler.NewBasket(basketRepo)
	h.Register(app.Group("/basket"))

	if err := app.Start("0.0.0.0:1373"); err != nil {
		log.Fatalf("server failed to start %v", err)
	}
}
