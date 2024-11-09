package logger

import (
	"os"

	"go.uber.org/zap"
)

var Log *zap.Logger

// InitializeLogger global logger'ı başlatır ve yapılandırır
func InitializeLogger() {
	// Dosyaya yazmak için bir Zap logger'ı oluşturuyoruz
	_, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("Error opening log file: " + err.Error())
	}

	// Log formatını ve log seviyesini belirliyoruz (JSON ve Info seviyesi)
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout", "logs/app.log"} // Hem stdout'a hem de dosyaya yazılacak
	config.ErrorOutputPaths = []string{"stderr"}

	// Log seviyesini belirleyin (örneğin, InfoLevel)
	config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)

	// Zap logger'ı oluştur
	var err2 error
	Log, err2 = config.Build()
	if err2 != nil {
		panic("Error building zap logger: " + err2.Error())
	}

	// Global logger olarak yapılandırılıyor
	zap.ReplaceGlobals(Log)
}
