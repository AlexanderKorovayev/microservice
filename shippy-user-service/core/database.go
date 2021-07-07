package core

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func CreateConnection() (*gorm.DB, error) {

	// Get database details from environment variables
	host := "127.0.0.1:5432" // тут надо будет походу тоже использовать мостик докера и прописать просто database:5432
	db_name := "microservice"
	user := "postgres"
	password := "postgres"

	return gorm.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s user=%s dbname=%s sslmode=disable password=%s",
			host, user, db_name, password,
		),
	)
}
