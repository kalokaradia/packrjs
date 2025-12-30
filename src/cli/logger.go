package cli

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
)

// LogLevel defines the severity of the log message
type LogLevel int

const (
	// Error level for errors
	Error LogLevel = iota
	// Warn level for warnings
	Warn
	// Info level for general info
	Info
	// Debug level for debugging
	Debug
)

// LogLevelFromString converts a string to LogLevel
func LogLevelFromString(level string) LogLevel {
	switch level {
	case "error":
		return Error
	case "warn":
		return Warn
	case "debug":
		return Debug
	default:
		return Info
	}
}

// Logger represents a CLI logger with levels and consistent styling
type Logger struct {
	level      LogLevel
	showTime   bool
	useIcons   bool
	colors     *color.Color
	warnColor  *color.Color
	infoColor  *color.Color
	debugColor *color.Color
	successCol *color.Color
	errorCol   *color.Color
}

// New creates a new logger with the specified log level
func New(level string) *Logger {
	lvl := LogLevelFromString(level)

	return &Logger{
		level:      lvl,
		showTime:   false,
		useIcons:   true,
		colors:     color.New(color.FgWhite),
		warnColor:  color.New(color.FgYellow),
		infoColor:  color.New(color.FgCyan),
		debugColor: color.New(color.FgWhite),
		successCol: color.New(color.FgGreen, color.Bold),
		errorCol:   color.New(color.FgRed, color.Bold),
	}
}

// WithTimestamp enables timestamp in log messages
func (l *Logger) WithTimestamp(enabled bool) *Logger {
	l.showTime = enabled
	return l
}

// WithIcons enables or disables icons in log messages
func (l *Logger) WithIcons(enabled bool) *Logger {
	l.useIcons = enabled
	return l
}

// format formats the message with optional timestamp
func (l *Logger) format(format string, args ...any) string {
	if len(args) == 0 {
		return format
	}
	return fmt.Sprintf(format, args...)
}

// getTimestamp returns the current timestamp string
func (l *Logger) getTimestamp() string {
	return time.Now().Format("15:04:05")
}

// Error prints error messages (always shown)
func (l *Logger) Error(format string, args ...any) {
	prefix := "Error"
	if l.useIcons {
		prefix = "✗ " + prefix
	}
	msg := l.format(format, args...)
	if l.showTime {
		l.errorCol.Printf("[%s] %s: %s\n", l.getTimestamp(), prefix, msg)
	} else {
		l.errorCol.Printf("%s: %s\n", prefix, msg)
	}
}

// Fatal prints error and exits with code 1
func (l *Logger) Fatal(format string, args ...any) {
	l.Error(format, args...)
	os.Exit(1)
}

// FatalErr prints error with context and exits
func (l *Logger) FatalErr(err error, context string) {
	prefix := "Fatal"
	if l.useIcons {
		prefix = "✗ " + prefix
	}
	if l.showTime {
		l.errorCol.Printf("[%s] %s: %s: %v\n", l.getTimestamp(), prefix, context, err)
	} else {
		l.errorCol.Printf("%s: %s: %v\n", prefix, context, err)
	}
	os.Exit(1)
}

// Warn prints warning if level >= Warn
func (l *Logger) Warn(format string, args ...any) {
	if l.level < Warn {
		return
	}
	prefix := "Warn"
	if l.useIcons {
		prefix = "⚠ " + prefix
	}
	msg := l.format(format, args...)
	if l.showTime {
		l.warnColor.Printf("[%s] %s: %s\n", l.getTimestamp(), prefix, msg)
	} else {
		l.warnColor.Printf("%s: %s\n", prefix, msg)
	}
}

// WarnWithTip prints warning with a helpful tip
func (l *Logger) WarnWithTip(warnMsg, tipMsg string) {
	if l.level < Warn {
		return
	}
	prefix := "Warn"
	tipPrefix := "Tip"
	if l.useIcons {
		prefix = "⚠ " + prefix
		tipPrefix = "💡 " + tipPrefix
	}
	if l.showTime {
		l.warnColor.Printf("[%s] %s: %s\n", l.getTimestamp(), prefix, warnMsg)
		l.infoColor.Printf("[%s] %s: %s\n", l.getTimestamp(), tipPrefix, tipMsg)
	} else {
		l.warnColor.Printf("%s: %s\n", prefix, warnMsg)
		l.infoColor.Printf("%s: %s\n", tipPrefix, tipMsg)
	}
}

// Info prints info if level >= Info
func (l *Logger) Info(format string, args ...any) {
	if l.level < Info {
		return
	}
	prefix := "Info"
	if l.useIcons {
		prefix = "ℹ " + prefix
	}
	msg := l.format(format, args...)
	if l.showTime {
		l.infoColor.Printf("[%s] %s: %s\n", l.getTimestamp(), prefix, msg)
	} else {
		l.infoColor.Printf("%s: %s\n", prefix, msg)
	}
}

// Success prints success messages
func (l *Logger) Success(format string, args ...any) {
	prefix := "Done"
	if l.useIcons {
		prefix = "✓ " + prefix
	}
	msg := l.format(format, args...)
	if l.showTime {
		l.successCol.Printf("[%s] %s: %s\n", l.getTimestamp(), prefix, msg)
	} else {
		l.successCol.Printf("%s: %s\n", prefix, msg)
	}
}

// Debug prints debug if level >= Debug
func (l *Logger) Debug(format string, args ...any) {
	if l.level < Debug {
		return
	}
	prefix := "Debug"
	if l.useIcons {
		prefix = "🔍 " + prefix
	}
	msg := l.format(format, args...)
	if l.showTime {
		l.debugColor.Printf("[%s] %s: %s\n", l.getTimestamp(), prefix, msg)
	} else {
		l.debugColor.Printf("%s: %s\n", prefix, msg)
	}
}

// Print prints a raw message without any formatting
func (l *Logger) Print(args ...any) {
	fmt.Print(args...)
}

// Println prints a raw message with newline
func (l *Logger) Println(args ...any) {
	fmt.Println(args...)
}

// Printf prints a formatted message
func (l *Logger) Printf(format string, args ...any) {
	fmt.Printf(format, args...)
}

// PrintSuccess prints a success banner
func (l *Logger) PrintSuccess() {
	if l.useIcons {
		l.successCol.Println("✓ Build succeeded")
	} else {
		l.successCol.Println("Build succeeded")
	}
}

// PrintError prints an error banner
func (l *Logger) PrintError(msg string) {
	if l.useIcons {
		l.errorCol.Println("✗ " + msg)
	} else {
		l.errorCol.Println(msg)
	}
}

// PrintWatch prints watch mode status
func (l *Logger) PrintWatch(path string) {
	if l.useIcons {
		l.infoColor.Printf("👀 Watching: %s\n", path)
	} else {
		l.infoColor.Printf("Watching: %s\n", path)
	}
}

// PrintRebuild prints rebuild notification
func (l *Logger) PrintRebuild() {
	if l.useIcons {
		l.infoColor.Println("↻ Rebuilding...")
	} else {
		l.infoColor.Println("Rebuilding...")
	}
}

// PrintDirCreated prints directory creation message
func (l *Logger) PrintDirCreated(path string) {
	if l.useIcons {
		l.successCol.Printf("📁 Created directory: %s\n", path)
	} else {
		l.successCol.Printf("Created directory: %s\n", path)
	}
}

// PrintBuildStart prints build start message
func (l *Logger) PrintBuildStart() {
	if l.useIcons {
		l.infoColor.Println("⚙️  Building...")
	} else {
		l.infoColor.Println("Building...")
	}
}

