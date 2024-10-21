package alert_contacts

import (
	"encoding/json"
	"fmt"
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

type Status int64

// Status of the Alert Contact
const (
	StatusNotActivated Status = 0
	StatusPaused       Status = 1
	StatusActive       Status = 2
)

type AlertContact struct {
	ID           int64  `json:"id"`
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
