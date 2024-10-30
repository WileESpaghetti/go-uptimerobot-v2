package alert_contact

import "github.com/WileESpaghetti/go-uptimerobot-v2/uptime_robot/api"

type GetAlertContacts struct {
	api.Envelope
	AlertContacts []AlertContact `json:"alert_contacts,omitempty"`
}

type GetAlertContactsRequest struct {
	AlertContacts AlertContacts `form:"alert_contacts,omitempty"`
}

// AlertContactOptions configures a GetAlertContactsRequest
type AlertContactOptions func(*GetAlertContactsRequest)

// WithOptions allows setting all options for [getAlertContacts] request at once.
//
// This is a good option if you are setting many query parameters and do not
// want to pass several AlertContactOptions.
//
// [getAlertContacts]: https://uptimerobot.com/api/
func WithOptions(getAlertContacts *GetAlertContactsRequest) AlertContactOptions {
	return func(options *GetAlertContactsRequest) {
		*options = *getAlertContacts
	}
}

// WithAlertContacts allows you to limit the data returned to specific alert contacts
func WithAlertContacts(contacts AlertContacts) AlertContactOptions {
	return func(options *GetAlertContactsRequest) {
		options.AlertContacts = contacts
	}
}
