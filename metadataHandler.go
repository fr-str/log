package log

import (
	"context"
	"fmt"
	"log/slog"
	"runtime"

	"github.com/fr-str/log/level"
)

type ContextMeta string

const (
	// CorrelationIdKey is a default key name for go contexts
	CorrelationIdKey ContextMeta = "correlation_id"

	// callerID is a key for context used by TextHandler
	callerID ContextMeta = "callerID"

	// PC - program counter, where log is created
	PC ContextMeta = "program_counter"
)

// metadataHandler handles Metadata type and adds attrs to logs
type metadataHandler struct {
	slog.Handler
	level slog.Level
	cfg   Config

	attrs []slog.Attr
}

// MetadataHandler returns slog.Handler with metadata
func MetadataHandler(handler slog.Handler, cfg Config) *metadataHandler {
	return &metadataHandler{
		Handler: handler,
		cfg:     cfg,
		level:   level.TextToSlog(cfg.Level),
	}
}

func (h *metadataHandler) Handle(ctx context.Context, r slog.Record) error {
	attrs := []slog.Attr{
		slog.String(string(CorrelationIdKey), h.getCorrelationId(ctx)),
	}

	caller := h.getCaller(ctx)
	if h.cfg.Format == "json" {
		attrs = append(attrs, slog.String("caller", h.getCaller(ctx)))
	}
	// Add default attrs
	r.AddAttrs(attrs...)

	// Add callerID for TextHandler
	ctx = context.WithValue(ctx, callerID, caller)
	return h.Handler.Handle(ctx, r)
}

// getCorrelationId extracts Correlation ID from ctx
func (h *metadataHandler) getCorrelationId(ctx context.Context) string {
	correlationId, ok := ctx.Value(CorrelationIdKey).(string)
	if !ok {
		return "unknown"
	}
	return correlationId
}

// getCaller returns caller with package, file and line number
func (h *metadataHandler) getCaller(ctx context.Context) string {
	pc, ok := ctx.Value(PC).(uintptr)
	if !ok {
		// Extract correct caller pc
		var pcs [1]uintptr
		runtime.Callers(6+h.cfg.pcOffset, pcs[:])
		pc = pcs[0]
	}

	frame, _ := runtime.CallersFrames([]uintptr{pc}).Next()
	return fmt.Sprintf("%s:%d", frame.File, frame.Line)
}

// Enabled returns true if level is enabled
func (h *metadataHandler) Enabled(ctx context.Context, l slog.Level) bool {
	return h.level <= l
}

// WithAttrs returns new handler with added attrs
func (h *metadataHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &metadataHandler{
		Handler: h.Handler,
		level:   h.level,
		cfg:     h.cfg,

		attrs: append(h.attrs, attrs...),
	}
}

func (h *metadataHandler) WithGroup(string) slog.Handler {
	return h
}
