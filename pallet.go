// Package pallet offers a clean and simple stackdriver logging interface
// for google services.
//
/*
Grouping Logs by Request

To group all the log entries written during a single HTTP request, create two Loggers, a "parent" and a "child," with different log IDs. Both should be in the same project, and have the same MonitoredResource type and labels.

- Parent entries must have HTTPRequest.Request populated. (Strictly speaking, only the URL is necessary.)

- A child entry's timestamp must be within the time interval covered by the parent request. (i.e., before the parent.Timestamp and after the parent.Timestamp - parent.HTTPRequest.Latency. This assumes the parent.Timestamp marks the end of the request.)

- The trace field must be populated in all of the entries and match exactly.

You should observe the child log entries grouped under the parent on the console. The parent entry will not inherit the severity of its children; you must update the parent severity yourself.
*/
package pallet

import (
	"fmt"
	"net/http"
	"strings"

	"cloud.google.com/go/logging"
	"google.golang.org/genproto/googleapis/api/monitoredres"
)

const (
	logName                = "app_log"
	traceContextHeaderName = "X-Cloud-Trace-Context"
)

// Logger is a standard logging client struct.
type Logger struct {
	Client            *logging.Client
	Logger            *logging.Logger
	MonitoredResource *monitoredres.MonitoredResource
	TraceID           string
}

// Options are used to configure a client.
type Options struct {
	ProjectID string
	ServiceID string
	VersionID string
	LogType   string
	Labels    map[string]string
}

// NewOptionSet will create a new optionset, passed to the client.
func NewOptionSet() *Options {
	return &Options{
		ServiceID: "default",
		VersionID: "default",
		LogType:   "gae_app",
	}
}

// SetProject will set a project
func (o *Options) SetProject(p string) {
	o.ProjectID = p
}

// SetService will set a service
func (o *Options) SetService(s string) {
	o.ServiceID = s
}

// SetVersion will set a verson
func (o *Options) SetVersion(v string) {
	o.VersionID = v
}

// SetLabels will set labels.
func (o *Options) SetLabels(l map[string]string) {
	o.Labels = l
}

// getTraceID takes an http.Request and using the GOOGLE_CLOUD_PROJECT env var,
// generates the stackdriver trace id.
func getTraceID(project, traceid string) string {
	return fmt.Sprintf("projects/%s/traces/%s", project, traceid)
}

// NewLoggerForRequest is a helper method to create a new client
func NewLoggerForRequest(r *http.Request, options *Options, loggingOptions ...logging.LoggerOption) (*Logger, error) {
	var logger Logger
	if len(options.ProjectID) == 0 {
		return &logger, ErrorInvalidConfiguration{"ProjectID"}
	}

	if len(options.LogType) == 0 {
		return &logger, ErrorInvalidConfiguration{"LogType"}
	}

	traceContext := r.Header.Get(traceContextHeaderName)
	if len(traceContext) == 0 {
		return &logger, fmt.Errorf("missing trace context, switching to std logging")
	}

	client, err := logging.NewClient(r.Context(), fmt.Sprintf("projects/%s", options.ProjectID))
	if err != nil {
		return &logger, err
	}

	monRes := &monitoredres.MonitoredResource{
		Labels: map[string]string{
			"project_id": options.ProjectID,
			"module_id":  options.ServiceID,
			"version_id": options.VersionID,
		},
		Type: options.LogType,
	}

	if len(options.Labels) > 0 {
		for k, v := range options.Labels {
			monRes.Labels[k] = v
		}
	}

	logger = Logger{
		Client:            client,
		Logger:            client.Logger(logName, loggingOptions...),
		MonitoredResource: monRes,
		TraceID:           getTraceID(options.ProjectID, strings.Split(traceContext, "/")[0]),
	}

	return &logger, nil
}

// Close closes the Logger, ensuring all logs are flushed and closing the underlying
// Stackdriver Logging client.
func (l *Logger) Close() error {
	if l.Client != nil {
		return l.Client.Close()
	}

	return nil
}
