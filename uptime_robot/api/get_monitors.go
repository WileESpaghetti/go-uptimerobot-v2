package api

import (
	"github.com/WileESpaghetti/go-uptimerobot-v2/uptime_robot/models"
	"github.com/WileESpaghetti/go-uptimerobot-v2/uptime_robot/monitors"
)

type GetMonitors struct {
	Envelope
	Monitors []models.Monitor `json:"monitors,omitempty"`
}

type GetMonitorsRequest struct {
	Monitors models.Monitors `form:"monitors,omitempty"`
	Types    monitors.Types  `form:"types,omitempty"`
}

type GetMonitorsOptions struct {
	all_time_uptime_ratio     bool `` //   "all_time_uptime_ratio": "97.890"
	all_time_uptime_durations bool
	logs                      bool
	response_times            bool
	alert_contacts            bool
	mwindows                  bool
	ssl                       bool
	custom_http_statuses      bool
	timezone                  bool
}

type MonitorOptions func(*GetMonitorsRequest)

func WithOptions(getMonitorsRequest *GetMonitorsRequest) MonitorOptions {
	return func(options *GetMonitorsRequest) {
		*options = *getMonitorsRequest
	}
}

func WithMonitors(monitors models.Monitors) MonitorOptions {
	return func(options *GetMonitorsRequest) {
		options.Monitors = monitors
	}
}

func WithTypes(types monitors.Types) MonitorOptions {
	return func(options *GetMonitorsRequest) {
		options.Types = types
	}
}
