package uptimerobot

import (
	"encoding/json"
	"fmt"
)

type Monitor struct {
	Id           json.Number `schema:"id,omitempty",json:"id,omitempty"`
	FriendlyName string      `schema:"friendly_name",json:"friendly_name"`
	Url          string      `schema:"url",json:"url"`
	Type         int         `schema:"type",json:"type"`
	Status       int         `schema:"status",json:"status"`
}

type SingleMonitorResponse struct {
	ApiResponse
	Monitor Monitor `json:"monitor"`
}

type NewMonitorResponse struct {
	ApiResponse
	Monitor Monitor `json:"monitor"`
}

func (c *Client) NewMonitor(m *Monitor) (*Monitor, error) {
	newMonitorRequest, err := c.NewRequest("newMonitor", m)
	if err != nil {
		return nil, err
	}

	r, err := c.HttpClient.Do(newMonitorRequest)
	if err != nil {
		return nil, err
	}

	monitor := &NewMonitorResponse{}
	err = json.NewDecoder(r.Body).Decode(monitor)
	if err != nil {
		return nil, err
	}

	if monitor.Error != nil {
		return nil, fmt.Errorf("%s: %s", monitor.Error.Type, monitor.Error.Message)
	}

	m.Id = monitor.Monitor.Id

	return m, nil
}

func (c *Client) DeleteMonitor(m *Monitor) error {
	deleteMonitorRequest, err := c.NewRequest("deleteMonitor", m)
	if err != nil {
		return err
	}

	r, err := c.HttpClient.Do(deleteMonitorRequest)
	if err != nil {
		return err
	}

	monitor := &SingleMonitorResponse{}
	err = json.NewDecoder(r.Body).Decode(monitor)
	if err != nil {
		return err
	}

	if monitor.Error != nil {
		return fmt.Errorf("%s: %s", monitor.Error.Type, monitor.Error.Message)
	}

	return nil
}