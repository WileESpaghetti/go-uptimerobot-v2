package api

type Envelope struct {
// Envelope contains common API response metadata
	Stat  string `json:"stat"`
	Error Error  `json:"error,omitempty"`
	// TODO pagination
}
