package logger

import (
	"fmt"
	"time"

	"github.com/MohammedElattar/movie-reservation/internal/config"
)

func CreateNewLogger(
	driver config.LoggerDriver,
	level LogLevel,
	location *time.Location,
	logTimeFormat string,
) (Logger, error) {
	switch driver {
	case config.StdoutLogger:
		return newStdoutLogger(level, location, logTimeFormat), nil
	default:
		return nil, fmt.Errorf("unknown log driver %s", driver)
	}
}
