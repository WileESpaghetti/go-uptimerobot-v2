package public_status_page

import (
	"encoding/json"
	"net/url"
	"strconv"
	"strings"
)

// FIXME need to polish up how monitor IDs are stored and how we retrieve monitors

type PublicStatusPage struct {
	ID           int64    `json:"id"`
	FriendlyName string   `json:"friendly_name"`
	monitorIDs   []int64  `json:"monitors"`
	Sort         Sort     `json:"sort"`
	Status       Status   `json:"status"`
	StandardURL  *url.URL `json:"standard_url"`
	CustomURL    *url.URL `json:"custom_url"`
}

//func (psp *PublicStatusPage) Monitors(c *uptime_robot.Client) ([]monitor.Monitor, error) {
//	if c == nil {
//		// FIXME should this just return an nil []monitor.Monitor or []monitor.Monitor{monitor.Monitor{ID: 1234}}
//		return errors.New("no UptimeRobot client given")
//	}
//	return c.GetMonitors(monitor.WithMonitors(psp.monitorIDs))
//}

type psp PublicStatusPage

type jsonPSP struct {
	psp
	StandardURL string `json:"standard_url"`
	CustomURL   string `json:"custom_url"`
	//MonitorIDs  string `json:"monitors"`
}

func (ac *PublicStatusPage) UnmarshalJSON(b []byte) error {
	var raw jsonPSP

	err := json.Unmarshal(b, &raw)
	if err != nil {
		return err
	}

	standardURL, err := url.Parse(raw.StandardURL)
	if err != nil {
		return err
	}

	customURL, err := url.Parse(raw.CustomURL)
	if err != nil {
		return err
	}

	//rawMonitorIDs := strings.Split(raw.MonitorIDs, "-")
	//monitorIDs := make([]int64, len(rawMonitorIDs))
	//for i, rawMonitorID := range rawMonitorIDs {
	//	id, err := strconv.ParseInt(rawMonitorID, 10, 64)
	//	if err != nil {
	//		return err
	//	}
	//	monitorIDs[i] = id
	//}

	*ac = PublicStatusPage{
		ID:           raw.ID,
		FriendlyName: raw.FriendlyName,
		//monitorIDs:   monitorIDs,
		monitorIDs:  raw.monitorIDs,
		Sort:        raw.Sort,
		Status:      raw.Status,
		StandardURL: standardURL,
		CustomURL:   customURL,
	}

	return nil
}

////////////////////

type PublicStatusPages []PublicStatusPage

func (psps PublicStatusPages) String() string {
	ids := make(map[string]string, len(psps))

	var combined strings.Builder
	for i, psp := range psps {
		id := strconv.FormatInt(psp.ID, 10)
		if _, ok := ids[id]; ok {
			continue
		}

		if i > 0 {
			combined.WriteString("-")
		}

		ids[id] = id
		combined.WriteString(id)
	}

	return combined.String()
}

func (psps *PublicStatusPages) MarshalText() ([]byte, error) {
	// FIXME not sure which is faster this or strings.builder
	//return []byte(psps.String()), nil
	var ids []string

	for _, psp := range *psps {
		id := strconv.FormatInt(psp.ID, 10)
		ids = append(ids, id)
	}

	combined := strings.Join(ids, "-")

	return []byte(combined), nil
}

func (psps *PublicStatusPages) UnmarshalText(text []byte) error {
	textIDs := strings.Split(string(text), "-")

	for _, textID := range textIDs {
		id, err := strconv.ParseInt(textID, 10, 64)
		if err != nil {
			return err
		}
		*psps = append(*psps, PublicStatusPage{ID: id})
	}

	return nil
}

// Set is used to create a list of Monitor from a dash-separated list of ID
func (psps *PublicStatusPages) Set(s string) error {
	return psps.UnmarshalText([]byte(s))
}
