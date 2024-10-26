package alert_contact

import "github.com/WileESpaghetti/go-uptimerobot-v2/uptime_robot/api"

type GetAlertContacts struct {
	api.Envelope
	AlertContacts []AlertContact `json:"alert_contacts,omitempty"`
}

type GetAlertContactsRequest struct {
	AlertContacts AlertContacts `form:"alert_contact,omitempty"`
}

// AlertContactOptions configures a GetAlertContactsRequest
type AlertContactOptions func(*GetAlertContactsRequest)

// WithAlertContacts allows you to limit the data returned to specific alert contacts
func WithAlertContacts(contacts AlertContacts) AlertContactOptions {
	return func(options *GetAlertContactsRequest) {
		options.AlertContacts = contacts
	}
}
