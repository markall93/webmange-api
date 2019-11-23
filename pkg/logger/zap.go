package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var logger *zap.SugaredLogger

func InitLog(level string,filePath string) {
	hook := lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    1,
		MaxAge:     0,
		Compress:   false,
	}
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
	var writers []zapcore.WriteSyncer
	writers = append(writers, os.Stdout)
	writers = append(writers, zapcore.AddSync(&hook))
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(getLevel(level))
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(writers...),
		atomicLevel,
	)
	caller := zap.AddCaller()
	development := zap.Development()
	logger= zap.New(core, caller, development,zap.AddCallerSkip(1)).Sugar()
}


func Error(args ...interface{}) {
	logger.Error(args...)
}

func getLevel(level string)(l zapcore.Level) {
	switch level{
	case "debug":
		l=zap.DebugLevel
	case "info":
		l=zap.InfoLevel
	case "warn":
		l=zap.WarnLevel
	case "error":
		l=zap.ErrorLevel
	case "panic":
		l=zap.PanicLevel
	case "fatal":
		l=zap.FatalLevel
	default:
		l=zap.InfoLevel
	}
	return
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

