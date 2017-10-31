package uptimerobot

import (
	"net/http"
	"net/url"
)

// Client constants
const (
	defaultUserAgent = "go-uptimerobot-v2/1.0.0 (Language=Go)"
)

// API constants
const (
	baseUrl = "https://api.uptimerobot.com/v2/"
)

const (
	_POST_API_KEY = "api_key"
)

type Client struct {
	ApiKey string
	HttpClient http.Client
}

func New(apiKey string) *Client {
	return &Client{
		ApiKey: apiKey,
		HttpClient: http.Client{},
	}
}


func (c *Client) NewRequest(apiMethod string) (*http.Request, error) {
	apiUrl := baseUrl + apiMethod

	req, err := http.NewRequest("POST", apiUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", defaultUserAgent)

	form := url.Values{}
	form.Set(_POST_API_KEY, c.ApiKey)
	req.PostForm = form

	return req, nil
}


//func (*c Client) Do(req *http.Request) (*http.Response, error) {
//	return nil, nil
//}
