package monitors

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrUnknownStatus = errors.New("unknown monitor status")

type Status int64

// Supported Status ID's
const (
	StatusPaused     Status = 0
	StatusNotChecked Status = 1
	StatusUp         Status = 2
	StatusSeemsDown  Status = 8
	StatusDown       Status = 9
)

func (st Status) String() string {
	switch st {
	case StatusPaused:
		return "Paused"
	case StatusNotChecked:
		return "Not Checked"
	case StatusUp:
		return "Up"
	case StatusSeemsDown:
		return "Seems Down"
	case StatusDown:
		return "Down"
	default:
		return "Unknown"
	}
}

func NewStatus(s string) (Status, error) {
	if s == "" {
		return Status(0), ErrUnknownStatus
	}

	// convert by Status ID
	i, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		return Status(i), nil
	}

	// convert from status name
	for statusID := 1; ; statusID++ {
		guessedStatus := Status(statusID)

		if bytes.EqualFold([]byte("unknown"), []byte(guessedStatus.String())) {
			// we have exceeded the number of actual monitor statuses and haven't found a match
			return Status(0), ErrUnknownStatus
		}

		if bytes.EqualFold([]byte(s), []byte(guessedStatus.String())) {
			return guessedStatus, nil
		}
	}
}

////////////////////////////////////////////////////////////
// TODO see if we can use either generics or something like Status > Enumable > int64 and just have status override String() and some type of Equals(string)

type Statuses []Status

func (ss *Statuses) String() string {
	s, err := ss.MarshalText()
	if err != nil {
		return ""
	}

	return string(s)
}

func (ss *Statuses) MarshalText() ([]byte, error) {
	var ids []string

	for _, monitor := range *ss {
		id := strconv.FormatInt(int64(monitor), 10)
		ids = append(ids, id)
	}

	combined := strings.Join(ids, "-")

	return []byte(combined), nil
}

func (ss *Statuses) UnmarshalText(text []byte) error {
	fmt.Println(string(text))
	textIDs := strings.Split(string(text), "-")

	for _, sID := range textIDs {
		id, err := strconv.ParseInt(sID, 10, 64)
		if err != nil {
			return errors.New("monitor status must be an integer")
		}

		*ss = append(*ss, Status(id))
	}

	return nil
}

// Set is used to create a list of Status from a dash-separated list of ID
func (ss *Statuses) Set(s string) error {
	return ss.UnmarshalText([]byte(s))
}
