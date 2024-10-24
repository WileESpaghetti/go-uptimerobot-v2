package api

import (
	"github.com/WileESpaghetti/go-uptimerobot-v2/uptime_robot/alert_contacts"
)

type GetAlertContacts struct {
	Envelope
	AlertContacts []alert_contacts.AlertContact `json:"alert_contacts,omitempty"`
}

type GetAlertContactsRequest struct {
	AlertContacts alert_contacts.AlertContacts `form:"alert_contacts,omitempty"`
}

// AlertContactOptions configures a GetAlertContactsRequest
type AlertContactOptions func(*GetAlertContactsRequest)

// WithAlertContacts allows you to limit the data returned to specific alert contacts
func WithAlertContacts(contacts alert_contacts.AlertContacts) AlertContactOptions {
	return func(options *GetAlertContactsRequest) {
		options.AlertContacts = contacts
	}
}
