package level

import (
	"log/slog"
	"strings"
)

// Those are missing levels from slog.Logger
const (
	Trace = slog.Level(-8)
)

// Colors used by TextHandler renderer
const (
	TraceText = "\x1b[37;1mTRACE\x1b[0m"
	DebugText = "\x1b[35;1mDEBUG\x1b[0m"
	InfoText  = "\x1b[34;1mINFO\x1b[0m"
	WarnText  = "\x1b[33;1mWARN\x1b[0m"
	ErrorText = "\x1b[31;1mERROR\x1b[0m"
)

func TextToSlog(level string) slog.Level {
	switch strings.ToLower(level)[0] {
	case 't':
		return Trace
	case 'd':
		return slog.LevelDebug
	case 'i':
		return slog.LevelInfo
	case 'w':
		return slog.LevelWarn
	case 'e':
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func TextFromSlog(level slog.Level) string {
	switch level {
	case Trace:
		return TraceText
	case slog.LevelDebug:
		return DebugText
	case slog.LevelInfo:
		return InfoText
	case slog.LevelWarn:
		return WarnText
	case slog.LevelError:
		return ErrorText
	default:
		return level.String()
	}
}
