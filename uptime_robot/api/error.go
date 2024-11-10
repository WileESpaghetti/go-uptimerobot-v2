package api

import "fmt"

// [api.Error.Type] values
const (
	ErrorTypeParameterMissing = "missing_parameter"
	ErrorTypeNotAuthorized    = "not_authorized"
	ErrorTypeInvalidParameter = "invalid_parameter"
	ErrorTypeBadRequest       = "bad_request"
	ErrorTypeInternalError    = "internal_error"
)

// Error messages that are more user-friendly (wordy) than what the UptimeRobot provides
const (
	ErrParameterMissing = "no `%s` parameter found in the API request"
	ErrNotAuthorized    = "API key is not authorized to make this request"
)

// Error provides details for any API request failures.
// An API request failure is indicated when [api.Envelope.Stat] matches [api.StatFail]
type Error struct {
	Type          string `json:"type,omitempty"`
	ParameterName string `json:"parameter_name,omitempty"`
	Message       string `json:"message,omitempty"`
}

// Error enables using API errors as Go errors
func (e Error) Error() string {
	switch e.Type {
	case ErrorTypeParameterMissing:
		return fmt.Sprintf(ErrParameterMissing, e.ParameterName)
	case ErrorTypeNotAuthorized:
		return ErrNotAuthorized
	case ErrorTypeInvalidParameter:
		return e.Message
	case ErrorTypeBadRequest:
		return e.Message
	case ErrorTypeInternalError:
		return e.Message
	default:
		if len(e.Message) > 0 {
			return e.Message
		}
		return e.Type
	}
}
