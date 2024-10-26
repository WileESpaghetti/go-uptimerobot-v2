package number

import (
	"encoding/json"
	"strconv"
)

// An Optional number provides a way to decode loosely typed integer JSON values
//
// The UptimeRobot API sometimes uses an empty string instead of a null, or undefined when a value only optionally needs
// to be included in the API response. Empty strings cause issues with decoding because using an int type or
// json.Number is unable to unmarshal an empty string, and strings are unable to unmarshal cases where a number literal
// is unquoted. json.Number also supports floating point numbers which isn't useful for what these attributes represent.
type Optional int64

// UnmarshalJSON implements the [encoding/json.Unmarshaler] interface.
func (n *Optional) UnmarshalJSON(data []byte) error {
	// handle empty string
	if len(data) == 2 && data[0] == '"' && data[len(data)-1] == '"' {
		*n = Optional(0)
		return nil
	}

	var numVal json.Number
	err := json.Unmarshal(data, &numVal)
	if err != nil {
		*n = Optional(0)
		return err
	}

	if len(numVal) == 0 {
		*n = Optional(0)
		return nil
	}

	i, err := numVal.Int64()
	if err != nil {
		*n = Optional(0)
		return err // probably a decimal
	}

	*n = Optional(i)
	return nil
}

func (n Optional) String() string {
	// FIXME is `0` an allowed value?
	if n == 0 {
		return ""
	}

	return strconv.FormatInt(int64(n), 10)
}
