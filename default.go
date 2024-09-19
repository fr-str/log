package log

import (
	"context"
	"encoding/json"
	"fmt"
)

var DefaultLogger Logger = NewWithConfiguration(Config{
	pcOffset: 1,
})

func Trace(msg string, args ...any) {
	DefaultLogger.Trace(msg, args...)
}

func Debug(msg string, args ...any) {
	DefaultLogger.Debug(msg, args...)
}

func Info(msg string, args ...any) {
	DefaultLogger.Info(msg, args...)
}

func Warn(msg string, args ...any) {
	DefaultLogger.Warn(msg, args...)
}

func Error(msg string, args ...any) {
	DefaultLogger.Error(msg, args...)
}

func TraceCtx(ctx context.Context, msg string, args ...any) {
	DefaultLogger.TraceCtx(ctx, msg, args...)
}

func DebugCtx(ctx context.Context, msg string, args ...any) {
	DefaultLogger.DebugCtx(ctx, msg, args...)
}

func InfoCtx(ctx context.Context, msg string, args ...any) {
	DefaultLogger.InfoCtx(ctx, msg, args...)
}

func WarnCtx(ctx context.Context, msg string, args ...any) {
	DefaultLogger.WarnCtx(ctx, msg, args...)
}

func ErrorCtx(ctx context.Context, msg string, args ...any) {
	DefaultLogger.ErrorCtx(ctx, msg, args...)
}

func PrintJSON(v any) {
	b, err := json.MarshalIndent(v, "", " ")
	if err != nil {
		Error("------ PrintJSON -----", Err(err))
		return
	}
	fmt.Println(string(b))
}
