package models

import (
	"encoding/json"
	"errors"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Monitors []Monitor

// FIXME CreateDatetime not in API docs. need to send email
type Monitor struct {
	Id             int                `schema:"id,omitempty"  json:"id,omitempty"`
	FriendlyName   string             `schema:"friendly_name" json:"friendly_name"`
	Url            *url.URL           `schema:"-"             json:"-"`
	Type           MonitorType        `schema:"type"          json:"type"`
	Status         MonitorStatus      `schema:"status"        json:"status"`
	SubType        MonitorSubType     `schema:"sub_type"      json:"sub_type"`
	KeywordType    MonitorKeywordType `schema:"keyword_type"  json:"keyword_type"`
	KeywordValue   string             `schema:"keyword_value" json:"keyword_value"`
	HttpUsername   string             `schema:"http_username" json:"http_username"`
	HttpPassword   string             `schema:"http_password" json:"http_password"`
	Port           OptionalNumber     `schema:"port"          json:"port"`
	Interval       int64              `schema:"interval"      json:"interval"`
	CreateDatetime time.Time          `schema:"-"             json:"-"`
}

type unencodableMonitor Monitor

type JsonMonitor struct {
	unencodableMonitor
	Url            string `schema:"url"             json:"url"`
	CreateDatetime int64  `schema:"create_datetime" json:"create_datetime"`
}

func (m *Monitors) String() string {
	// TODO have a type for structs that reduce to `-` separated data
	var ids []string

	for _, monitor := range *m {
		id := strconv.Itoa(monitor.Id)
		ids = append(ids, id)
	}

	combined := strings.Join(ids, "-")

	return combined
}

// TODO need to be for []Monitor and parse #-#-# format instead
// needed to be used as flag value
func (m *Monitors) Set(s string) error {
	ids := strings.Split(s, "-")

	for _, idStr := range ids {
		id, err := strconv.ParseInt(idStr, 10, 0) // TODO extract ID validation for reuse
		if err != nil {
			return errors.New("invalid ID format")
		}

		*m = append(*m, Monitor{Id: int(id)})
	}

	return nil
}

func (m *Monitor) String() string {
	return strconv.Itoa(m.Id)
}

func (jm JsonMonitor) ToMonitor() Monitor {
	parsedUrl, err := url.Parse(jm.Url)
	if err != nil {
		// FIXME handle url eoncoding error. Dashboard does some frontend validation, but need to check if API also validates
	}

	return Monitor{
		Id:           jm.unencodableMonitor.Id,
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
		Url:            parsedUrl}
}

func (m *Monitor) UnmarshalJSON(data []byte) error {
	var jm JsonMonitor

	if err := json.Unmarshal(data, &jm); err != nil {
		return err
	}

	*m = jm.ToMonitor()

	return nil
}
