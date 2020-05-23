package log

import "fmt"

type BaseLogger struct{
	Level Level
}

type LogrusLogger struct {
	BaseLogger
}

func (logger *LogrusLogger) Panic(args ...interface{})  {
	fmt.Println("Logrus PANIC")
}


func (logger *LogrusLogger) Fatal(args ...interface{})  {
	fmt.Println("Logrus Fatal")
}


func (logger *LogrusLogger) Error(args ...interface{})  {
	fmt.Println("Logrus Error")
}


func (logger *LogrusLogger) Warn(args ...interface{})  {
	fmt.Println("Logrus Warn")
}

func (logger *LogrusLogger) Info(args ...interface{})  {
	fmt.Println("Logrus Info")
}


func (logger *LogrusLogger) Debug(args ...interface{})  {
	fmt.Println("Logrus Debug")
	if len(args) == 1 {
		switch v := args[0].(type) { // type switch
		case *Entry:
			fmt.Println(v.Fields)
		}

	}

}

func (logger *LogrusLogger) Trace(args ...interface{})  {
	fmt.Println("Logrus Trace")

	switch v := args[0].(type) {
	case *Entry:
		fmt.Println(v.Fields)
	}
}

func (logger *LogrusLogger) newEntry(fields Fields) *Entry {
	return NewEntry(logger, fields)
}

func (logger *LogrusLogger) WithFields(fields Fields) *Entry {
	 return  logger.newEntry(fields)
}

func (logger *LogrusLogger) SetLevel(level Level)  {
	logger.Level = level
}

func (logger *LogrusLogger) GetLevel() Level  {

	return logger.Level
}

func NewLogrusLogger() *LogrusLogger {

	return &LogrusLogger{
		BaseLogger: BaseLogger{
			Level:DebugLevel,
		},
		//Level:DebugLevel, //Default
	}
}


