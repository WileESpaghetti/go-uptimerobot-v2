package account

import (
	"github.com/WileESpaghetti/go-uptimerobot-v2/uptime_robot/api"
)

// GetAccountDetails is details about your account and API response metadata
type GetAccountDetails struct {
	api.Envelope
	Account Account `json:"account,omitempty"`
}
