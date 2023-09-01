package api

type Envelope struct {
	Stat string `json:"stat"`
	Error Error `json:"error,omitempty"`
}
