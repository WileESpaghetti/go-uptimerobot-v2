package models

const (
	MonitorTypeHttp    = 1
	MonitorTypeKeyword = 2
	MonitorTypePing    = 3
	MonitorTypePort    = 4
)

type MonitorType int

func (mt MonitorType) String() string {
	if mt == MonitorTypeHttp {
		return "HTTP"
	} else if mt == MonitorTypeKeyword {
		return "Keyword"
	} else if mt == MonitorTypePing {
		return "Ping"
	} else if mt == MonitorTypePort {
		return "Port"
	} else {
		return "Unknown"
	}
}
