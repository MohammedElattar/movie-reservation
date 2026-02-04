// Package logger
package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/MohammedElattar/movie-reservation/internal/ports/logger"
)

type StdoutLogger struct {
	level          logger.Level
	location       *time.Location
	datetimeFormat string
}

func NewStdoutLogger(
	level logger.Level,
	location *time.Location,
	datetimeFormat string,
) *StdoutLogger {
	return &StdoutLogger{
		level:          level,
		location:       location,
		datetimeFormat: datetimeFormat,
	}
}

func (l *StdoutLogger) Debug(msg string, fields ...logger.LogField) {
	l.log(logger.DebugLevel, msg, l.location, l.datetimeFormat, fields...)
}

func (l *StdoutLogger) Info(msg string, fields ...logger.LogField) {
	l.log(logger.InfoLevel, msg, l.location, l.datetimeFormat, fields...)
}

func (l *StdoutLogger) Warn(msg string, fields ...logger.LogField) {
	l.log(logger.WarnLevel, msg, l.location, l.datetimeFormat, fields...)
}

func (l *StdoutLogger) Error(msg string, fields ...logger.LogField) {
	l.log(logger.ErrorLevel, msg, l.location, l.datetimeFormat, fields...)
}

func (l *StdoutLogger) log(
	level logger.Level,
	msg string,
	location *time.Location,
	timeFormat string,
	fields ...logger.LogField,
) {
	if level < l.level {
		return
	}

	var output string

	if len(fields) > 0 {
		fbytes, err := json.MarshalIndent(fields, "", " ")

		if err != nil {
			fmt.Fprintf(os.Stdout, "unable to indent fields in log with message %s\n", msg)
			output = fmt.Sprintf(
				"%v [%v] %v\n",
				time.Now().In(location).Format(timeFormat),
				level.String(),
				msg,
			)
		} else {
			output = fmt.Sprintf(
				"%v [%v] %v %v\n",
				time.Now().In(location).Format(timeFormat),
				level.String(),
				msg,
				string(fbytes),
			)
		}
	} else {
		output = fmt.Sprintf(
			"%v [%v] %v\n",
			time.Now().In(location).Format(timeFormat),
			level.String(),
			msg,
		)
	}

	fmt.Fprint(os.Stdout, output)
}
