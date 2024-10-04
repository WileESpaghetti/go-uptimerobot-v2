package api

import "github.com/WileESpaghetti/go-uptimerobot-v2/uptime_robot/models"

// GetAccountDetails is details about your account and API response metadata
type GetAccountDetails struct {
	Envelope
	Account models.Account `json:"account,omitempty"`
}
