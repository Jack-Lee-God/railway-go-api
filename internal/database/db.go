package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		host := os.Getenv("PGHOST")
		if host == "" {
			host = "localhost"
		}
		user := os.Getenv("PGUSER")
		if user == "" {
			user = "postgres"
		}
		password := os.Getenv("PGPASSWORD")
		dbname := os.Getenv("PGDATABASE")
		if dbname == "" {
			dbname = "postgres"
		}
		port := os.Getenv("PGPORT")
		if port == "" {
			port = "5432"
		}

		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			host,
			user,
			password,
			dbname,
			port,
		)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
