package facade

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog"
)

type apiLogger struct {
	Logger zerolog.Logger
}

type apiLoggerEntry struct {
	event *zerolog.Event
}

// TODO: sanitize and log requests and/or responses
func (l *apiLogger) NewLogEntry(r *http.Request) middleware.LogEntry {
	entry := &apiLoggerEntry{event: l.Logger.Info()}

	if rec := recover(); rec != nil {
		entry = &apiLoggerEntry{event: l.Logger.Error()}
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
	entry.event = entry.event.Fields(logFields)

	return entry
}

func (l *apiLoggerEntry) Write(status, bytes int, header http.Header, elapsed time.Duration, extra any) {
	l.event.Fields(map[string]interface{}{
		"res_length":       bytes,
		"res_status":       status,
		"res_time_elapsed": strconv.FormatFloat(float64(elapsed.Nanoseconds())/1000000.0, 'f', -1, 64) + "ms",
	}).Msg("endpoint hit")
}

func (l *apiLoggerEntry) Panic(v interface{}, stack []byte) {
	l.event = l.event.Fields(map[string]interface{}{
		"stack": string(stack),
		"panic": fmt.Sprintf("%+v", v),
	})
	l.event.Msg("request failed")
}
