// Package httpapi exposes a register node over HTTP.
//
// Handlers here do no domain work. They decode a request, hand it to the package
// that owns the decision, and encode the answer. Anything a fonctionnaire or a
// citizen can do must be testable without starting a server.
package httpapi

import (
	"log/slog"
	"net/http"
	"time"
)

// NewRouter builds the node's HTTP handler.
func NewRouter(logger *slog.Logger) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", health)

	return logging(logger, mux)
}

func health(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"status":"ok"}` + "\n"))
}

// logging records every request, reads included.
//
// Who consulted which record is a control in its own right here, not only a
// diagnostic: in an administration where information itself is sold, a
// consultation trail is one of the few things that makes quiet abuse visible.
func logging(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		logger.Info("request",
			"method", r.Method,
			"path", r.URL.Path,
			"duration_ms", time.Since(start).Milliseconds(),
		)
	})
}
