package api

type Error struct {
	Type string `json:"type,omitempty"`
}

func (e Error) Error() string {
	return e.Type
}
