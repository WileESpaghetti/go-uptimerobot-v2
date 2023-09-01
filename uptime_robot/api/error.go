package api

import "fmt"

// Error.Type values
const (
	ErrorTypeParameterMissing = "missing_parameter"
	ErrorTypeNotAuthorized    = "not_authorized"
)

// Error messages
const (
	ErrParameterMissing = "no `%s` parameter found in the API request"
	ErrNotAuthorized    = "API key is not authorized to make this request"
)

type Error struct {
	Type          string `json:"type,omitempty"`
	ParameterName string `json:"parameter_name,omitempty"`
}

func (e Error) Error() string {
	switch e.Type {
	case ErrorTypeParameterMissing:
		return fmt.Sprintf(ErrParameterMissing, e.ParameterName)
	case ErrorTypeNotAuthorized:
		return ErrNotAuthorized
	default:
		return e.Type
	}
}
