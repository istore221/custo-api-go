package log

import "fmt"

var logger ILogger = nil

func GetLogger() ILogger{

	if logger == nil {
		fmt.Println("init")
		logger = LogFactory("Logrus") // change default logger
		logger.SetLevel(WarnLevel)
	}
	return logger

}