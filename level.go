package log

// Level type of log level
type Level int

const (
	// LevelDebug logging in debug level
	LevelDebug Level = iota

	// LevelInfo logging in info level
	LevelInfo

	// LevelWarning logging in warning level
	LevelWarning

	// LevelError logging in error level
	LevelError

	// LevelFatal logging in fatal level
	LevelFatal
)

// String returns name of level
func (lvl Level) String() string {
	switch lvl {
	case LevelDebug:
		return "Debug"
	case LevelInfo:
		return "Info"
	case LevelWarning:
		return "Warning"
	case LevelError:
		return "Error"
	case LevelFatal:
		return "Fatal"
	default:
		return "Unknown"
	}
}
