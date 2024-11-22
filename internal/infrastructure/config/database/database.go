package database

import (
	"github.com/ahmetilboga2004/internal/domain/models"
	"github.com/ahmetilboga2004/internal/infrastructure/utils/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	dsn := "host=localhost user=gorm password=gorm dbname=godrive port=5432 sslmode=disable TimeZone=Europe/Istanbul"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		logger.Log.Sugar().Fatalf("database connection error: %v", err)
	}

	logger.Log.Info("database connection successfully")

	err = DB.AutoMigrate(&models.User{}, &models.File{})
	if err != nil {
		logger.Log.Sugar().Fatalf("database migration failed: %v", err)
	}

	logger.Log.Info("migration completed")
}
