package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
)

var logger *Logger

type Logger struct {
	*zap.Logger
}

func Init() {
	logger = &Logger{initZapLogger()}
}

func initZapLogger() *zap.Logger {
	var err error
	_ = os.Mkdir("logs", 0755)

	allFile, err := os.OpenFile("logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		panic(fmt.Sprintf("Error: %s", err))
	}

	zapCfg := zap.NewProductionConfig()
	zapCfg.OutputPaths = []string{"stdout", allFile.Name()}

	zapLogger, err := zapCfg.Build()
	if err != nil {
		panic(err)
	}

	return zapLogger
}

func Debug(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)

}
func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}
