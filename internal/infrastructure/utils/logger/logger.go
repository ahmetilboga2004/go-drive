package logger

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Log *zap.Logger

func init() {
	if err := os.MkdirAll("logs", 0755); err != nil {
		panic(fmt.Sprintf("failed to create logs directory: %v", err))
	}
	rotatingLogger := &lumberjack.Logger{
		Filename:  fmt.Sprintf("logs/app-%s.log", time.Now().Format("2006-01-02")),
		MaxSize:   10,
		LocalTime: true,
	}

	encodingConfig := zap.NewProductionEncoderConfig()
	encodingConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encodingConfig),
		zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(os.Stdout),
			zapcore.AddSync(rotatingLogger),
		),
		zap.InfoLevel,
	)

	Log = zap.New(core)
	zap.ReplaceGlobals(Log)
}
