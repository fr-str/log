package log

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"path/filepath"
	"strings"

	"github.com/fr-str/log/level"
)

const (
	// time.Format used by TextHandler
	timeFormat = "15:04:05.000"

	// format: <time> <level> <caller> <message> <attrs>
	logFormatStr = "%s %s \x1b[90m%s\x1b[0m \x1b[97m%s %s\x1b[0m\n"
)

// textHandler is a custom slog.Handler which renders logs with correct format
type textHandler struct {
	out       io.Writer
	marshaler func(any) ([]byte, error)
}

// TextHandler returns slog.Handler for rendering message in text format
func TextHandler(out io.Writer, multiline bool) *textHandler {
	// Choose if json.marshal must be idented or not
	marshaler := json.Marshal
	if multiline {
		marshaler = func(x any) ([]byte, error) {
			return json.MarshalIndent(x, "", "  ")
		}
	}

	return &textHandler{
		out:       out,
		marshaler: marshaler,
	}
}

// Handle renders log message and writes on output
func (h *textHandler) Handle(ctx context.Context, r slog.Record) error {
	b, err := h.marshaler(mapFromAttrs(r))
	if err != nil {
		return err
	}
	if len(b) <= 2 {
		b = nil
	}

	_, err = fmt.Fprintf(h.out, logFormatStr,
		r.Time.Format(timeFormat),
		level.TextFromSlog(r.Level),
		h.getCaller(ctx),
		r.Message,
		string(b),
	)
	return err
}

func (*textHandler) getCaller(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	caller, ok := ctx.Value(callerID).(string)
	if !ok {
		return ""
	}

	s := strings.Split(filepath.Dir(caller), "/")
	path := s[len(s)-1] + "/" + filepath.Base(caller)

	return path
}

func (h *textHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return true
}

func (h *textHandler) WithAttrs([]slog.Attr) slog.Handler {
	return h
}

func (h *textHandler) WithGroup(string) slog.Handler {
	return h
}

type attrser interface {
	Attrs(f func(Attr) bool)
	NumAttrs() int
}

// mapFromAttrs creates map from attr slice
func mapFromAttrs(r attrser) map[string]any {
	fields := make(map[string]any, r.NumAttrs())
	r.Attrs(func(a Attr) bool {
		switch a.Value.Kind() {
		case slog.KindGroup:
			fields[a.Key] = mapFromAttrs(groupAttrs(a.Value.Group()))
		default:
			fields[a.Key] = a.Value.Any()
		}
		return true
	})
	return fields
}

type groupAttrs []slog.Attr

func (g groupAttrs) Attrs(f func(Attr) bool) {
	for _, attr := range g {
		if !f(attr) {
			return
		}
	}
}

func (g groupAttrs) NumAttrs() int {
	return len(g)
}
