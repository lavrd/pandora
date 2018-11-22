package log

import (
	"log"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	StackTraceFormat = "%v%+v"
)

// StackTracer
type StackTracer interface {
	StackTrace() errors.StackTrace
}

func init() {
	logrus.SetLevel(logrus.PanicLevel)
}

// SetVerbose set verbose output
func SetVerbose(verbose bool) {
	if verbose {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

// Debug print debug log
func Debug(args ...interface{}) {
	caller().Debug(args...)
}

// Debugf print formatted debug log
func Debugf(format string, args ...interface{}) {
	caller().Debugf(format, args...)
}

// Error print error log
func Error(err error) {
	log.Printf(StackTraceFormat, err, stackTrace(err))
}

// Fatal print fatal log
func Fatal(err error) {
	log.Fatalf(StackTraceFormat, err, stackTrace(err))
}

func stackTrace(err error) errors.StackTrace {
	st, _ := err.(StackTracer)
	return st.StackTrace()[1:]
}

func caller() *logrus.Entry {
	if pc, file, line, ok := runtime.Caller(2); ok {
		return logrus.WithFields(logrus.Fields{
			"file":  file[strings.Index(file, "/pkg")+len("/pkg"):],
			"fname": filepath.Base(runtime.FuncForPC(pc).Name()),
			"line":  line,
		})
	}
	return logrus.WithFields(logrus.Fields{})
}
