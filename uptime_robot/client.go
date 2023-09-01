package uptime_robot

import (
	"github.com/WileESpaghetti/go-uptimerobot-v2/uptime_robot/models"
	"net/http"
	"strings"
)

const (
	baseUrl = "https://api.uptimerobot.com/v2/"
	userAgent = "go-uptimerobot-v2/1.0.0 (Language=Go)"
)

type Client struct {
	HttpClient *http.Client
}

func New() *Client {
	return &Client{HttpClient: http.DefaultClient}
}

func NewRequest(apiMethod string) (*http.Request, error) {
	apiUrl := baseUrl + apiMethod

	encodedForm := strings.NewReader("")

	req, err := http.NewRequest(http.MethodPost, apiUrl, encodedForm)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return req, nil
}

func (c *Client) GetAccountDetails() (models.Account, error) {
	return models.Account{Email: "example@example.com",
		DownMonitors: 2,
		MonitorLimit: 50,
		MonitorInterval: 5,
		PausedMonitors: 1,
		UpMonitors: 5}, nil
}
