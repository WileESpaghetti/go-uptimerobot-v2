package api

import (
	"fmt"
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

type GetMonitorsQuery struct {
	Monitors []models.Monitor `schema:"monitors,omitempty"`
}

func (m GetMonitorsQuery) RegisterEncoders(e *schema.Encoder) {
	e.RegisterEncoder(m.Monitors, MonitorsSchemaEncoder)
}

func MonitorsSchemaEncoder(v reflect.Value) string {
	var ids []string

	m := v.Interface().([]models.Monitor)

	for _, monitor := range m {
		id := strconv.Itoa(monitor.Id)
		ids = append(ids, id)
	}

	combined := strings.Join(ids, "-")
	fmt.Println(combined)

	return combined
}
