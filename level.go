package logz

import (
	"strings"

	"github.com/go-kit/kit/log/level"
)

// Level represents the log level of a logger
type Level string

// All specifies all messages should be written. It is an alias for Debug.
const All = Level("all")
// Error specifies that only messages of Level Error are written
const Error = Level("error")
// Warn specifies that only messages of Level Error or Warn are written
const Warn = Level("warn")
// Debug specifies that only messages of Level Info, Warn, Debug or Error are written
const Debug = Level("debug")
// Error specifies that only messages of Level Info, Warn, or Error are written
const Info = Level("info")

func ParseLevel(lvl string) Level {

	switch strings.ToLower(strings.TrimSpace(lvl)) {
	case "error":
		return Error
	case "warn":
		return Warn
	case "debug":
		return Debug
	case "all":
		return All
	default:
		return Info
	}
}

func (l Level) Option() level.Option {
	switch l {
	case Error:
		return level.AllowError()
	case Warn:
		return level.AllowWarn()
	case Debug:
		return level.AllowDebug()
	case All:
		return level.AllowAll()
	default:
		return level.AllowInfo()
	}
}
