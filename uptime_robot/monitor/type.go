package monitor

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrUnknownType = errors.New("unknown monitor type")

// Type of monitor
type Type int64

// Type ID provided by the UptimeRobot API
const (
	TypeHttp    Type = 1
	TypeKeyword Type = 2
	TypePing    Type = 3
	TypePort    Type = 4
)

// String provides a user-friendly name for the monitor Type
func (t Type) String() string {
	switch t {
	case TypeHttp:
		return "HTTP"
	case TypeKeyword:
		return "Keyword"
	case TypePing:
		return "Ping"
	case TypePort:
		return "Port"
	default:
		return "Unknown"
	}
}

func NewType(s string) (Type, error) {
	if s == "" {
		return Type(0), ErrUnknownType
	}

	// convert by Type ID
	i, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		return Type(i), nil
	}

	// convert from type name
	for typeID := 1; ; typeID++ {
		guessedType := Type(typeID)

		if bytes.EqualFold([]byte("unknown"), []byte(guessedType.String())) {
			// we have exceeded the number of actual monitor types and haven't found a match
			return Type(0), ErrUnknownType
		}

		if bytes.EqualFold([]byte(s), []byte(guessedType.String())) {
			return guessedType, nil
		}
	}
}

////////////////////////////////////////////////////////////

type Types []Type

func (ts *Types) String() string {
	s, err := ts.MarshalText()
	if err != nil {
		return ""
	}

	return string(s)
}

func (ts *Types) MarshalText() ([]byte, error) {
	var ids []string

	for _, monitor := range *ts {
		id := strconv.FormatInt(int64(monitor), 10)
		ids = append(ids, id)
	}

	combined := strings.Join(ids, "-")

	return []byte(combined), nil
}

func (ts *Types) UnmarshalText(text []byte) error {
	fmt.Println(string(text))
	textIDs := strings.Split(string(text), "-")

	for _, sID := range textIDs {
		id, err := strconv.ParseInt(sID, 10, 64)
		if err != nil {
			return errors.New("monitor type must be an integer")
		}

		*ts = append(*ts, Type(id))
	}

	return nil
}

// Set is used to create a list of Type from a dash-separated list of ID
func (ts *Types) Set(s string) error {
	return ts.UnmarshalText([]byte(s))
}

func (ts *Types) GetSlice() *[]Type {
	return (*[]Type)(ts)
}
