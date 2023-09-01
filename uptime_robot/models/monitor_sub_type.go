package models

import (
	"encoding/json"
	"strconv"
)

const (
	MonitorSubTypeHttp       = 1
	MonitorSubTypeHttps      = 2
	MonitorSubTypeFtp        = 3
	MonitorSubTypeSmtp       = 4
	MonitorSubTypePop3       = 5
	MonitorSubTypeImap       = 6
	MonitorSubTypeCustomPort = 99
)

type MonitorSubType string

func (mst *MonitorSubType) UnmarshalJSON(data []byte) error {
	// handle empty string
	if len(data) == 2 && data[0] == '"' && data[len(data)-1] == '"' {
		var str string
		strErr := json.Unmarshal(data, &str)
		if strErr != nil {
			return strErr
		}

		*mst = MonitorSubType(str)
		return nil
	}

	var numVal json.Number
	err := json.Unmarshal(data, &numVal)
	if err != nil {
		return err
	}

	mstVal := MonitorSubType(numVal.String())
	*mst = mstVal
	return nil
}

func (mst MonitorSubType) String() string {
	mstVal, err := strconv.ParseInt(string(mst), 10, 64)

	if err != nil {
		return ""
	} else if mstVal == MonitorSubTypeHttp {
		return "HTTP"
	} else if mstVal == MonitorSubTypeHttps {
		return "HTTPS"
	} else if mstVal == MonitorSubTypeFtp {
		return "FTP"
	} else if mstVal == MonitorSubTypeSmtp {
		return "SMTP"
	} else if mstVal == MonitorSubTypePop3 {
		return "POP3"
	} else if mstVal == MonitorSubTypeImap {
		return "IMAP"
	} else if mstVal == MonitorSubTypeCustomPort {
		return "Custom Port"
	} else {
		return "Unknown"
	}
}
