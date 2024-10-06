package uptime_robot

import (
	"encoding/json"
	"github.com/WileESpaghetti/go-uptimerobot-v2/uptime_robot/monitors"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/ajg/form"

	"github.com/WileESpaghetti/go-uptimerobot-v2/uptime_robot/api"
	"github.com/WileESpaghetti/go-uptimerobot-v2/uptime_robot/models"
)

const (
	baseUrl   = "https://api.uptimerobot.com/v2/"
	userAgent = "go-uptimerobot-v2/1.0.0 (Language=Go)"
)

// Client makes requests to the UptimeRobot API version 2.0
type Client struct {
	// Can be any of the API key types supported by the API
	ApiKey string
	// User agent used when communicating with the UptimeRobot API.
	UserAgent string
	// Url should always be specified with a trailing slash
	Url string
	// Uses http.DefaultClient by default
	HttpClient *http.Client
}

type Envelope = api.Envelope

// NewClient returns a new UptimeRobot API version 2.0 client using the default settings and the given API key.
func NewClient(apiKey string) *Client {
	return &Client{ApiKey: apiKey,
		Url:        baseUrl,
		UserAgent:  userAgent,
		HttpClient: http.DefaultClient}
}

func (c *Client) NewRequest(apiMethod string, options interface{}) (*http.Request, error) {
	endpoint := c.Url + apiMethod

	postData := url.Values{}                                   // FIXME this section still feels awkward because we double assign postData, but we need to handle nil options
	if !(options == nil || reflect.ValueOf(options).IsNil()) { // FIXME https://mangatmodi.medium.com/go-check-nil-interface-the-right-way-d142776edef1
		optionsData, err := form.EncodeToValues(options)
		if err != nil {
			return nil, err
		}
		postData = optionsData
	}
	postData.Set("api_key", c.ApiKey)

	encodedForm := strings.NewReader(postData.Encode())

	req, err := http.NewRequest(http.MethodPost, endpoint, encodedForm)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return req, nil
}

func (c *Client) Get(method string, response interface{}, options interface{}) error {
	request, err := c.NewRequest(method, options)
	if err != nil {
		return err
	}

	r, err := c.HttpClient.Do(request)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	err = json.NewDecoder(r.Body).Decode(response)
	if err != nil {
		return err
	}

	if envelop, ok := response.(*Envelope); ok {
		if envelop.Stat == api.StatFail {
			return envelop.Error
		}
	}

	return err
}

func (c *Client) GetAccountDetails() (*models.Account, error) {
	env := &api.GetAccountDetails{}
	err := c.Get("getAccountDetails", env, nil)
	if err != nil {
		return nil, err
	}

	return &env.Account, err
}

func (c *Client) GetMonitors(options ...api.MonitorOptions) (monitors.Monitors, error) {
	env := &api.GetMonitors{}

	params := &api.GetMonitorsRequest{}
	for _, o := range options {
		o(params)
	}

	err := c.Get("getMonitors", env, params)
	if err != nil {
		return nil, err
	}

	return env.Monitors, err
}
