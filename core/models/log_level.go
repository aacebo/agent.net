package models

type LogLevel string

const (
	LOG_LEVEL_INFO  LogLevel = "info"
	LOG_LEVEL_WARN  LogLevel = "warn"
	LOG_LEVEL_ERROR LogLevel = "error"
	LOG_LEVEL_DEBUG LogLevel = "debug"
)

func (self LogLevel) Valid() bool {
	switch self {
	case LOG_LEVEL_INFO:
		fallthrough
	case LOG_LEVEL_WARN:
		fallthrough
	case LOG_LEVEL_ERROR:
		fallthrough
	case LOG_LEVEL_DEBUG:
		return true
	}

	return false
}
