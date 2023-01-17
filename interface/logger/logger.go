package logger

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
)

// implementation based on https://github.com/go-chi/chi/blob/master/_examples/logging/main.go

// StructuredLogger constructor
func NewStructuredLogger(logger *logrus.Logger) func(next http.Handler) http.Handler {
	return middleware.RequestLogger(&StructuredLogger{logger})
}

// Structured Logger required for chi logger middleware
type StructuredLogger struct {
	Logger *logrus.Logger
}

// Create log session
func (s *StructuredLogger) NewLogEntry(r *http.Request) middleware.LogEntry {
	logFields := logrus.Fields{}

	logFields["ts"] = time.Now().UTC().Format(time.RFC1123)

	if reqID := middleware.GetReqID(r.Context()); reqID != "" {
		logFields["req_id"] = reqID
	}

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}

	logFields["http_scheme"] = scheme
	logFields["http_proto"] = r.Proto
	logFields["http_method"] = r.Method
	logFields["remote_addr"] = r.RemoteAddr
	logFields["user_agent"] = r.UserAgent()
	logFields["uri"] = fmt.Sprintf("%s://%s%s", scheme, r.Host, r.RequestURI)

	entry := &StructuredLoggerEntry{Logger: logrus.NewEntry(s.Logger)}

	entry.Logger = entry.Logger.WithFields(logFields)

	entry.Logger.Info("HTTP Request started")

	return entry
}

type StructuredLoggerEntry struct {
	Logger logrus.FieldLogger
}

// write log from the log session
func (s *StructuredLoggerEntry) Write(status, bytes int, header http.Header, elapsed time.Duration, extra interface{}) {
	s.Logger = s.Logger.WithFields(logrus.Fields{
		"resp_status":       status,
		"resp_bytes_length": bytes,
		"resp_elapsed":      float64(elapsed.Nanoseconds()) / 1000000.0,
	})

	s.Logger.Info("HTTP Request completed")
}

func (s *StructuredLoggerEntry) Panic(v interface{}, stack []byte) {
	s.Logger = s.Logger.WithFields(logrus.Fields{
		"stack": string(stack),
		"panic": fmt.Sprintf("%+v", v),
	})
}

func GetLogEntry(r *http.Request) logrus.FieldLogger {
	entry := middleware.GetLogEntry(r).(*StructuredLoggerEntry)
	return entry.Logger
}

// set field for log session
func LogEntrySetField(r *http.Request, key string, value interface{}) {
	if entry, ok := r.Context().Value(middleware.LogEntryCtxKey).(*StructuredLoggerEntry); ok {
		fields := logrus.Fields{}
		fields[key] = value
		entry.Logger = entry.Logger.WithFields(fields)
	}
}

// set fields for log session
func LogEntrySetFields(r *http.Request, fields map[string]interface{}) {
	if entry, ok := r.Context().Value(middleware.LogEntryCtxKey).(*StructuredLoggerEntry); ok {
		fields := logrus.Fields{}
		for k, v := range fields {
			fields[k] = v
		}
		entry.Logger = entry.Logger.WithFields(fields)
	}
}
