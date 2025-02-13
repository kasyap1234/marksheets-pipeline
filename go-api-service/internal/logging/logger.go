package logging

import (
	"go.uber.org/zap"
	"net/http"
	"time"
)

// generic implementation for use in other modules
type Logger interface {
	Infof(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	Debugf(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Fatalf(template string, args ...interface{})
}

// concrete implementation for ZapLogger
type ZapLogger struct {
	logger *zap.Logger
}

// Debugf implements Logger.
func (z *ZapLogger) Debugf(template string, args ...interface{}) {
	z.logger.Sugar().Debugf(template, args...)
}

// Errorf implements Logger.
func (z *ZapLogger) Errorf(template string, args ...interface{}) {
	z.logger.Sugar().Errorf(template, args...)
}

// Fatalf implements Logger.
func (z *ZapLogger) Fatalf(template string, args ...interface{}) {
	z.logger.Sugar().Fatalf(template, args...)
}

// Infof implements Logger.
func (z *ZapLogger) Infof(template string, args ...interface{}) {
	z.logger.Sugar().Infof(template, args...)
}

// Warnf implements Logger.
func (z *ZapLogger) Warnf(template string, args ...interface{}) {
	z.logger.Sugar().Warnf(template, args...)
}

// this is the method to be used in other files , will implement a new logger and will not care about the implementaiton
func NewLogger() Logger {
	logger := zap.Must(zap.NewProduction())
	defer logger.Sync()

	return &ZapLogger{
		logger: logger,
	}

}

// using Logger interface to be able to implement all its methods ;
func RequestLogger(logger Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			ww := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
			next.ServeHTTP(ww, r)
			duration := time.Since(start)

			logger.Infof("Request: %s %s %s, Status: %d, Duration: %v", r.Method, r.URL.Path, r.RemoteAddr, ww.statusCode, duration)
		})
	}
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
