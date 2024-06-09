package api

import (
	"github.com/WileESpaghetti/go-uptimerobot-v2/uptime_robot/monitors"

	"github.com/WileESpaghetti/go-uptimerobot-v2/uptime_robot/models"
)

type GetMonitors struct {
	Envelope
	Monitors []models.Monitor `json:"monitors,omitempty"`
}

type GetMonitorsRequest struct {
	Monitors models.Monitors `schema:"monitors,omitempty"`
	Monitors models.Monitors `form:"monitors,omitempty"`
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
