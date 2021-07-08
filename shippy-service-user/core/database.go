package core

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateConnection() (*gorm.DB, error) {

	// Get database details from environment variables
	host := "127.0.0.1" // тут надо будет походу тоже использовать мостик докера и прописать просто database:7432
	port := "7432"
	user := "postgres"
	password := "postgres"
	dbName := "microservice"

	db, err := gorm.Open(postgres.Open(fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=disable password=%s port=%s",
		host, user, dbName, password, port)), &gorm.Config{})
	return db, err
}
