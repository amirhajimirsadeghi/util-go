package logger

import (
	"fmt"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

type logger struct {
	entry         *logrus.Entry
	initialFields logrus.Fields
}

// globalLogger is the logrus.Entry item which holds all globally stored fields
var globalLogger *logger

// parseFields takes in a list of key-value pairs and returns a logrus.Fields object
func parseFields(fields ...interface{}) logrus.Fields {
	// If client does not provide an even number of fields, panic
	if len(fields)%2 != 0 {
		logrus.Panicf("Provided field values are invalid: %+v\n", fields)
	}

	// Create fields array
	f := make(map[string]interface{}, len(fields)/2)
	for i := 0; i < len(fields); i += 2 {
		key := fmt.Sprintf("%v", fields[i])
		f[key] = fields[i+1]
	}

	return logrus.Fields(f)
}

// Initialize the global logger
func Initialize(fields ...interface{}) {
	InitializeWithWriter(os.Stdout, fields...)
}

// Initialize the global logger with a custom writer
func InitializeWithWriter(out io.Writer, fields ...interface{}) {
	if globalLogger != nil {
		return
	}

	// Set default logger parameters
	logrus.SetOutput(out)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.TraceLevel)

	// Initialize globalLogger
	globalLogger = &logger{
		entry:         logrus.WithFields(parseFields(fields...)),
		initialFields: parseFields(fields...),
	}
}

// Reset resets the global logger
func Reset() {
	globalLogger = &logger{
		entry:         logrus.WithFields(globalLogger.initialFields),
		initialFields: globalLogger.initialFields,
	}
}

// Implement select Logrus logging functions
func Log(level logrus.Level, args ...interface{}) {
	globalLogger.entry.Log(level, args...)
}

func Logf(level logrus.Level, format string, args ...interface{}) {
	globalLogger.entry.Logf(level, format, args...)
}

func Logln(level logrus.Level, args ...interface{}) {
	globalLogger.entry.Logln(level, args...)
}

func Print(args ...interface{}) {
	globalLogger.entry.Print(args...)
}

func Printf(fmtString string, args ...interface{}) {
	globalLogger.entry.Printf(fmtString, args...)
}

func Println(args ...interface{}) {
	globalLogger.entry.Println(args...)
}

func Fatal(args ...interface{}) {
	globalLogger.entry.Fatal(args...)
}

func Fatalf(fmtString string, args ...interface{}) {
	globalLogger.entry.Fatalf(fmtString, args...)
}

func Fatalln(args ...interface{}) {
	globalLogger.entry.Fatalln(args...)
}

func Panic(args ...interface{}) {
	globalLogger.entry.Panic(args...)
}

func Panicf(fmtString string, args ...interface{}) {
	globalLogger.entry.Panicf(fmtString, args...)
}

func Panicln(args ...interface{}) {
	globalLogger.entry.Panicln(args...)
}

// AddFields adds more fields to globalLogger
func AddFields(fields ...interface{}) {
	globalLogger.entry = globalLogger.entry.WithFields(parseFields(fields...))
}

// {Log level}w functions takes in a static msg string and a list of key-value
// pairs to be included as log fields and emits them through the default logger
func Logw(level logrus.Level, msg string, fields ...interface{}) {
	globalLogger.entry.WithFields(parseFields(fields...)).Log(level, msg)
}

func Tracew(msg string, fields ...interface{}) {
	globalLogger.entry.WithFields(parseFields(fields...)).Trace(msg)
}

func Debugw(msg string, fields ...interface{}) {
	globalLogger.entry.WithFields(parseFields(fields...)).Debug(msg)
}

func Infow(msg string, fields ...interface{}) {
	globalLogger.entry.WithFields(parseFields(fields...)).Info(msg)
}

func Warnw(msg string, fields ...interface{}) {
	globalLogger.entry.WithFields(parseFields(fields...)).Warn(msg)
}

func Errorw(msg string, fields ...interface{}) {
	globalLogger.entry.WithFields(parseFields(fields...)).Error(msg)
}

func Fatalw(msg string, fields ...interface{}) {
	globalLogger.entry.WithFields(parseFields(fields...)).Fatal(msg)
}

func Panicw(msg string, fields ...interface{}) {
	globalLogger.entry.WithFields(parseFields(fields...)).Panic(msg)
}
