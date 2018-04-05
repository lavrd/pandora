package log

import (
	"io"
	"path/filepath"
	"runtime"

	"github.com/sirupsen/logrus"
	ufp "github.com/spacelavr/pandora/pkg/utils/filepath"
)

var logger *logrus.Logger

// Init initialize logger
func Init(verbose bool) {

	logger = logrus.New()

	if verbose {
		logger.SetLevel(logrus.DebugLevel)
	}
}

// SetOut set output stream
func SetOut(out io.Writer) {
	logger.Out = out
}

// Debug print debug log
func Debug(args ...interface{}) {
	prepare().Debug(args)
}

// Error print error log
func Error(args ...interface{}) {
	prepare().Error(args)
}

func prepare() *logrus.Entry {

	if pc, file, line, ok := runtime.Caller(2); ok {
		fname := runtime.FuncForPC(pc).Name()

		return logger.WithFields(logrus.Fields{
			"file":  ufp.PKG(file),
			"fname": filepath.Base(fname),
			"line":  line,
		})
	}
	return logger.WithFields(logrus.Fields{})
}
