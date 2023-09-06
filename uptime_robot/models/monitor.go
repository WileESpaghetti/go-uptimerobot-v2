package models

import (
	"encoding/json"
	"errors"
	"github.com/WileESpaghetti/go-uptimerobot-v2/uptime_robot/breakcircle"
	"github.com/WileESpaghetti/go-uptimerobot-v2/uptime_robot/monitors"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Monitor struct {
	ID              int64                      `schema:"id,omitempty"  json:"id,omitempty"`
	FriendlyName    string                     `schema:"friendly_name" json:"friendly_name"`
	Url             *url.URL                   `schema:"-"             json:"-"`
	Type            monitors.Type              `schema:"type"          json:"type"`
	Status          monitors.Status            `schema:"status"        json:"status"`
	SubType         monitors.SubType           `schema:"sub_type"      json:"sub_type"`
	KeywordType     monitors.KeywordType       `schema:"keyword_type"  json:"keyword_type"`
	KeywordValue    string                     `schema:"keyword_value" json:"keyword_value"`
	HttpUsername    string                     `schema:"http_username" json:"http_username"`
	HttpPassword    string                     `schema:"http_password" json:"http_password"`
	Port            breakcircle.OptionalNumber `schema:"port"          json:"port"` // FIXME might make more sense to move the option numbers to the json version, unless we want to print out "" whenever = 0
	Interval        int64                      `schema:"interval"      json:"interval"`
	CreateDatetime  time.Time                  `schema:"-"             json:"-"` // FIXME not in API docs. need to send email
	KeywordCaseType monitors.KeywordCaseType   `schema:"-"             json:"keyword_case_type"`
	Timeout         int64                      `schema:"-"             json:"timeout"`
}

type unencodableMonitor Monitor

type JsonMonitor struct {
	unencodableMonitor
	Url            string                     `schema:"url"             json:"url"`
	CreateDatetime int64                      `schema:"create_datetime" json:"create_datetime"`
	SubType        breakcircle.OptionalNumber `schema:"sub_type"      json:"sub_type"`
	KeywordType    breakcircle.OptionalNumber `schema:"keyword_type"  json:"keyword_type"`
}

func (jm JsonMonitor) ToMonitor() Monitor {
	parsedUrl, err := url.Parse(jm.Url)
	if err != nil {
		// FIXME handle url eoncoding error. Dashboard does some frontend validation, but need to check if API also validates
	}

	return Monitor{
		ID:           jm.unencodableMonitor.ID,
		FriendlyName: jm.unencodableMonitor.FriendlyName,
		Type:         jm.unencodableMonitor.Type,
		Status:       jm.unencodableMonitor.Status,
		SubType:      jm.unencodableMonitor.SubType,
		KeywordType:  jm.unencodableMonitor.KeywordType,
		KeywordValue: jm.unencodableMonitor.KeywordValue,
		HttpUsername: jm.unencodableMonitor.HttpUsername,
		HttpPassword: jm.unencodableMonitor.HttpPassword,
		Port:         jm.unencodableMonitor.Port,
		Interval:     jm.unencodableMonitor.Interval,

		// This will be inaccurate for some older monitors. I'm guessing that since the `create_datetime` attribute
		// was added to the API at a later date this results in the creation date older pre-existing monitors to
		// default to `0` (1969-12-31 18:00:00 -0600 CST) since the date of the monitor's creation is unknown.
		// The only workaround for this seems to be deleting the monitor and adding it again.
		CreateDatetime: time.Unix(jm.CreateDatetime, 0),
		Url:            parsedUrl,
	}
}

func (m *Monitor) String() string {
	return strconv.FormatInt(m.ID, 10)
}

func (m *Monitor) UnmarshalJSON(data []byte) error {
	var jm JsonMonitor

	if err := json.Unmarshal(data, &jm); err != nil {
		return err
	}

	*m = jm.ToMonitor()

	return nil
}

////////////////////////////////////////////////////////////

type Monitors []Monitor

func (ms Monitors) String() string {
	ids := make(map[int64]int64, len(ms))

	var combined strings.Builder
	for i, monitor := range ms {
		if _, ok := ids[monitor.ID]; ok {
			continue
		}

		if i > 0 {
			combined.WriteString("-")
		}

		ids[monitor.ID] = monitor.ID
		id := strconv.FormatInt(monitor.ID, 10)
		combined.WriteString(id)
	}

	return combined.String()
}

func (ms Monitors) UnmarshalText(text []byte) error {
	textIDs := strings.Split(string(text), "-")

	for _, sID := range textIDs {
		id, err := strconv.ParseInt(sID, 10, 64)
		if err != nil {
			return errors.New("monitor ID must be an integer")
		}

		ms = append(ms, Monitor{ID: id})
	}

	return nil
}

// Set is used to create a list of Monitor from a dash-separated list of ID
func (ms Monitors) Set(s string) error {
	return ms.UnmarshalText([]byte(s))
}
