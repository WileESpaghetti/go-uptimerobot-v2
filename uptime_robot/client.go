package uptime_robot

import (
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
