package openlogging

var logger Logger

func SetLogger(l Logger) {
	logger = l
}

func GetLogger() Logger {
	return logger
}
