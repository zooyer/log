package log

var std = NewLog(defaultLogger())

func defaultLogger() Logger {
	var logger = NewLogger()
	stdout := NewStdoutRecorder(FormatColorText)
	stderr := NewStdoutRecorder(FormatColorText)
	logger.SetDefault(stdout)
	logger.SetRecorder(stdout, "DEBUG", "INFO")
	logger.SetRecorder(stderr, "WARNING", "ERROR")

	return logger
}

func Debug(v ...interface{}) {
	std.Debug(v...)
}

func Info(v ...interface{}) {
	std.Info(v...)
}

func Warning(v ...interface{}) {
	std.Warning(v...)
}

func Error(v ...interface{}) {
	std.Error(v...)
}

func D(v ...interface{}) {
	std.Debug(v...)
}

func I(v ...interface{}) {
	std.Info(v...)
}

func W(v ...interface{}) {
	std.Warning(v...)
}

func E(v ...interface{}) {
	std.Error(v...)
}
