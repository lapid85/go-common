package log

import "golang.org/x/exp/slog"

// Warn 警告日志
func Warn(msg string, args ...interface{}) {
	slog.Warn(msg, args...)
}

// Debug 调试日志
func Debug(msg string, args ...interface{}) {
	slog.Debug(msg, args...)
}

// Info 信息日志
func Info(msg string, args ...interface{}) {
	slog.Info(msg, args...)
}

// Error 错误日志
func Error(msg string, args ...interface{}) {
	slog.Error(msg, args...)
}
