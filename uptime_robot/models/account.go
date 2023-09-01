package models

type Account struct {
	DownMonitors    int64  `json:"down_monitors"    xml:"down_monitors,attr"`
	Email           string `json:"email"            xml:"email,attr"`
	MonitorInterval int64  `json:"monitor_interval" xml:"monitor_interval,attr"`
	MonitorLimit    int64  `json:"monitor_limit"    xml:"monitor_limit,attr"`
	PausedMonitors  int64  `json:"paused_monitors"  xml:"paused_monitors,attr"`
	UpMonitors      int64  `json:"up_monitors"      xml:"up_monitors,attr"`
}

