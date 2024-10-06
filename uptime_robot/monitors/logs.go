package monitors

import "time"

// LogType is an Enum
type LogType int64

// LogTypes
const (
	LogTypeDown    LogType = 1
	LogTypeUp      LogType = 2
	LogTypeStarted LogType = 98
	LogTypePaused  LogType = 99
)

func (lt LogType) String() string {
	switch lt {
	case LogTypeDown:
		return "Down"
	case LogTypeUp:
		return "Up"
	case LogTypeStarted:
		return "Started"
	case LogTypePaused:
		return "Paused"
	default:
		return ""
	}
}

type LogEntry struct {
	Id       int64   `json:"id"`
	Type     LogType `json:"type"`
	Datetime int     `json:"datetime"`

	Duration time.Duration `json:"duration"` // API downtime duration is in seconds
	Reason   struct {
		Code   string `json:"code"`
		Detail string `json:"detail"`
	} `json:"reason"`
}
