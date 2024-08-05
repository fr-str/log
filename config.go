package log

import (
	"io"
	"os"

	"github.com/fr-str/env"
)

// Logging configuration parameters.
var (
	logLevel     string = env.Get("LOG_LEVEL", "debug")
	logFormat    string = env.Get("LOG_FORMAT", "json")
	logMultiline bool   = env.Get("LOG_MULTILINE", false)
)

type Config struct {
	Level     string
	Format    string
	Multiline *bool

	Output   io.Writer
	pcOffset int
}

// defaults sets default value for fields that have not been set
func (cfg *Config) defaults() {
	if cfg.Level == "" {
		cfg.Level = logLevel
	}
	if cfg.Format == "" {
		cfg.Format = logFormat
	}
	if cfg.Multiline == nil {
		cfg.Multiline = &logMultiline
	}
	if cfg.Output == nil {
		cfg.Output = os.Stdout
	}
}
