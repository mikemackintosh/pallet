package pallet

import (
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/logging"
)

// Logf logs with the given severity. Remaining arguments are handled in the manner of fmt.Printf.
func (l *Logger) Logf(severity logging.Severity, format string, v ...interface{}) {
	if l.Logger == nil {
		log.Printf(format, v...)
		return
	}

	monRes := *l.MonitoredResource
	entry := logging.Entry{
		Timestamp: time.Now(),
		Severity:  severity,
		Payload:   fmt.Sprintf(format, v...),
		Resource:  &monRes,
	}
	fmt.Printf("entry: %#v\n", entry)

	l.Logger.Log(entry)
}

// Log logs with the given severity. v must be either a string, or something that
// marshals via the encoding/json package to a JSON object (and not any other type
// of JSON value).
func (l *Logger) Log(severity logging.Severity, v interface{}) {
	if l.Logger == nil {
		log.Print(v)
		return
	}

	monRes := *l.MonitoredResource
	entry := logging.Entry{
		Timestamp: time.Now(),
		Severity:  severity,
		Payload:   v,
		Resource:  &monRes,
	}
	fmt.Printf("entry: %#v\n", entry)

	l.Logger.Log(entry)
}

// Debugf calls Logf with debug severity.
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.Logf(logging.Debug, format, v...)
}

// Debug calls Log with debug severity.
func (l *Logger) Debug(v interface{}) {
	l.Log(logging.Debug, v)
}

// Infof calls Logf with info severity.
func (l *Logger) Infof(format string, v ...interface{}) {
	l.Logf(logging.Info, format, v...)
}

// Info calls Log with info severity.
func (l *Logger) Info(v interface{}) {
	l.Log(logging.Info, v)
}

// Noticef calls Logf with notice severity.
func (l *Logger) Noticef(format string, v ...interface{}) {
	l.Logf(logging.Notice, format, v...)
}

// Notice calls Log with notice severity.
func (l *Logger) Notice(v interface{}) {
	l.Log(logging.Notice, v)
}

// Warningf calls Logf with warning severity.
func (l *Logger) Warningf(format string, v ...interface{}) {
	l.Logf(logging.Warning, format, v...)
}

// Warning calls Log with warning severity.
func (l *Logger) Warning(v interface{}) {
	l.Log(logging.Warning, v)
}

// Errorf calls Logf with error severity.
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.Logf(logging.Error, format, v...)
}

// Error calls Log with error severity.
func (l *Logger) Error(v interface{}) {
	l.Log(logging.Error, v)
}

// Criticalf calls Logf with critical severity.
func (l *Logger) Criticalf(format string, v ...interface{}) {
	l.Logf(logging.Critical, format, v...)
}

// Critical calls Log with critical severity.
func (l *Logger) Critical(v interface{}) {
	l.Log(logging.Critical, v)
}

// Alertf calls Logf with alert severity.
func (l *Logger) Alertf(format string, v ...interface{}) {
	l.Logf(logging.Alert, format, v...)
}

// Alert calls Log with alert severity.
func (l *Logger) Alert(v interface{}) {
	l.Log(logging.Alert, v)
}

// Emergencyf calls Logf with emergency severity.
func (l *Logger) Emergencyf(format string, v ...interface{}) {
	l.Logf(logging.Emergency, format, v...)
}

// Emergency calls Log with emergency severity.
func (l *Logger) Emergency(v interface{}) {
	l.Log(logging.Emergency, v)
}
