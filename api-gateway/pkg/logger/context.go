package logger

import (
	"context"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ctxLogger struct{}

// ContextWithLogger adds logger to context
func ContextWithLogger(ctx context.Context, l *logrus.Logger) context.Context {
	return context.WithValue(ctx, ctxLogger{}, l)
}

// LoggerFromContext returns logger from context
func LoggerFromContext(ctx context.Context) *logrus.Logger {
	if l, ok := ctx.Value(ctxLogger{}).(*logrus.Logger); ok {
		return l
	}
	return logrus.StandardLogger()
}

// EntryWithRequestIDFromContext returns logger from context with request ID
func EntryWithRequestIDFromContext(ctx context.Context) *logrus.Entry {
	if l, ok := ctx.Value(ctxLogger{}).(*logrus.Logger); ok {
		requestID, ok := ctx.Value(middleware.RequestIDHeader).(string)
		if ok {
			return l.WithField("request_id", requestID)
		}

		return l.WithField("request_id", "unknown")
	}
	return logrus.StandardLogger().WithField("request_id", "unknown")
}

// WithLogger middleware для обогащение контекста запроса логером
func WithLogger(l *logrus.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(context.WithValue(r.Context(), ctxLogger{}, l))
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
