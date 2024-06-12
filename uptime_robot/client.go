package uptime_robot

import (
	"encoding/json"
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

type Client struct {
	ApiKey     string       `form:"api_key"`
	UserAgent  string       `form:"-"`
	Url        string       `form:"-"`
	HttpClient *http.Client `form:"-"`
}

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

	if envelop, ok := response.(*api.Envelope); ok {
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

func (c *Client) GetMonitors(options ...api.MonitorOptions) (models.Monitors, error) {
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
