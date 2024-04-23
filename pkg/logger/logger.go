package logger

import (
	"github.com/charmbracelet/log"
	"github.com/muesli/termenv"
	"log/slog"
	"os"
)

// Log is the global charmbracelet/log logger
var Log *log.Logger

// Slog is the global slog logger provided by charmbracelet/log implementation
var Slog *slog.Logger

func InitCharmLogger(verbosity int) {

	// Create a new logger with options
	l := log.NewWithOptions(os.Stdout, charmOpts())
	// Set the styles
	l.SetStyles(charmStyles())
	// Set the color profile
	l.SetColorProfile(termenv.TrueColor)
	// Set the log level
	l.SetLevel(charmLevel(verbosity))

	// Instatiate the global log and slog
	Log = l
	Slog = slog.New(l)
}

// Logging options for charmbracelet/log
func charmOpts() log.Options {
	return log.Options{
		ReportTimestamp: false,
		ReportCaller:    false,
	}
}

func charmStyles() *log.Styles {
	// Default styles
	styles := log.DefaultStyles()
	return styles
}

func charmLevel(verbosity int) log.Level {
	var level log.Level
	switch verbosity {
	case 1:
		level = log.ErrorLevel // Least verbose
	case 2:
		level = log.WarnLevel
	case 3:
		level = log.InfoLevel
	case 4:
		level = log.DebugLevel
	default:
		level = log.ErrorLevel
	}
	return level
}
