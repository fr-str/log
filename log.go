package log

import (
	"context"
	"log/slog"

	"github.com/fr-str/log/level"
)

// Logger is a custom slog.Logger implementation
type Logger struct {
	Logger *slog.Logger
}

// New returns logger with default configuration
func New() Logger {
	return NewWithConfiguration(Config{})
}

// NewWithConfiguration returns logger with given configuration
func NewWithConfiguration(cfg Config) Logger {
	cfg.defaults()

	var handler slog.Handler

	switch {
	case cfg.Format == "json":
		// Production/JSON logger
		handler = slog.NewJSONHandler(cfg.Output, &slog.HandlerOptions{
			Level: level.TextToSlog(cfg.Level),
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				// Fix TRACE and FATAL levels
				if a.Key == slog.LevelKey {
					l := a.Value.Any().(slog.Level)
					switch l {
					case level.Trace:
						a.Value = slog.StringValue("TRACE")
					}
				}
				return a
			},
		})
	default:
		handler = TextHandler(cfg.Output, *cfg.Multiline)
	}

	handler = MetadataHandler(handler, cfg)
	return Logger{
		Logger: slog.New(handler),
	}
}

// Logs with default logger at level TRACE
func (l *Logger) Trace(msg string, args ...any) {
	l.Logger.Log(context.Background(), level.Trace, msg, args...)
}

// Logs with default logger at level DEBUG

func (l *Logger) Debug(msg string, args ...any) {
	l.Logger.Debug(msg, args...)
}

// Logs with default logger at level INFO
func (l *Logger) Info(msg string, args ...any) {
	l.Logger.Info(msg, args...)
}

// Logs with default logger at level WARN
func (l *Logger) Warn(msg string, args ...any) {
	l.Logger.Warn(msg, args...)
}

// Logs with default logger at level ERROR
func (l *Logger) Error(msg string, args ...any) {
	l.Logger.Error(msg, args...)
}

// Logs with default logger at level TRACE with given context
func (l *Logger) TraceCtx(ctx context.Context, msg string, args ...any) {
	l.Logger.Log(ctx, level.Trace, msg, args...)
}

// Logs with default logger at level DEBUG with given context
func (l *Logger) DebugCtx(ctx context.Context, msg string, args ...any) {
	l.Logger.DebugContext(ctx, msg, args...)
}

// Logs with default logger at level INFO with given context
func (l *Logger) InfoCtx(ctx context.Context, msg string, args ...any) {
	l.Logger.InfoContext(ctx, msg, args...)
}

// Logs with default logger at level WARN with given context
func (l *Logger) WarnCtx(ctx context.Context, msg string, args ...any) {
	l.Logger.WarnContext(ctx, msg, args...)
}

// Logs with default logger at level ERROR with given context
func (l *Logger) ErrorCtx(ctx context.Context, msg string, args ...any) {
	l.Logger.ErrorContext(ctx, msg, args...)
}
