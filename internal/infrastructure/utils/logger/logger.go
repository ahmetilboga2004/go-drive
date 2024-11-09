package logger

import (
	"os"

	"go.uber.org/zap"
)

var Log *zap.Logger

func InitializeLogger() {
	_, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("Error opening log file: " + err.Error())
	}

	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout", "logs/app.log"}
	config.ErrorOutputPaths = []string{"stderr"}

	config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)

	var err2 error
	Log, err2 = config.Build()
	if err2 != nil {
		panic("Error building zap logger: " + err2.Error())
	}

	zap.ReplaceGlobals(Log)
}
