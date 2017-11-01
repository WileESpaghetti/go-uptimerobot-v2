package uptimerobot

import (
	"fmt"
	"io/ioutil"
)

type Monitor struct {
	FriendlyName string `schema:"friendly_name"`
	Url          string `schema:"url"`
	Type         int    `schema:"type"`
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

	body, err := ioutil.ReadAll(r.Body)
	fmt.Println(err)
	fmt.Printf("%s", body)

	return m, nil
}
