package log

import (
	"fmt"
)



type MockLogger struct {
	Level Level
}

func (logger *MockLogger) Panic(args ...interface{})  {
	fmt.Println("Mock Panic Logger ")
}


func (logger *MockLogger) Fatal(args ...interface{})  {
	fmt.Println("Mock Fatal Logger ")
}


func (logger *MockLogger) Error(args ...interface{})  {
	fmt.Println("Mock Error Logger ")
}


func (logger *MockLogger) Warn(args ...interface{})  {
	fmt.Println("Mock Warn Logger ")
}

func (logger *MockLogger) Info(args ...interface{})  {
	fmt.Println("Mock Info Logger ")
}


func (logger *MockLogger) Debug(args ...interface{})  {
	fmt.Println("Mock Debug Logger ")
	if len(args) == 1 {
		switch v := args[0].(type) { // type switch
		case *Entry:
			fmt.Println(v.Fields)
		}

	}

}


func (logger *MockLogger) Trace(args ...interface{})  {
	fmt.Println("Mock Trace Logger ")
}

func (logger *MockLogger) WithFields(fields Fields) *Entry {
	return  logger.newEntry(fields)
}

func (logger *MockLogger) SetLevel(level Level)  {
	fmt.Println("Mock Logger SetLevel ",level)
	logger.Level = level
}


func (logger *MockLogger) newEntry(fields Fields) *Entry {
	return NewEntry(logger, fields)
}

func (logger *MockLogger) GetLevel() Level  {

	return logger.Level
}

func NewMockLogger() *MockLogger {

	return &MockLogger{
		Level:DebugLevel, //Default
	}
}
