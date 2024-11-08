package alert_contact

import (
	"encoding/json"
	"fmt"
	"github.com/go-test/deep"
	"strings"
	"testing"
	"time"
)

func TestType_String(t *testing.T) {
	tests := map[string]struct {
		in   Type
		want string
	}{
		"sms": {
			in:   TypeSMS,
			want: "SMS",
		},
		"email": {
			in:   TypeEmail,
			want: "Email",
		},
		"twitter": {
			in:   TypeTwitter,
			want: "Twitter",
		},
		"webhook": {
			in:   TypeWebhook,
			want: "Webhook",
		},
		"pushbullet": {
			in:   TypePushbullet,
			want: "Pushbullet",
		},
		"zapier": {
			in:   TypeZapier,
			want: "Zapier",
		},
		"prosms": {
			in:   TypeProSMS,
			want: "ProSMS",
		},
		"pushover": {
			in:   TypePushover,
			want: "Pushover",
		},
		"slack": {
			in:   TypeSlack,
			want: "Slack",
		},
		"voicecall": {
			in:   TypeVoiceCall,
			want: "Voice Call",
		},
		"splunk": {
			in:   TypeSplunk,
			want: "Splunk",
		},
		"pagerduty": {
			in:   TypePagerDuty,
			want: "PagerDuty",
		},
		"opsgenie": {
			in:   TypeOpsgenie,
			want: "Opsgenie",
		},
		"msteams": {
			in:   TypeMSTeams,
			want: "MS Teams",
		},
		"googlechat": {
			in:   TypeGoogleChat,
			want: "Google Chat",
		},
		"discord": {
			in:   TypeDiscord,
			want: "Discord",
		},
		"unknown": {
			in:   Type(0),
			want: "Unknown",
		},
		"unknown 2": {
			in:   Type(123),
			want: "Unknown",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			//t.Parallel() // marks each test case as capable of running in parallel with each other
			if got := tt.in.String(); !strings.EqualFold(got, tt.want) {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatus_String(t *testing.T) {
	tests := map[string]struct {
		in   Status
		want string
	}{
		"sms": {
			in:   StatusNotActivated,
			want: "Not Activated",
		},
		"email": {
			in:   StatusPaused,
			want: "Paused",
		},
		"twitter": {
			in:   StatusActive,
			want: "Active",
		},
		"unknown": {
			in:   Status(123),
			want: "Unknown",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			//t.Parallel() // marks each test case as capable of running in parallel with each other
			if got := tt.in.String(); !strings.EqualFold(got, tt.want) {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlertContact_UnmarshalJSON(t *testing.T) {
	// TODO add test to make sure leading zeros are not treated as octal and show up in output
	acJson := []byte(`{
	"id": "12345678",
	"type": 2,
	"friendly_name": "test contact",
	"value": "example@example.com",
	"status": 2,
	"threshold": 137,
	"recurrence": 15
}`)

	want := AlertContact{
		ID:           "12345678",
		Type:         2,
		FriendlyName: "test contact",
		Value:        "example@example.com",
		Status:       2,
		Threshold:    time.Duration(137) * time.Minute,
		Recurrence:   time.Duration(15) * time.Minute,
	}

	got := AlertContact{}
	err := json.Unmarshal(acJson, &got)
	if err != nil {
		t.Errorf("unexpected error unmarshalling json: %v", err)
	}

	if diff := deep.Equal(want, got); diff != nil {
		t.Error(diff)
	}
}

func TestAlertContact_MarshalJSON(t *testing.T) {
	wantJson := []byte(`{
	"id": "12345678",
	"type": 2,
	"friendly_name": "test contact",
	"value": "example@example.com",
	"status": 2,
	"threshold": 137,
	"recurrence": 15
}`)

	ac := AlertContact{
		ID:           "12345678",
		Type:         2,
		FriendlyName: "test contact",
		Value:        "example@example.com",
		Status:       2,
		Threshold:    time.Duration(137) * time.Minute,
		Recurrence:   time.Duration(15) * time.Minute,
	}

	gotJson, err := json.Marshal(ac)
	if err != nil {
		t.Errorf("unexpected error unmarshalling json: %v", err)
	}
	fmt.Printf("%s\n\n", gotJson)

	var want map[string]any
	var got map[string]any

	err = json.Unmarshal(wantJson, &want)
	if err != nil {
		t.Fatalf("unexpected error unmarshalling wanted value for json map: %v", err)
	}

	err = json.Unmarshal(gotJson, &got)
	if err != nil {
		t.Fatalf("unexpected error unmarshalling wanted value for json map: %v", err)
	}

	if diff := deep.Equal(want, got); diff != nil {
		t.Error(diff)
	}
}
