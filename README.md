# log_sentry

usage
```
	logSentryConfig := &LogSentryConfig{
		LogLevel:	"debug",
		DSN:		"sentry_dsn",
		ReportLogLevel:	[]logrus.Level{
			logrus.PanicLevel,
			logrus.ErrorLevel,
			logrus.WarnLevel,
		},
	}
	InitLogSentry(logSentryConfig)
	Error("test")
```
