package api

import "fmt"

const (
	ErrorParameterMissing = "missing_parameter"
	ErrorNotAuthorized = "not_authorized"
)

type Error struct {
	Type string `json:"type,omitempty"`
	ParameterName string `json:"parameter_name,omitempty"`
}

func (e Error) Error() string {
	if e.Type == ErrorParameterMissing {
		return e.ParameterMissingError()
	} else if e.Type == ErrorNotAuthorized {
		return e.NotAuthorizedError()
	}

	return e.Type
}

func (e Error) ParameterMissingError() string {
	return fmt.Sprintf("No `%s` parameter found in the API request", e.ParameterName)
}

func (e Error) NotAuthorizedError() string {
	return "API key is not authorized to make this request"
}
