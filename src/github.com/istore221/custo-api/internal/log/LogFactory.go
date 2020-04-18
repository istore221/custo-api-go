package log

func LogFactory(logger string) ILogger {

	switch logger {
	case "Mock":
		return NewMockLogger()
		break
	case "Logrus":
		return NewLogrusLogger()
		break
	default:
		return nil

	}

	return nil;

}