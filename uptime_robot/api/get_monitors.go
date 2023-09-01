package api

import (
	"github.com/WileESpaghetti/go-uptimerobot-v2/uptime_robot/models"
)

type GetMonitors struct {
	Envelope
	Monitors []models.Monitor `json:"monitors,omitempty"`
}

type GetMonitorsQuery struct {
	Monitors []models.Monitor `schema:"monitors,omitempty"`
}
