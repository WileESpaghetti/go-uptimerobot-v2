package monitors

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// LogType is an Enum
type LogType int64

// LogTypes
const (
	LogTypeDown    LogType = 1
	LogTypeUp      LogType = 2
	LogTypeStarted LogType = 98
	LogTypePaused  LogType = 99
)

var ErrUnknownLogType = errors.New("unknown monitor log type")

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
	Datetime time.Time
	Duration time.Duration // API downtime duration is in seconds
	Reason   struct {
		Code   string `json:"code"`
		Detail string `json:"detail"`
	} `json:"reason"`
}

type unencodableLogEntry LogEntry

type jsonLogEntry struct {
	unencodableLogEntry
	Datetime int64 `json:"datetime"`
	Duration int64 `json:"duration"`
}

func (jl jsonLogEntry) LogEntry() LogEntry {
	return LogEntry{
		Id:       jl.Id,
		Type:     jl.Type,
		Datetime: time.Unix(jl.Datetime, 0),
		Duration: time.Duration(jl.Duration) * time.Second,
		Reason:   jl.Reason,
	}
}

func (l *LogEntry) UnmarshalJSON(data []byte) error {
	var jm jsonLogEntry

	if err := json.Unmarshal(data, &jm); err != nil {
		return err
	}

	*l = jm.LogEntry()

	return nil
}

func (l *LogEntry) MarshalJSON() ([]byte, error) {
	jm := jsonLogEntry{
		unencodableLogEntry: unencodableLogEntry{
			Id:     l.Id,
			Type:   l.Type,
			Reason: l.Reason,
		},
		Datetime: l.Datetime.Unix(),
		Duration: int64(l.Duration.Seconds()),
	}

	return json.Marshal(jm)
}

////////////////////////////////////////////////////////////

type LogTypes []LogType

func (ss *LogTypes) String() string {
	s, err := ss.MarshalText()
	if err != nil {
		return ""
	}

	return string(s)
}

func (ss *LogTypes) GetSlice() *[]LogType {
	return (*[]LogType)(ss)
}

func (ss *LogTypes) MarshalText() ([]byte, error) {
	var ids []string

	for _, monitor := range *ss {
		id := strconv.FormatInt(int64(monitor), 10)
		ids = append(ids, id)
	}

	combined := strings.Join(ids, "-")

	return []byte(combined), nil
}

func (ss *LogTypes) UnmarshalText(text []byte) error {
	fmt.Println(string(text))
	textIDs := strings.Split(string(text), "-")

	for _, sID := range textIDs {
		id, err := strconv.ParseInt(sID, 10, 64)
		if err != nil {
			return errors.New("monitor status must be an integer")
		}

		*ss = append(*ss, LogType(id))
	}

	return nil
}

// Set is used to create a list of Status from a dash-separated list of ID
func (ss *LogTypes) Set(s string) error {
	return ss.UnmarshalText([]byte(s))
}

func NewLogType(s string) (LogType, error) {
	if len(s) == 0 {
		return LogType(0), ErrUnknownLogType
	}

	// convert by LogType ID
	i, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		return LogType(i), nil
	}

	// convert from status name
	for logTypeID := 1; ; logTypeID++ {
		guessedLogType := LogType(logTypeID)

		if bytes.EqualFold([]byte("unknown"), []byte(guessedLogType.String())) {
			// we have exceeded the number of actual monitor log types and haven't found a match
			return LogType(0), ErrUnknownLogType
		}

		if bytes.EqualFold([]byte(s), []byte(guessedLogType.String())) {
			return guessedLogType, nil
		}
	}
}
