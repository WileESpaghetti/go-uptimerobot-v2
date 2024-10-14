package monitors

import (
	"encoding/json"
	"time"
)

// ResponseTimeEntry is the result of an uptime check and can be used for checking for latency over time
type ResponseTimeEntry struct {
	// Datetime is provided as a Unix timestamp in API responses
	Datetime time.Time

	// Value is provided as the number of milliseconds in API responses
	Value time.Duration
}

// jsonResponseTimeEntry is an intermediate form that we will use to convert between more strict time formats
type jsonResponseTimeEntry struct {
	UnixTimestamp int64 `json:"datetime"`
	ValueMs       int64 `json:"value"`
}

// MarshalJSON will convert the time formats between what the API uses and what the API uses
func (r *ResponseTimeEntry) MarshalJSON() ([]byte, error) {
	jsonResponseTimeEntry := jsonResponseTimeEntry{
		UnixTimestamp: r.Datetime.Unix(),
		ValueMs:       r.Value.Milliseconds(),
	}
	return json.Marshal(jsonResponseTimeEntry)
}

// UnmarshalJSON is needed to convert between an untyped API response to use more accurate [time] types
func (r *ResponseTimeEntry) UnmarshalJSON(b []byte) error {
	var raw jsonResponseTimeEntry

	err := json.Unmarshal(b, &raw)
	if err != nil {
		return err
	}

	r.Datetime = time.Unix(raw.UnixTimestamp, 0)
	r.Value = time.Duration(raw.ValueMs) * time.Millisecond

	return nil
}
