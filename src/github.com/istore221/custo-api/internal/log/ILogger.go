package log

import (
	"time"
)

type Level uint32

type Fields map[string]interface{}

const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel Level = iota
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	TraceLevel
)

// Logger interface
type ILogger interface {
	Panic(args ...interface{})
	Fatal(args ...interface{})
	Error(args ...interface{})
	Warn(args ...interface{})
	Info(args ...interface{})
	Debug(args ...interface{})
	Trace(args ...interface{})
	WithFields(args Fields) *Entry
	SetLevel(level Level)
	GetLevel() Level

}



// An entry is the final or intermediate  logging entry. It contains all
// the fields passed with WithField{,s}. It's finally logged when Trace, Debug,
// Info, Warn, Error, Fatal or Panic is called on it. These objects can be
// reused and passed around as much as you wish to avoid field duplication.
type Entry struct {

	Logger ILogger

	// log fields
	Fields Fields

	// Time at which the log entry was created
	Time time.Time

	// Level the log entry was logged at: Trace, Debug, Info, Warn, Error, Fatal or Panic
	// This field will be set on entry firing and the value will be equal to the one in Logger struct field.
	Level Level

	// Message passed to Trace, Debug, Info, Warn, Error, Fatal or Panic
	Message string

}

func (entry *Entry) Panic(args ...interface{})  {

	if entry.supportLog(PanicLevel){
		entry.Logger.Panic(entry)
	}

}


func (entry *Entry) Fatal(args ...interface{})  {

	if entry.supportLog(FatalLevel){
		entry.Logger.Fatal(entry)
	}

}


func (entry *Entry) Error(args ...interface{})  {

	if entry.supportLog(ErrorLevel){
		entry.Logger.Error(entry)
	}

}


func (entry *Entry) Warn(args ...interface{})  {

	if entry.supportLog(WarnLevel){
		entry.Logger.Warn(entry)
	}

}

func (entry *Entry) Info(args ...interface{})  {

	if entry.supportLog(InfoLevel){
		entry.Logger.Info(entry)
	}

}


func (entry *Entry) Debug(args ...interface{})  {

	if entry.supportLog(DebugLevel){
		entry.Fields["msg"] = args[0]
		entry.Logger.Debug(entry)
	}


}

func (entry *Entry) Trace(args ...interface{})  {

	if entry.supportLog(TraceLevel){
		entry.Logger.Trace(entry)
	}

}

func (entry *Entry) supportLog(level Level) bool{
	return entry.Logger.GetLevel() >= level
}

func NewEntry(logger ILogger, fields Fields) *Entry {
	return &Entry{Logger:logger, Fields:fields}
}



