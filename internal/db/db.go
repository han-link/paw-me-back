package db

import (
	"fmt"
	"paw-me-back/internal/env"
	"paw-me-back/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func New(debug bool) (*gorm.DB, error) {
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

	logLevel := logger.Error

	if debug {
		logLevel = logger.Info
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})

	if err != nil {
		return nil, err
	}

	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)

	if err = db.SetupJoinTable(&model.Group{}, "Members", &model.UserGroup{}); err != nil {
		return nil, err
	}
	if err = db.SetupJoinTable(&model.User{}, "Groups", &model.UserGroup{}); err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&model.User{},
		&model.Group{},
		&model.Expense{},
		&model.Payee{},
		&model.UserGroup{},
	)

	if err != nil {
		return nil, err
	}

	return db, nil
}
