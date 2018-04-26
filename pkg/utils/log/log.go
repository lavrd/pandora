package log

import (
	"fmt"
	"io"
	"path/filepath"
	"runtime"

	"github.com/sirupsen/logrus"
	ufp "github.com/spacelavr/pandora/pkg/utils/filepath"
)

var (
	// CommonLogFormat http request log format
	// 127.0.0.1 - - [Sun, 08 Apr 2018 06:50:15 +0000] "GET /health HTTP/1.1" 501 40 1.0019ms curl
	CommonLogFormat = "%s %s %s [%s] \"%s %s %v\" %d %d %s %s\n"
)

func init() {
	logrus.SetLevel(logrus.ErrorLevel)
}

// SetVerbose set verbose output
func SetVerbose(verbose bool) {
	if verbose {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

// SetOut set output stream
func SetOut(out io.Writer) {
	logrus.SetOutput(out)
}

// Debug print debug log
func Debug(args ...interface{}) {
	prepare().Debug(args...)
}

// Debugf print formatted debug log
func Debugf(format string, args ...interface{}) {
	prepare().Debugf(format, args...)
}

// Errorf print formatted error log
func Errorf(format string, args ...interface{}) {
	prepare().Errorf(format, args...)
}

// Error print error log
func Error(args ...interface{}) {
	prepare().Error(args...)
}

// Fatal print fatal log
func Fatal(args ...interface{}) {
	prepare().Fatal(args...)
}

// Http print http log in common log format to out stream
func Http(out io.Writer, ip, time, method, route, proto, duration, userAgent string, code, size int) {
	fmt.Fprintf(out, CommonLogFormat, ip, "-", "-", time, method, route, proto, code, size, duration, userAgent)
}

func prepare() *logrus.Entry {
	if pc, file, line, ok := runtime.Caller(2); ok {
		fname := runtime.FuncForPC(pc).Name()

		return logrus.WithFields(logrus.Fields{
			"file":  ufp.PKG(file),
			"fname": filepath.Base(fname),
			"line":  line,
		})
	}

	return logrus.WithFields(logrus.Fields{})
}
