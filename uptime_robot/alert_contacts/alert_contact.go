package alert_contacts

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// Type is an Enum
type Type int64

// Type supported by the API
const (
	TypeSMS        Type = 1
	TypeEmail      Type = 2
	TypeTwitter    Type = 3
	TypeWebhook    Type = 5
	TypePushbullet Type = 6
	TypeZapier     Type = 7
	TypeProSMS     Type = 8
	TypePushover   Type = 9
	TypeSlack      Type = 11
	TypeVoiceCall  Type = 14
	TypeSplunk     Type = 15
	TypePagerDuty  Type = 16
	TypeOpsgenie   Type = 17
	TypeMSTeams    Type = 20
	TypeGoogleChat Type = 21
	TypeDiscord    Type = 23
)

func (t Type) String() string {
	switch t {
	case TypeSMS:
		return "SMS"
	case TypeEmail:
		return "Email"
	case TypeTwitter:
		return "Twitter"
	case TypeWebhook:
		return "Webhook"
	case TypePushbullet:
		return "Pushbullet"
	case TypeZapier:
		return "Zapier"
	case TypeProSMS:
		return "ProSMS"
	case TypePushover:
		return "Pushover"
	case TypeSlack:
		return "Slack"
	case TypeVoiceCall:
		return "VoiceCall"
	case TypeSplunk:
		return "Splunk"
	case TypePagerDuty:
		return "PagerDuty"
	case TypeOpsgenie:
		return "Opsgenie"
	case TypeMSTeams:
		return "MSTeams"
	case TypeGoogleChat:
		return "GoogleChat"
	case TypeDiscord:
		return "Discord"
	default:
		return "Unknown"
	}
}

type Status int64

// Status of the Alert Contact
const (
	StatusNotActivated Status = 0
	StatusPaused       Status = 1
	StatusActive       Status = 2
)

func (s Status) String() string {
	switch s {
	case StatusNotActivated:
		return "Not Activated"
	case StatusPaused:
		return "Paused"
	case StatusActive:
		return "Active"
	default:
		return "Unknown"
	}
}

type AlertContact struct {
	// ID can have leading zeros so it must be a string
	ID           string `json:"id"`
	Type         Type   `json:"type"`
	FriendlyName string `json:"friendly_name"`
	Value        string `json:"value"`
	Status       Status `json:"status"`

	// minutes in the API
	Threshold time.Duration `json:"threshold"`

	// minutes in the API
	Recurrence time.Duration `json:"recurrence"`
}

// unencodableMonitor is used to break encoding loops for jsonMonitor
type unencodableAlertContact AlertContact

// jsonResponseTimeEntry is an intermediate form that we will use to convert between more strict time formats
type jsonAlertContact struct {
	unencodableAlertContact

	// minutes
	Threshold int64 `json:"threshold"`

	// minutes
	Recurrence int64 `json:"recurrence"`
}

func (jac *jsonAlertContact) AlertContact() AlertContact {
	return AlertContact{
		ID:           jac.ID,
		Type:         jac.Type,
		FriendlyName: jac.FriendlyName,
		Value:        jac.Value,
		Status:       jac.Status,
		Threshold:    time.Duration(jac.Threshold) * time.Minute,
		Recurrence:   time.Duration(jac.Recurrence) * time.Minute,
	}
}

// UnmarshalJSON is needed to convert between an untyped API response to use more accurate [time] types
func (ac *AlertContact) UnmarshalJSON(b []byte) error {
	var raw jsonAlertContact

	err := json.Unmarshal(b, &raw)
	if err != nil {
		return err
	}

	*ac = raw.AlertContact()

	return nil
}

// MarshalJSON will convert the time formats between what the API uses and what the API uses
func (ac AlertContact) MarshalJSON() ([]byte, error) {
	a := int64(ac.Threshold / time.Minute)
	fmt.Println(a)
	uac := unencodableAlertContact(ac)
	return json.Marshal(jsonAlertContact{
		unencodableAlertContact: uac,
		Threshold:               int64(ac.Threshold / time.Minute),
		Recurrence:              int64(ac.Recurrence / time.Minute),
	})
}

////////////////////

type AlertContacts []AlertContact

func (acs AlertContacts) String() string {
	ids := make(map[string]string, len(acs))

	var combined strings.Builder
	for i, contact := range acs {
		if _, ok := ids[contact.ID]; ok {
			continue
		}

		if i > 0 {
			combined.WriteString("-")
		}

		ids[contact.ID] = contact.ID
		combined.WriteString(contact.ID)
	}

	return combined.String()
}

func (acs *AlertContacts) MarshalText() ([]byte, error) {
	// FIXME not sure which is faster this or strings.builder
	//return []byte(acs.String()), nil
	var ids []string

	for _, ac := range *acs {
		ids = append(ids, ac.ID)
	}

	combined := strings.Join(ids, "-")

	return []byte(combined), nil
}

func (acs *AlertContacts) UnmarshalText(text []byte) error {
	textIDs := strings.Split(string(text), "-")

	for _, id := range textIDs {
		*acs = append(*acs, AlertContact{ID: id})
	}

	return nil
}

// Set is used to create a list of Monitor from a dash-separated list of ID
func (acs *AlertContacts) Set(s string) error {
	return acs.UnmarshalText([]byte(s))
}
