package api

import "fmt"

const (
	ErrorParameterMissing = "missing_parameter"
)

type Error struct {
	Type string `json:"type,omitempty"`
	ParameterName string `json:"parameter_name,omitempty"`
}

func (e Error) Error() string {
	if e.Type == ErrorParameterMissing {
		return e.ParameterMissingError()
	}

	return e.ParameterName
}

func (e Error) ParameterMissingError() string {
	return fmt.Sprintf("No `%s` parameter found in the API request", e.ParameterName)
}
