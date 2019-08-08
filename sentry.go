package log_sentry

import (
	"github.com/evalphobia/logrus_sentry"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
	"fmt"
)

func enableSentry(logger *logrus.Logger, dsn string) {
	split := strings.Split(os.Args[0], "/")
	program := split[len(split)-1]
	tags := map[string]string{
		"program": program,
	}

	levels := []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
	}

	// add logrus hook
	hook, err := logrus_sentry.NewWithTagsSentryHook(dsn, tags, levels)
	fmt.Printf("err:%+v\n", err)
	if err == nil {
		logger.Hooks.Add(hook)
	}
}
