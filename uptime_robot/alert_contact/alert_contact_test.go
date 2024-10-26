package alert_contact

import (
	"encoding/json"
	"fmt"
	"github.com/go-test/deep"
	"testing"
	"time"
)

func TestAlertContact_UnmarshalJSON(t *testing.T) {
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
