package migrations

import (
	"github.com/ahmetilboga2004/internal/domain/models"
	"github.com/ahmetilboga2004/internal/infrastructure/config"
	"github.com/ahmetilboga2004/internal/infrastructure/utils/logger"
	"go.uber.org/zap"
)

func init() {
	err := config.DB.AutoMigrate(&models.User{}, &models.File{})
	if err != nil {
		logger.Log.Fatal("failed to migrate models: %v", zap.Error(err))
	}
	logger.Log.Info("Migration complated successfully")
}
