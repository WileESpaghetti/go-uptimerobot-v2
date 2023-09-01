package models

import "fmt"

type Account struct {
	DownMonitors    int64  `json:"down_monitors"    xml:"down_monitors,attr"`
	Email           string `json:"email"            xml:"email,attr"`
	MonitorInterval int64  `json:"monitor_interval" xml:"monitor_interval,attr"`
	MonitorLimit    int64  `json:"monitor_limit"    xml:"monitor_limit,attr"`
	PausedMonitors  int64  `json:"paused_monitors"  xml:"paused_monitors,attr"`
	UpMonitors      int64  `json:"up_monitors"      xml:"up_monitors,attr"`
}

func (a Account) String() string {
	accountFormat := "Account Details:" +
		"\n Email:            %s" +
		"\n Monitor Limit:    %d" +
		"\n Monitor Interval: %d minute(s)" +
		"\n\nMonitor Details:" +
		"\n Up:     %d" +
		"\n Down:   %d" +
		"\n Paused: %d\n"

	return fmt.Sprintf(accountFormat,
		a.Email,
		a.MonitorLimit,
		a.MonitorInterval,
		a.UpMonitors,
		a.DownMonitors,
		a.PausedMonitors)
}
