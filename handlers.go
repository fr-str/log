package log

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

// CorrelationIDHeaderKey is the HTTP header key used for transmitting correlation IDs.
const CorrelationIDHeaderKey = "X-Correlation-ID"

// HTTPHandler is a logging middleware for http server
// ordering multiple middlewares for http.Mux is important
// logger should be the last one due to Handler call order
//
// Example:
//
//	mux := http.NewServeMux()
//	mux.HandleFunc("/v1/hello", HelloHandler)
//	http.ListenAndServe(addr, log.HTTPHandler(mux))
func HTTPHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		correlationID := r.Header.Get(CorrelationIDHeaderKey)
		if correlationID == "" {
			correlationID = fmt.Sprintf("unknown-%s", uuid.NewString())
		}
		ctx := context.WithValue(r.Context(), CorrelationIdKey, correlationID)
		r = r.WithContext(ctx)

		remoteAddr := r.RemoteAddr
		fwdAddr := r.Header.Get("X-Forwarded-For")
		if fwdAddr != "" {
			// Got X-Forwarded-For
			remoteAddr = fwdAddr

			// If we got an array, grab the first IP
			ips := strings.Split(fwdAddr, ", ")
			if len(ips) > 1 {
				remoteAddr = ips[0]
			}
		}

		ts := time.Now()
		handler.ServeHTTP(w, r)
		InfoCtx(r.Context(), r.Method,
			String("path", r.URL.Path),
			Duration("duration", time.Since(ts)),
			String("remote", remoteAddr),
			String("agent", r.UserAgent()),
		)
	})
}
