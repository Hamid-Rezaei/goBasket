package main

import (
	"github.com/Hamid-Rezaei/goBasket/internal/infra/db"
	"github.com/Hamid-Rezaei/goBasket/internal/infra/http/handler"
	"github.com/Hamid-Rezaei/goBasket/internal/infra/repository"
	"github.com/Hamid-Rezaei/goBasket/internal/infra/router"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Cannot find env file")
	}

	r := router.New()

	dbConnection, err := db.New()
	if err != nil {
		log.Fatal("Cannot connect to database.")
	}

	db.AutoMigrate(dbConnection)

	v1 := r.Group("/api")

	br := repository.NewBasketRepo(dbConnection)
	ur := repository.NewUserRepo(dbConnection)

	h := handler.NewHandler(ur, br)
	h.Register(v1)
	if err := r.Start("0.0.0.0:1373"); err != nil {
		log.Fatalf("server failed to start %v", err)
	}
}
