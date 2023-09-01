package api

import "github.com/WileESpaghetti/go-uptimerobot-v2/uptime_robot/models"

type GetAccountDetails struct {
	Envelope
	Account models.Account `json:"account,omitempty"`
}
