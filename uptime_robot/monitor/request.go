package monitor

import "github.com/WileESpaghetti/go-uptimerobot-v2/uptime_robot/api"

// GetMonitors provides monitor configuration and status information.
type GetMonitors struct {
	api.Envelope
	Monitors []Monitor `json:"monitors,omitempty"`
}

// GetMonitorsRequest provides a way to configure Monitor data
type GetMonitorsRequest struct {
	Monitors              Monitors `form:"monitors,omitempty"`
	Types                 Types    `form:"types,omitempty"`
	Statuses              Statuses `form:"statuses,omitempty"`
	UptimeRatios          []int64  `form:"custom_uptime_ratios,omitempty"` // days
	HasAllTimeUptimeRatio bool     `form:"all_time_uptime_ratio,omitempty"`
	HasLogs               bool     `form:"logs,omitempty"`
	LogTypes              LogTypes `form:"log_types,omitempty"`
	LogsLimit             int64    `form:"logs_limit,omitempty"`
	HasResponseTimes      bool     `form:"response_times,omitempty"`
	HasCustomHTTPStatuses bool     `form:"custom_http_statuses,omitempty"`

	// ResponseTimesLimit only takes effect when not using response_times_start_date and response_times_end_date
	// If empty, last 24 hours of logs are returned (
	ResponseTimesLimit int64 `form:"response_times_limit,omitempty"`

	// ResponseTimesAverage averages the response times in intervals using the specified number of minutes.
	// 0 is the default and is unaveraged. The dashboard uses 30 minutes
	ResponseTimesAverage int64 `form:"response_times_limit,omitempty"`
	HasAlertContacts     bool  `form:"alert_contacts,omitempty"`
	HasAuthType          bool  `form:"auth_type,omitempty"`
}

/*
type getMonitorsOptions struct {
	all_time_uptime_ratio     bool `` //   "all_time_uptime_ratio": "97.890"
	all_time_uptime_durations bool
	response_times            bool
	alert_contacts            bool
	mwindows                  bool
	ssl                       bool
	custom_http_statuses      bool
	timezone                  bool
	//custom_down_durations // only accepted values are 0 or 1, but doesn't appear to do anything or give an invalid parameter error. need to file docs correction
}
*/

// MonitorOptions configures a GetMonitorsRequest
type MonitorOptions func(*GetMonitorsRequest)

// WithOptions allows setting all options for [getMonitors] request at once.
//
// This is a good option if you are setting many query parameters and do not
// want to pass several MonitorOptions.
//
// [getMonitors]: https://uptimerobot.com/api/
func WithOptions(getMonitorsRequest *GetMonitorsRequest) MonitorOptions {
	return func(options *GetMonitorsRequest) {
		*options = *getMonitorsRequest
	}
}

// WithMonitors allows you to limit the data returned to specific monitors
func WithMonitors(monitors Monitors) MonitorOptions {
	return func(options *GetMonitorsRequest) {
		options.Monitors = monitors
	}
}

// WithTypes allows you to limit returned monitors to specific [Type]
//
// Supported monitor [Type]s
//   - [TypeHttp]
//   - [TypeKeyword]
//   - [TypePing]
//   - [TypePort]
func WithTypes(types Types) MonitorOptions {
	return func(options *GetMonitorsRequest) {
		options.Types = types
	}
}

func WithStatuses(statuses Statuses) MonitorOptions {
	return func(options *GetMonitorsRequest) {
		options.Statuses = statuses
	}
}

func WithUptimeRatios(uptimeRatios []int64) MonitorOptions {
	return func(options *GetMonitorsRequest) {
		options.UptimeRatios = uptimeRatios
	}
}

func WithAllTimeUptimeRatio(shouldInclude bool) MonitorOptions {
	return func(options *GetMonitorsRequest) {
		options.HasAllTimeUptimeRatio = shouldInclude
	}
}

func WithLogs(shouldInclude bool) MonitorOptions {
	return func(options *GetMonitorsRequest) {
		options.HasLogs = shouldInclude
	}
}

func WithLogsLimit(limit int64) MonitorOptions {
	return func(options *GetMonitorsRequest) {
		options.LogsLimit = limit
	}
}

func WithResponseTimes(shouldInclude bool) MonitorOptions {
	return func(options *GetMonitorsRequest) {
		options.HasResponseTimes = shouldInclude
	}
}

func WithResponseTimesLimit(limit int64) MonitorOptions {
	return func(options *GetMonitorsRequest) {
		options.ResponseTimesLimit = limit
	}
}

func WithResponseTimesAverage(minutesPerInterval int64) MonitorOptions {
	return func(options *GetMonitorsRequest) {
		options.ResponseTimesAverage = minutesPerInterval
	}
}

func WithLogTypes(types LogTypes) MonitorOptions {
	return func(options *GetMonitorsRequest) {
		options.LogTypes = types
	}
}

func WithAlertContacts(shouldInclude bool) MonitorOptions {
	return func(options *GetMonitorsRequest) {
		options.HasAlertContacts = shouldInclude
	}
}

func WithAuthType(shouldInclude bool) MonitorOptions {
	return func(options *GetMonitorsRequest) {
		options.HasAuthType = shouldInclude
	}
}

func WithCustomHTTPStatuses(shouldInclude bool) MonitorOptions {
	return func(options *GetMonitorsRequest) {
		options.HasCustomHTTPStatuses = shouldInclude
	}
}
