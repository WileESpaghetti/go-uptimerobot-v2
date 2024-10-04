package api

// stat exists only for JSON responses to show if any records are returned or not.
const (
	StatFail = "fail"
	StatOk   = "ok"
)

// Envelope contains common API response metadata
type Envelope struct {
	Stat  string `json:"stat"`
	Error Error  `json:"error,omitempty"`
	// TODO pagination
}
