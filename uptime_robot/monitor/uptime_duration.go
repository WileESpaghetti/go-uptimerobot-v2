package monitor

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	indexUp     = 0
	indexDown   = 1
	indexPaused = 2
)

var ErrInvalidUptimeDurationFormat = errors.New("format for uptime duration must be \"up-down-paused\"")

type UptimeDuration struct {
	Up     time.Duration
	Down   time.Duration
	Paused time.Duration
}

func (ud *UptimeDuration) UnmarshalJSON(b []byte) error {

	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		*ud = UptimeDuration{}
		return fmt.Errorf("value must be a string: %w", ErrInvalidUptimeDurationFormat)
	}

	strDurations := strings.Split(s, "-")
	if len(strDurations) != 3 {
		*ud = UptimeDuration{}
		return fmt.Errorf("wrong number of segments: %w", ErrInvalidUptimeDurationFormat)
	}

	for i, d := range strDurations {
		seconds, err := strconv.ParseInt(d, 10, 64)
		if err != nil {
			*ud = UptimeDuration{}
			return fmt.Errorf("%w: duration \"%s\" must be an integer: %w", err, d, ErrInvalidUptimeDurationFormat)
		}

		switch i {
		case indexUp:
			ud.Up = time.Duration(seconds) * time.Second
		case indexDown:
			ud.Down = time.Duration(seconds) * time.Second
		case indexPaused:
			ud.Paused = time.Duration(seconds) * time.Second
		}
	}

	return nil
}
