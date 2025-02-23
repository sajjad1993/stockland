package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"time"

	"github.com/hjhussaini/shared-go/log"
)

type Logger interface {
	Info(string, ...any)
	Fatal(string, ...any)
	Error(string, ...any)
}

// hjLogger implements Logger interface
type hjLogger struct {
	logger log.Logger
}

func (hLog *hjLogger) Info(format string, values ...any) {
	hLog.logger.Info(fmt.Sprintf(format, values...))
}

func (hLog *hjLogger) Fatal(format string, values ...any) {
	hLog.logger.Fatal(fmt.Sprintf(format, values...))
}
func (hLog *hjLogger) Error(format string, values ...any) {
	hLog.logger.Error(fmt.Sprintf(format, values...))
}

func NewLogger() Logger {
	logger := &hjLogger{
		logger: log.NewLogger(log.ErrorLevel, log.NewStdoutCore(), newFileLogger()),
	}

	return logger
}

func newFileLogger() zapcore.Core {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	fileName := "log/web/" + time.Now().Format("2006.01.02") + ".log"
	level := getLoggerLevel("debug")
	syncWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   1 << 30, //1G
		LocalTime: true,
		Compress:  true,
	})
	encoder := zap.NewProductionEncoderConfig()
	encoder.TimeKey = "timestamp"
	encoder.CallerKey = "module"
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	var enc zapcore.Encoder
	enc = zapcore.NewJSONEncoder(encoder)

	core := zapcore.NewCore(enc, syncWriter, zap.NewAtomicLevelAt(level))
	return core
}

func getLoggerLevel(lvl string) zapcore.Level {
	if level, ok := levelMap[lvl]; ok {
		return level
	}
	return zapcore.InfoLevel
}

var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}
