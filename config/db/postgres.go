package db

import (
	"fmt"
	"micro_backend_film/common/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres() *gorm.DB {
	dbUser := "postgres"
	dbPass := "minh123"
	dbHost := "localhost"
	dbName := "go_compare_price"
	dbPort := 5432

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", dbHost, dbUser, dbPass, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}

	db.AutoMigrate(&entity.User{}, &entity.Film{}, &entity.Bookmark{})
	fmt.Println("Connect DB successfully")
	return db
}

func ClosePostgres(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}
