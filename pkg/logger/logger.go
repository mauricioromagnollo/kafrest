package logger

// Logger defines a generic logging interface with methods for various log levels.
// Each method accepts a message string and optional extra properties for structured logging.
// The props parameter allows passing additional context or metadata to be included with the log entry.
type Logger interface {
	// Info logs an informational message with optional extra properties.
	Info(msg string, props ...ExtraProps)
	// Error logs an error message with optional extra properties.
	Error(msg string, props ...ExtraProps)
	// Fatal logs a fatal error message with optional extra properties.
	Fatal(msg string, props ...ExtraProps)
	// Warn logs a warning message with optional extra properties.
	Warn(msg string, props ...ExtraProps)
	// Debug logs a debug message with optional extra properties.
	Debug(msg string, props ...ExtraProps)
}
