package uptime_robot

import (
	"encoding/json"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/WileESpaghetti/go-uptimerobot-v2/uptime_robot/api"
	"github.com/WileESpaghetti/go-uptimerobot-v2/uptime_robot/models"
	"github.com/gorilla/schema"
)

const (
	baseUrl   = "https://api.uptimerobot.com/v2/"
	userAgent = "go-uptimerobot-v2/1.0.0 (Language=Go)"
)

type Client struct {
	ApiKey     string       `schema:"api_key"`
	UserAgent  string       `schema:"-"`
	Url        string       `schema:"-"`
	HttpClient *http.Client `schema:"-"`
}

func New(apiKey string) *Client {
	return &Client{ApiKey: apiKey,
		Url:        baseUrl,
		UserAgent:  userAgent,
		HttpClient: http.DefaultClient}
}

func (c *Client) NewRequest(apiMethod string, options interface{}) (*http.Request, error) {
	endpoint := c.Url + apiMethod

	form := url.Values{}

	encoder := schema.NewEncoder()
	err := encoder.Encode(c, form)
	if err != nil {
		return nil, err
	}

	if !(options == nil || reflect.ValueOf(options).IsNil()) { // FIXME https://mangatmodi.medium.com/go-check-nil-interface-the-right-way-d142776edef1
		// FIXME need a better way to register encoders. also encoder might be a long lived object (https://web.archive.org/web/20190418003941/www.gorillatoolkit.org/pkg/schema)
		//var mq api.GetMonitorsRequest
		if mq, ok := options.(*api.GetMonitorsRequest); ok {
			mq.RegisterEncoders(encoder)
			err = encoder.Encode(mq, form)
			if err != nil {
				return nil, err
			}
		}

	}

	encodedForm := strings.NewReader(form.Encode())
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

	if envelop, ok := response.(api.Envelope); ok {
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

func (c *Client) GetMonitors(options *api.GetMonitorsRequest) (*models.Monitors, error) {
	env := &api.GetMonitors{}
	err := c.Get("getMonitors", env, options)
	if err != nil {
		return nil, err
	}

	return &env.Monitors, err
}
