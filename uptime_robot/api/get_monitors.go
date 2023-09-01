package api

import (
	"github.com/WileESpaghetti/go-uptimerobot-v2/uptime_robot/models"
	"github.com/gorilla/schema"
	"reflect"
	"strconv"
	"strings"
)

type GetMonitors struct {
	Envelope
	Monitors []models.Monitor `json:"monitors,omitempty"`
}

type GetMonitorsRequest struct {
	Monitors []models.Monitor `schema:"monitors,omitempty"`
}

type GetMonitorsOptions struct {
	all_time_uptime_ratio     bool //   "all_time_uptime_ratio": "97.890"
	all_time_uptime_durations bool
	logs                      bool
	response_times            bool
	alert_contacts            bool
	mwindows                  bool
	ssl                       bool
	custom_http_statuses      bool
	timezone                  bool
}

func (m GetMonitorsRequest) RegisterEncoders(e *schema.Encoder) {
	e.RegisterEncoder(m.Monitors, MonitorsSchemaEncoder)
}

// TODO dedup monitors
// TODO may be able to replace this with unmarshalText
func MonitorsSchemaEncoder(v reflect.Value) string {
	var ids []string

	m := v.Interface().([]models.Monitor)

	for _, monitor := range m {
		id := strconv.Itoa(monitor.Id)
		ids = append(ids, id)
	}

	combined := strings.Join(ids, "-")

	return combined
}
