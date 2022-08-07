package facade

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog"
)

type StructuredLogger struct {
	Logger zerolog.Logger
}

type StructuredLoggerEntry struct {
	Logger *zerolog.Event
}

func NewStructuredLogger(logger zerolog.Logger) func(next http.Handler) http.Handler {
	return middleware.RequestLogger(&StructuredLogger{logger})
}

func (l *StructuredLogger) NewLogEntry(r *http.Request) middleware.LogEntry {
	entry := &StructuredLoggerEntry{Logger: l.Logger.Info()}

	if rec := recover(); rec != nil {
		entry = &StructuredLoggerEntry{Logger: l.Logger.Error()}
	}

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}

	logFields := map[string]interface{}{}

	if reqID := middleware.GetReqID(r.Context()); reqID != "" {
		logFields["req_id"] = reqID
	}
	logFields["req_method"] = r.Method
	logFields["req_address"] = r.RemoteAddr
	logFields["req_uri"] = fmt.Sprintf("%s://%s%s", scheme, r.Host, r.RequestURI)
	entry.Logger = entry.Logger.Fields(logFields)

	return entry
}

func (l *StructuredLoggerEntry) Write(status, bytes int, header http.Header, elapsed time.Duration, extra any) {
	l.Logger = l.Logger.Fields(map[string]interface{}{
		"res_length":       bytes,
		"res_status":       status,
		"res_time_elapsed": strconv.FormatFloat(float64(elapsed.Nanoseconds())/1000000.0, 'f', -1, 64) + "ms",
	})
	l.Logger.Msg("")
}

func (l *StructuredLoggerEntry) Panic(v interface{}, stack []byte) {
	l.Logger = l.Logger.Fields(map[string]interface{}{
		"stack": string(stack),
		"panic": fmt.Sprintf("%+v", v),
	})
	l.Logger.Msg("request failed")
}
