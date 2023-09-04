package logger

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.Logger

	LOG_OUTPUT = "LOG_OUTPUT"
	LOG_LEVEL  = "LOG_LEVEL"
)

// init initializes the logging configuration.
//
// No parameters.
// No return types.
func init() {
	logConfig := zap.Config{
		OutputPaths: []string{getOutputLogs()},
		Level:       zap.NewAtomicLevelAt(getLevelLogs()),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "message",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	log, _ = logConfig.Build()
}

// Info logs the provided message with the given tags as info level.
//
// Parameters:
// - message: the message to be logged (string).
// - tags: optional tags to be included in the log (zap.Field).
// Return type: void.
func Info(message string, tags ...zap.Field) {
	log.Info(message, tags...)
	log.Sync()
}

// Error logs an error message with optional tags.
//
// message: the error message to log.
// err: the error object to log.
// tags: optional additional tags to include in the log.
func Error(message string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.Info(message, tags...)
	log.Sync()
}

// getOutputLogs returns the output logs of the program.
//
// It does not take any parameters.
// It returns a string.
func getOutputLogs() string {
	output := strings.ToLower(strings.TrimSpace(os.Getenv(LOG_OUTPUT)))
	if output == "" {
		return "stdout"
	}

	return output
}

// getLevelLogs returns the zapcore.Level based on the value of the LOG_LEVEL environment variable.
//
// It takes no parameters.
// It returns a zapcore.Level.
func getLevelLogs() zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(os.Getenv(LOG_LEVEL))) {
	case "info":
		return zapcore.InfoLevel
	case "error":
		return zapcore.ErrorLevel
	case "debug":
		return zapcore.DebugLevel
	default:
		return zapcore.InfoLevel
	}
}
