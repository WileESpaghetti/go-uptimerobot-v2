package api

import (
	"github.com/WileESpaghetti/go-uptimerobot-v2/uptime_robot/alert_contacts"
)

type GetAlertContacts struct {
	Envelope
	AlertContacts []alert_contacts.AlertContact `json:"alert_contacts,omitempty"`
}
