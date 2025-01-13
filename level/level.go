package level

import (
	"fmt"
	"log/slog"
	"strconv"
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
	var mod int
	_, extra, found := strings.Cut(level, "-")
	if found {
		var err error
		mod, err = strconv.Atoi(extra)
		if err != nil {
			fmt.Printf("\n%s\n", err)
		}
	}

	var ret slog.Level
	switch strings.ToLower(level)[0] {
	case 't':
		ret = Trace
	case 'd':
		ret = slog.LevelDebug
	case 'i':
		ret = slog.LevelInfo
	case 'w':
		ret = slog.LevelWarn
	case 'e':
		ret = slog.LevelError
	default:
		ret = slog.LevelInfo
	}
	return ret - slog.Level(mod)
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
