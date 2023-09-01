package models

import (
	"encoding/json"
	"strconv"
)

const (
	MonitorKeywordTypeExists    = 1
	MonitorKeywordTypeNotExists = 2
)

type MonitorKeywordType string

func (mkt *MonitorKeywordType) UnmarshalJSON(data []byte) error {
	// handle empty string
	if len(data) == 2 && data[0] == '"' && data[len(data)-1] == '"' {
		var str string
		strErr := json.Unmarshal(data, &str)
		if strErr != nil {
			return strErr
		}

		*mkt = MonitorKeywordType(str)
		return nil
	}

	var numVal json.Number
	err := json.Unmarshal(data, &numVal)
	if err != nil {
		return err
	}

	mstVal := MonitorKeywordType(numVal.String())
	*mkt = mstVal
	return nil
}

func (mkt MonitorKeywordType) String() string {
	mstVal, err := strconv.ParseInt(string(mkt), 10, 64)

	if err != nil {
		return ""
	} else if mstVal == MonitorKeywordTypeExists {
		return "Exists"
	} else if mstVal == MonitorKeywordTypeNotExists {
		return "Not Exists"
	} else {
		return "Unknown"
	}
}
