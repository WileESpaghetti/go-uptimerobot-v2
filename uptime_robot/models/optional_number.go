package models

import "encoding/json"

type OptionalNumber string

func (t *OptionalNumber) UnmarshalJSON(data []byte) error {
	// handle empty string
	if len(data) == 2 && data[0] == '"' && data[len(data)-1] == '"' {
		var str string
		strErr := json.Unmarshal(data, &str)
		if strErr != nil {
			return strErr
		}

		*t = OptionalNumber(str)
		return nil
	}

	var numVal json.Number
	err := json.Unmarshal(data, &numVal)
	if err != nil {
		return err
	}

	mstVal := OptionalNumber(numVal.String())
	*t = mstVal
	return nil
}
