package uptime_robot

import (
	"encoding/json"
	"fmt"
	"github.com/WileESpaghetti/go-uptimerobot-v2/uptime_robot/api"
	"github.com/WileESpaghetti/go-uptimerobot-v2/uptime_robot/models"
	"net/http"
	"strings"
)

const (
	baseUrl = "https://api.uptimerobot.com/v2/"
	userAgent = "go-uptimerobot-v2/1.0.0 (Language=Go)"
)

type Client struct {
	ApiKey string
	HttpClient *http.Client
}

func New(apiKey string) *Client {
	return &Client{ApiKey: apiKey,
		HttpClient: http.DefaultClient}
}

func (c *Client) NewRequest(apiMethod string) (*http.Request, error) {
	apiUrl := baseUrl + apiMethod

	rawForm := fmt.Sprintf("api_key=%s", c.ApiKey)
	encodedForm := strings.NewReader(rawForm)

	req, err := http.NewRequest(http.MethodPost, apiUrl, encodedForm)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return req, nil
}

func (c *Client) GetAccountDetails() (*models.Account, error) {
	getAccountDetailsRequest, err := c.NewRequest("getAccountDetails")
	if err != nil {
		return nil, err
	}

	r, err := c.HttpClient.Do(getAccountDetailsRequest)
	if err != nil {
		return nil, err
	}

	defer r.Body.Close()
	accountDetailsResponse := &api.GetAccountDetails{}

	err = json.NewDecoder(r.Body).Decode(accountDetailsResponse)
	if err != nil {
		return nil, err
	}

	if accountDetailsResponse.Stat == api.StatFail {
		return nil, accountDetailsResponse.Error
	}

	return &accountDetailsResponse.Account, err
}
