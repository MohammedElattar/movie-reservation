package logger

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/MohammedElattar/movie-reservation/internal/config"
)

type StdoutLogger struct {
	level          LogLevel
	location       *time.Location
	datetimeFormat string
}

func newStdoutLogger(level LogLevel, location *time.Location, datetimeFormat string) *StdoutLogger {
	return &StdoutLogger{
		level:          level,
		location:       location,
		datetimeFormat: datetimeFormat,
	}
}

func (l *StdoutLogger) Debug(msg string, fields ...LogField) {
	if l.level > DebugLevel {
		return
	}

	printStructuredMessage(config.DebugLevel, msg, l.location, l.datetimeFormat, fields...)
}

func (l *StdoutLogger) Info(msg string, fields ...LogField) {
	if l.level > InfoLevel {
		return
	}

	printStructuredMessage(config.InfoLevel, msg, l.location, l.datetimeFormat, fields...)
}

func (l *StdoutLogger) Warn(msg string, fields ...LogField) {
	if l.level > WarnLevel {
		return
	}

	printStructuredMessage(config.WarnLevel, msg, l.location, l.datetimeFormat, fields...)
}

func (l *StdoutLogger) Error(msg string, fields ...LogField) {
	if l.level > ErrorLevel {
		return
	}

	printStructuredMessage(config.ErrorLevel, msg, l.location, l.datetimeFormat, fields...)
}

func printStructuredMessage(
	level config.LogLevel,
	msg string,
	location *time.Location,
	timeFormat string,
	fields ...LogField,
) {
	var output string

	if len(fields) > 0 {
		fbytes, err := json.MarshalIndent(fields, "", " ")

		if err != nil {
			fmt.Printf("unable to indent fields in log with message %s\n", msg)
			output = fmt.Sprintf(
				"%v [%v] %v\n",
				time.Now().In(location).Format(timeFormat),
				strings.ToUpper(string(level)),
				msg,
			)
		} else {
			output = fmt.Sprintf(
				"%v [%v] %v %v\n",
				time.Now().In(location).Format(timeFormat),
				strings.ToUpper(string(level)),
				msg,
				string(fbytes),
			)
		}
	} else {
		output = fmt.Sprintf(
			"%v [%v] %v\n",
			time.Now().In(location).Format(timeFormat),
			strings.ToUpper(string(level)),
			msg,
		)
	}

	fmt.Println(output)
}
