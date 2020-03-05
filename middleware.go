package pallet

import (
	"context"
	"log"
	"net/http"

	"cloud.google.com/go/logging"
)

// contextyKey is used to set context keys.
type contextKey string

// ContextCloudTrace is the context key for response data from routes to middleware.
const ContextCloudTrace contextKey = "cloudTraceID"

// ContextLoggingClient is the context key for response data from routes to middleware.
const ContextLoggingClient contextKey = "palletLoggingClient"

/*
// GetCloudTraceIDFromContext returns the XCloudTraceContent value from the context.
func GetCloudTraceIDFromContext(ctx context.Context) string {
	var traceid string
	if ctx != nil {
		if traceid, ok := ctx.Value(ContextCloudTrace).(string); ok {
			return traceid
		}
	}
	return traceid
}

// CloudTraceContextMiddleware sets the context variables for cloudtrace
func CloudTraceContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), ContextCloudTrace, r.Header.Get("X-Cloud-Trace-Context"))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
*/

// GetLoggerFromContext returns the XCloudTraceContent value from the context.
func GetLoggerFromContext(ctx context.Context) *logging.Client {
	if ctx != nil {
		if client, ok := ctx.Value(ContextLoggingClient).(*logging.Client); ok {
			return client
		}
	}
	return &logging.Client{}
}

// Flush will flush the active logger.
func (l *Logger) Flush() error {
	if l.Logger != nil {
		return l.Logger.Flush()
	}
	return nil
}

// DefaultMiddleware sets the context variables for cloudtrace
func DefaultMiddleware(options *Options) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Create the new logger
			logger, err := NewLoggerForRequest(r, options)
			if err != nil {
				logger.Errorf("error create new logger: %v", err)
			}
			defer func() {
				err := logger.Close()
				if err != nil {
					log.Printf("error closing logger: %v", err)
				}
			}()

			ctx := context.WithValue(r.Context(), ContextLoggingClient, logger)
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
	/*
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		    var startTime = time.Now()

			loggingWrapper := client.Logger(
				parentLogName,
				logging.CommonLabels(map[string]strin\g{
					"commonLabelParent": "commonLabelParentValue",
				}),
				logging.CommonResource(&mrpb.MonitoredResource{
					Labels: labels,
					Type: "global",
				}),
			)

				ctx := context.WithValue(r.Context(), ContextCloudTrace, r.Header.Get("X-Cloud-Trace-Context"))
				next.ServeHTTP(w, r.WithContext(ctx))
			})
	*/
}
