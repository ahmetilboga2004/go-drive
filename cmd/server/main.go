package main

import "github.com/ahmetilboga2004/internal/infrastructure/utils/logger"

func main() {
	logger.InitializeLogger()
	logger.Log.Info("Server started")
}
