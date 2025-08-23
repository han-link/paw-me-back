package db

import (
	"fmt"
	"paw-me-back/internal/env"
	"paw-me-back/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		`host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s`,
		env.GetString("POSTGRES_HOST", "localhost"),
		env.GetString("POSTGRES_USER", "postgres"),
		env.GetString("POSTGRES_PW", "postgres"),
		env.GetString("POSTGRES_DB", "paw-me-back"),
		env.GetString("POSTGRES_PORT", "5432"),
		env.GetString("POSTGRES_SSL_MODE", "disable"),
		env.GetString("TIMEZONE", "Europe/Berlin"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)

	err = db.AutoMigrate(
		&model.User{},
		&model.Group{},
		&model.Expense{},
		&model.Payee{},
	)

	if err != nil {
		return nil, err
	}

	return db, nil
}
