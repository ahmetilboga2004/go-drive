package main

import (
	"net/http"

	"github.com/ahmetilboga2004/internal/application/handlers"
	"github.com/ahmetilboga2004/internal/application/middlewares"
	"github.com/ahmetilboga2004/internal/domain/repositories"
	"github.com/ahmetilboga2004/internal/domain/services"
	"github.com/ahmetilboga2004/internal/infrastructure/config"
	"github.com/ahmetilboga2004/internal/infrastructure/config/database"
	"github.com/ahmetilboga2004/internal/infrastructure/utils/logger"
)

func main() {
	jwtConfig := services.JwtConfig{
		AccessTokenSecret:  config.JWT.AccessSecretKey,
		RefreshTokenSecret: config.JWT.RefreshSecretKey,
		AccessTokenExp:     config.JWT.AccessTokenExpiration,
		RefreshTokenExp:    config.JWT.RefreshTokenExpiration,
	}
	jwtService := services.NewJwtService(&jwtConfig)
	userRepo := repositories.NewUserRepository(database.DB)
	userService := services.NewUserService(userRepo, jwtService)
	userHandler := handlers.NewUserHandler(userService)

	fileRepo := repositories.NewFileRepository(database.DB)
	fileService := services.NewFileService(fileRepo)
	fileHandler := handlers.NewFileHandler(fileService)

	authMiddleware := middlewares.NewAuthMiddleware(jwtService)

	mux := http.NewServeMux()

	authMux := authMiddleware.Auth(mux)

	mux.HandleFunc("POST /users/register", authMiddleware.GuestOnly(userHandler.Register))
	mux.HandleFunc("POST /users/login", authMiddleware.GuestOnly(userHandler.Login))
	mux.HandleFunc("GET /users/refresh-token", userHandler.RefreshToken)

	mux.HandleFunc("GET /files/", fileHandler.GetAll)
	mux.HandleFunc("GET /files/{id}", fileHandler.GetByID)
	mux.HandleFunc("POST /files/", authMiddleware.RequireLogin(fileHandler.Upload))
	mux.HandleFunc("PUT /files/{id}", fileHandler.Update)

	logger.Log.Sugar().Infof("server started on %s", config.APP.BaseURL)
	logger.Log.Sugar().Fatal(http.ListenAndServe(":"+config.APP.Port, authMux))
}
