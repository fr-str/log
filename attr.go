package log

import (
	"encoding/json"
	"log/slog"
	"reflect"
	"time"

	"golang.org/x/exp/constraints"
)

// An Attr is a key-value pair.
type Attr = slog.Attr

// Any returns any Attr
//
// Use only if value is of a custom type or if specific Attr does not exist
func Any(key string, value any) Attr {
	return Attr{Key: key, Value: slog.AnyValue(value)}
}

// Bool returns bool Attr
func Bool(key string, v bool) Attr {
	return Attr{Key: key, Value: slog.BoolValue(v)}
}

// Duration returns duration Attr
func Duration(key string, v time.Duration) Attr {
	return Attr{
		Key:   key,
		Value: slog.StringValue(v.Round(time.Microsecond).String()),
	}
}

// Float returns float64 Attr
func Float[T constraints.Float](key string, v T) Attr {
	return Attr{Key: key, Value: slog.Float64Value(float64(v))}
}

// Group returns group Attr
func Group(key string, v ...Attr) Attr {
	return Attr{Key: key, Value: slog.GroupValue(v...)}
}

// Int returns int64 Attr
func Int[T constraints.Signed](key string, v T) Attr {
	return Attr{Key: key, Value: slog.Int64Value(int64(v))}
}

// String returns string Attr
func String(key, value string) Attr {
	return Attr{Key: key, Value: slog.StringValue(value)}
}

// Time returns time Attr
func Time(key string, v time.Time) Attr {
	return Attr{Key: key, Value: slog.TimeValue(v)}
}

// Uint returns uint64 Attr
func Uint[T constraints.Unsigned](key string, v T) Attr {
	return Attr{Key: key, Value: slog.Uint64Value(uint64(v))}
}

// Err returns error Attr
func Err(v error) Attr {
	return Attr{Key: "error", Value: slog.StringValue(v.Error())}
}

func JSON(v any) Attr {
	return NamedJSON(reflect.TypeOf(v).Name(), v)
}

func NamedJSON(k string, v any) Attr {
	b, err := json.MarshalIndent(v, "", " ")
	if err != nil {
		return Attr{Key: k, Value: slog.StringValue(err.Error())}
	}
	return Attr{Key: k, Value: slog.AnyValue(json.RawMessage(b))}
}
