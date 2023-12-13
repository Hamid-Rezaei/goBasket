package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func New() (*gorm.DB, error) {
	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  false,         // Disable color
		},
	)

	pgConfig := postgres.New(postgres.Config{DSN: CreateURI()})
	db, err := gorm.Open(pgConfig, &gorm.Config{Logger: dbLogger})

	if err != nil {
		return nil, err
	}
	return db, nil
}

func CreateURI() string {
	driver := os.Getenv("DB_DRIVER")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	db := os.Getenv("DB")

	DSN := fmt.Sprintf("%s://%s:%s@%s:%s/%s", driver, username, password, host, port, db)
	return DSN
}
