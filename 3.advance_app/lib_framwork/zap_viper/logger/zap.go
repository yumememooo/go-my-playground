package logger

import (
	"mywork/go-my-playground/3.advance_app/lib_framwork/zap/config"
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var ZapLogger *zap.Logger
var SugarLogger *zap.SugaredLogger

func InitLogger() {

	logWriter := []zapcore.WriteSyncer{zapcore.AddSync(os.Stdout)}

	if config.C.File != "" {
		hook := setFileWriter(config.C.File)
		logWriter = append(logWriter, hook)
	}
	encoderConfig := setEncoder()

	level := getLogLevel(config.C.Level)

	core := zapcore.NewCore(encoderConfig,
		zapcore.NewMultiWriteSyncer(logWriter...),
		level)

	ZapLogger = zap.New(core, zap.AddCaller()) //印出log的位置
	// ZapLogger.Debug("POK")                     // ZapLogger sample
	ZapLogger.Info("ZapLogger",
		zap.String("String", "ohoh"),
		zap.Int("Int", 3),
		zap.Duration("backoff", time.Second),
	)
	SugarLogger = ZapLogger.Sugar()
	//SugarLogger.Infof("Success! statusCode = %s for URL %s", "OK", "OK")  // SugarLogger sample

}

func setEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func setFileWriter(filePath string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getLogLevel(lv string) zapcore.Level {
	lv = strings.ToLower(lv)
	if level, ok := levelMap[lv]; ok {
		return level
	}
	return zapcore.InfoLevel
}

var levelMap = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
}
