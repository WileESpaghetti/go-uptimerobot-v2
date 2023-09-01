package breakcircle

// FIXME need to actually organize in a way that doesn't have circular imports

import (
	"encoding/json"
	"strconv"
)

// OptionalNumber handles cases where the API uses an empty string instead of a null, or undefined
// is used for an optional whole number value.
//
// int and json.Number types fail to unmarshal an empty string, and strings fail to unmarshal a number literal that is
// unquoted. json.Number also supports floating point numbers which isn't useful for what these attributes represent.
type OptionalNumber int64

func (n *OptionalNumber) UnmarshalJSON(data []byte) error {
	// handle empty string
	if len(data) == 2 && data[0] == '"' && data[len(data)-1] == '"' {
		*n = OptionalNumber(0)
		return nil
	}

	var numVal json.Number
	err := json.Unmarshal(data, &numVal)
	if err != nil {
		*n = OptionalNumber(0)
		return err
	}

	if len(numVal) == 0 {
		*n = OptionalNumber(0)
		return nil
	}

	i, err := numVal.Int64()
	if err != nil {
		*n = OptionalNumber(0)
		return err // probably a decimal
	}

	*n = OptionalNumber(i)
	return nil
}

func (n *OptionalNumber) String() string {
	if *n == 0 {
		return ""
	}

	return strconv.FormatInt(int64(*n), 10)
}
