package onvif

import (
	"flag"
	"fmt"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
	logrus_syslog "github.com/sirupsen/logrus/hooks/syslog"
	"log/syslog"
)

type Level uint8
var logger = logrus.New()

type levelFlag struct{}

// String implements flag.Value.
func (f levelFlag) String() string {
	return logger.Level.String()
}

// Set implements flag.Value.
func (f levelFlag) Set(level string) error {
	l, err := logrus.ParseLevel(level)
	if err != nil {
		return err
	}
	logger.Level = l
	return nil
}

func Set(level string) error {
	l := levelFlag{}
	return l.Set(level)
}

func init() {
	// In order for this flag to take effect, the user of the package must call
	// flag.Parse() before logging anything.
	flag.Var(levelFlag{}, "log.level", "Only log messages with the given severity or above. Valid levels: [debug, info, warn, error, fatal, panic].")

	hook, err := logrus_syslog.NewSyslogHook("", "", syslog.LOG_INFO, "")

	if err == nil {
		logger.Hooks.Add(hook)
	}
}

// fileLineEntry returns a logrus.Entry with file and line annotations for the
// original user log statement (two stack frames up from this function).
func fileLineEntry() *logrus.Entry {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[slash+1:]
		}
	}
	fileLine := fmt.Sprintf("%s:%d", file, line)
	return logger.WithFields(logrus.Fields{
		"caller": fileLine,
	})
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	fileLineEntry().Debug(args...)
}

// Debugln logs a message at level Debug on the standard logger.
func Debugln(args ...interface{}) {
	fileLineEntry().Debugln(args...)
}

// Debugf logs a message at level Debug on the standard logger.
func Debugf(format string, args ...interface{}) {
	fileLineEntry().Debugf(format, args...)
}

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {
	fileLineEntry().Info(args...)
}

// Infoln logs a message at level Info on the standard logger.
func Infoln(args ...interface{}) {
	fileLineEntry().Infoln(args...)
}

// Infof logs a message at level Info on the standard logger.
func Infof(format string, args ...interface{}) {
	fileLineEntry().Infof(format, args...)
}

// Print logs a message at level Info on the standard logger.
func Print(args ...interface{}) {
	fileLineEntry().Info(args...)
}

// Println logs a message at level Info on the standard logger.
func Println(args ...interface{}) {
	fileLineEntry().Infoln(args...)
}

// Printf logs a message at level Info on the standard logger.
func Printf(format string, args ...interface{}) {
	fileLineEntry().Infof(format, args...)
}

// Warn logs a message at level Warn on the standard logger.
func Warn(args ...interface{}) {
	fileLineEntry().Warn(args...)
}

// Warnln logs a message at level Warn on the standard logger.
func Warnln(args ...interface{}) {
	fileLineEntry().Warnln(args...)
}

// Warnf logs a message at level Warn on the standard logger.
func Warnf(format string, args ...interface{}) {
	fileLineEntry().Warnf(format, args...)
}

// Error logs a message at level Error on the standard logger.
func Error(args ...interface{}) {
	fileLineEntry().Error(args...)
}

// Errorln logs a message at level Error on the standard logger.
func Errorln(args ...interface{}) {
	fileLineEntry().Errorln(args...)
}

// Errorf logs a message at level Error on the standard logger.
func Errorf(format string, args ...interface{}) {
	fileLineEntry().Errorf(format, args...)
}

// Fatal logs a message at level Fatal on the standard logger.
func Fatal(args ...interface{}) {
	fileLineEntry().Fatal(args...)
}

// Fatalln logs a message at level Fatal on the standard logger.
func Fatalln(args ...interface{}) {
	fileLineEntry().Fatalln(args...)
}

// Fatalf logs a message at level Fatal on the standard logger.
func Fatalf(format string, args ...interface{}) {
	fileLineEntry().Fatalf(format, args...)
}

// Panic logs a message at level Panic on the standard logger.
func Panic(args ...interface{}) {
	fileLineEntry().Panicln(args...)
}

// Panicln logs a message at level Panic on the standard logger.
func Panicln(args ...interface{}) {
	fileLineEntry().Panicln(args...)
}

// Panicf logs a message at level Panic on the standard logger.
func Panicf(format string, args ...interface{}) {
	fileLineEntry().Panicf(format, args...)
}

